package git

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/neel1996/gitconvex-server/global"
	"github.com/neel1996/gitconvex-server/graph/model"
	"go/types"
	"io/ioutil"
	"strings"
)

type LsFileInfo struct {
	Content           []*string
	Commits           []*string
	TotalTrackedCount *int
}

type dirCommitDataModel struct {
	dirNameList   []*string
	dirCommitList []*string
}

type fileCommitDataModel struct {
	fileNameList   []*string
	fileCommitList []*string
}

var selectedDir string

// pathFilterCheck validates if the path held by the log iterator is tracked by the repo

func pathFilterCheck(filterPath string) bool {
	if strings.Contains(filterPath, selectedDir) {
		return true
	}
	return false
}

// dirCommitHandler collects the commit messages for the directories present in the target repo

func DirCommitHandler(repo *git.Repository, dirList []*string, dirCommitChan chan dirCommitDataModel) {
	var fileFilterList []*string
	var commitList []*string
	logger := global.Logger{}

	for _, dirName := range dirList {
		logOptions := git.LogOptions{
			Order:      git.LogOrderDFSPost,
			PathFilter: pathFilterCheck,
			All:        true,
		}
		selectedDir = *dirName

		dirIter, itrErr := repo.Log(&logOptions)

		if itrErr != nil {
			logger.Log(fmt.Sprintf("Failed while fetching commit for -> %s --> %s", *dirName, itrErr.Error()), global.StatusError)
		} else {
			if idx, _ := dirIter.Next(); idx != nil {
				if idx.Message != "" {
					logger.Log(fmt.Sprintf("Fetching commits for dir -> %v", *dirName), global.StatusInfo)

					commitMsg := idx.Message
					dirEntry := *dirName + ":directory"
					fileFilterList = append(fileFilterList, &dirEntry)

					if strings.Contains(commitMsg, "\n") {
						commitMsg = strings.Split(commitMsg, "\n")[0]
					}
					commitList = append(commitList, &commitMsg)
				}
				dirIter.Close()
				continue
			}
		}
	}

	dirCommitChan <- dirCommitDataModel{
		dirNameList:   fileFilterList,
		dirCommitList: commitList,
	}

	close(dirCommitChan)
}

// fileCommitHandler collects the commit messages for the files present in the target repo

func FileCommitHandler(repo *git.Repository, fileList []*string, fileCommitChan chan fileCommitDataModel) {
	var fileFilterList []*string
	var commitList []*string
	logger := global.Logger{}

	for _, file := range fileList {
		head, _ := repo.Head()
		hash := head.Hash()

		commitItr, itrErr := repo.Log(&git.LogOptions{
			From:     hash,
			FileName: file,
			Order:    git.LogOrderDFSPost,
			All:      false,
		})
		if itrErr != nil {
			logger.Log(fmt.Sprintf("error occurred in commit traversal for -> %s --> %v", *file, itrErr.Error()), global.StatusError)
		} else {
			if commit, _ := commitItr.Next(); commit != nil {
				var fileStr string
				logger.Log(fmt.Sprintf("Fetching commits for file -> %v", *file), global.StatusInfo)

				if strings.Contains(*file, "/") {
					splitEntry := strings.Split(*file, "/")
					fileStr = splitEntry[len(splitEntry)-1] + ":File"
				} else {
					fileStr = *file + ":File"
				}
				fileFilterList = append(fileFilterList, &fileStr)
				trimMsg := strings.TrimSpace(commit.Message)
				commitList = append(commitList, &trimMsg)
				commitItr.Close()
			}
		}
	}

	fileCommitChan <- fileCommitDataModel{
		fileNameList:   fileFilterList,
		fileCommitList: commitList,
	}
	close(fileCommitChan)
}

// TrackedFileCount returns the total number of files tracked by the target git repo

func TrackedFileCount(repo *git.Repository, trackedFileCountChan chan int) {
	var totalFileCount int
	logger := global.Logger{}

	head, _ := repo.Head()
	hash := head.Hash()

	allCommits, _ := repo.CommitObject(hash)
	tObj, _ := allCommits.Tree()

	err := tObj.Files().ForEach(func(file *object.File) error {
		if file != nil {
			totalFileCount++
			return nil
		} else {
			return types.Error{Msg: "File from the tree is empty"}
		}
	})
	tObj.Files().Close()

	if err != nil {
		logger.Log(err.Error(), global.StatusError)
		trackedFileCountChan <- 0
	} else {
		logger.Log(fmt.Sprintf("Total Tracked Files : %v", totalFileCount), global.StatusInfo)
		trackedFileCountChan <- totalFileCount
	}
	close(trackedFileCountChan)
}

// ListFiles collects the list of tracked files and their latest respective commit messages
// Used to visualize the git repo in the front-end file explorer in a github explorer based fashion

func ListFiles(repo *git.Repository, repoPath string, directoryName string) *model.GitFolderContentResults {
	logger := global.Logger{}
	logger.Log("Collecting tracked file list from the repo", global.StatusInfo)

	var targetPath string
	var fileList []*string
	var dirList []*string
	var fileFilterList []*string
	var commitList []*string

	fileFilterList = nil
	commitList = nil

	if directoryName != "" {
		targetPath = repoPath + "/" + directoryName
	} else {
		targetPath = repoPath
	}

	content, _ := ioutil.ReadDir(targetPath)

	for _, files := range content {
		var fileName string
		if directoryName != "" {
			fileName = directoryName + "/" + files.Name()
		} else {
			fileName = files.Name()
		}
		if files.IsDir() && fileName != ".git" {
			dirName := fileName
			dirList = append(dirList, &dirName)
		} else {
			if fileName != ".git" {
				fileList = append(fileList, &fileName)
			}
		}
	}
	content = nil

	var fileCommitChan = make(chan fileCommitDataModel)
	var dirCommitChan = make(chan dirCommitDataModel)

	go FileCommitHandler(repo, fileList, fileCommitChan)
	repoFileContent := <-fileCommitChan

	fileFilterList = append(fileFilterList, repoFileContent.fileNameList...)
	commitList = append(commitList, repoFileContent.fileCommitList...)

	go DirCommitHandler(repo, dirList, dirCommitChan)
	repoDirContent := <-dirCommitChan

	fileFilterList = append(fileFilterList, repoDirContent.dirNameList...)
	commitList = append(commitList, repoDirContent.dirCommitList...)

	if len(fileFilterList) == 0 {
		noFilesMsg := "NO_TRACKED_FILES"
		return &model.GitFolderContentResults{
			TrackedFiles:     []*string{&noFilesMsg},
			FileBasedCommits: []*string{&noFilesMsg},
		}
	}

	return &model.GitFolderContentResults{
		TrackedFiles:     fileFilterList,
		FileBasedCommits: commitList,
	}
}
