package git

import (
	"fmt"
	git2go "github.com/libgit2/git2go/v31"
	"github.com/neel1996/gitconvex-server/global"
	"github.com/neel1996/gitconvex-server/graph/model"
	"io/ioutil"
	"strings"
	"sync"
)

// This go file relies on git installed on the host or the git client packed with the build application -> ./gitclient{.exe}
// Git client dependency was induced as the go-git based log traversal was highly time consuming

type ListFilesInterface interface {
	DirCommitHandler(dir *string)
	FileCommitHandler(file *string)
	TrackedFileCount(trackedFileCountChan chan int)
	ListFiles() *model.GitFolderContentResults
}

type fileDiffStruct struct {
	diffPath string
}

type commitEntry struct {
	fileEntries   []fileDiffStruct
	commitMessage string
}

type ListFilesStruct struct {
	Repo                 *git2go.Repository
	RepoPath             string
	DirectoryName        string
	Commits              []git2go.Commit
	AllCommitTreeEntries []commitEntry
	FileName             *string
	fileChan             chan string
	commitChan           chan string
	waitGroup            *sync.WaitGroup
}

type LsFileInfo struct {
	Content           []*string
	Commits           []*string
	TotalTrackedCount *int
}

var logger global.Logger
var waitGroup sync.WaitGroup

// DirCommitHandler collects the commit messages for the directories present in the target repo
func (l ListFilesStruct) DirCommitHandler(dirName *string) {
	fileChan := l.fileChan
	commitChan := l.commitChan
	waitGroup := l.waitGroup
	allCommitEntries := l.AllCommitTreeEntries

	var dirEntry = ""
	var commitMsg = ""

	for _, entry := range allCommitEntries {
		fileEntries := entry.fileEntries
		if dirEntry != "" {
			break
		}

		for _, diffEntry := range fileEntries {
			if strings.Contains(diffEntry.diffPath, *dirName) {
				dirEntry = *dirName
				commitMsg = entry.commitMessage
				break
			}
		}
	}
	if dirEntry != "" {
		logger.Log(fmt.Sprintf("Fetching commits for directory -> %s --> %s", dirEntry, commitMsg), global.StatusInfo)
	}
	dirStr := dirEntry + ":directory"
	fileChan <- dirStr
	commitChan <- commitMsg
	waitGroup.Done()
}

// FileCommitHandler collects the commit messages for the files present in the target repo
func (l ListFilesStruct) FileCommitHandler(file *string) {
	fileChan := l.fileChan
	commitChan := l.commitChan
	waitGroup := l.waitGroup
	allCommitEntries := l.AllCommitTreeEntries

	var fileStr string
	var commitMsg string
	var fileEntry = ""

	for _, entry := range allCommitEntries {
		fileEntries := entry.fileEntries
		if fileEntry != "" {
			break
		}

		for _, diffEntry := range fileEntries {
			if diffEntry.diffPath == *file {
				fileEntry = diffEntry.diffPath
				commitMsg = entry.commitMessage
				break
			}
		}
	}

	if fileEntry != "" {
		logger.Log(fmt.Sprintf("Fetching commits for file -> %v --> %s", *file, commitMsg), global.StatusInfo)

		if strings.Contains(*file, "/") {
			splitEntry := strings.Split(*file, "/")
			fileStr = splitEntry[len(splitEntry)-1] + ":File"
		} else {
			fileStr = *file + ":File"
		}
	}

	fileChan <- fileStr
	commitChan <- commitMsg
	waitGroup.Done()
}

// TrackedFileCount returns the total number of files tracked by the target git repo
func (l ListFilesStruct) TrackedFileCount(trackedFileCountChan chan int) {
	var totalFileCount int
	logger := global.Logger{}

	repo := l.Repo

	head, headErr := repo.Head()
	if headErr != nil {
		logger.Log(fmt.Sprintf("Repo head is invalid -> %s", headErr.Error()), global.StatusError)
		trackedFileCountChan <- 0
	} else {
		hash := head.Branch().Target()
		headCommit, _ := repo.LookupCommit(hash)

		if headCommit != nil {
			tree, treeErr := headCommit.Tree()
			if treeErr != nil {
				logger.Log(treeErr.Error(), global.StatusError)
				trackedFileCountChan <- totalFileCount
			} else {
				err := tree.Walk(func(s string, entry *git2go.TreeEntry) int {
					if entry.Id != nil && entry.Type.String() == "Blob" {
						totalFileCount++
						return 1
					}
					return 0
				})

				if err != nil {
					logger.Log(err.Error(), global.StatusError)
					trackedFileCountChan <- 0
				} else {
					logger.Log(fmt.Sprintf("Total Tracked Files : %v", totalFileCount), global.StatusInfo)
					trackedFileCountChan <- totalFileCount
				}
			}
		}
	}
	close(trackedFileCountChan)
}

// ListFiles collects the list of tracked files and their latest respective commit messages
//
// Used to display the git repo structure in the front-end file explorer in a github explorer based fashion
func (l ListFilesStruct) ListFiles() *model.GitFolderContentResults {
	logger := global.Logger{}
	logger.Log("Collecting tracked file list from the repo", global.StatusInfo)

	directoryName := l.DirectoryName
	repoPath := l.RepoPath
	r, _ := git2go.OpenRepository(l.RepoPath)

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

	var fileListChan = make(chan string)
	var commitListChan = make(chan string)

	l.waitGroup = &waitGroup
	l.fileChan = fileListChan
	l.commitChan = commitListChan

	var allCommits []git2go.Commit
	var commitEntries []commitEntry

	head, headErr := r.Head()

	if headErr == nil {
		commit, _ := r.LookupCommit(head.Target())
		for commit != nil {
			numParents := commit.ParentCount()
			if numParents > 0 {
				parentCommit := commit.Parent(0)
				parentTree, _ := parentCommit.Tree()
				commitTree, _ := commit.Tree()
				if parentTree != nil && commitTree != nil {
					diff, _ := r.DiffTreeToTree(parentTree, commitTree, nil)
					if diff != nil {
						numDeltas, _ := diff.NumDeltas()
						if numDeltas > 0 {
							entry := commitEntry{
								commitMessage: commit.Message(),
							}
							for d := 0; d < numDeltas; d++ {
								delta, _ := diff.Delta(d)
								fileEntry := fileDiffStruct{diffPath: delta.NewFile.Path}
								entry.fileEntries = append(entry.fileEntries, fileEntry)
							}
							commitEntries = append(commitEntries, entry)
						}
					}
				}
			}
			commit = commit.Parent(0)
		}
	} else {
		logger.Log(headErr.Error(), global.StatusError)
		return &model.GitFolderContentResults{
			TrackedFiles:     fileFilterList,
			FileBasedCommits: commitList,
		}
	}

	l.Commits = allCommits
	l.AllCommitTreeEntries = commitEntries

	for _, file := range fileList {
		waitGroup.Add(1)
		go l.FileCommitHandler(file)
		fileName := <-fileListChan
		commitMsg := <-commitListChan

		if commitMsg != "" {
			fileFilterList = append(fileFilterList, &fileName)
			commitList = append(commitList, &commitMsg)
		}
	}

	for _, dir := range dirList {
		waitGroup.Add(1)
		go l.DirCommitHandler(dir)
		fileName := <-fileListChan
		commitMsg := <-commitListChan

		if commitMsg != "" {
			fileFilterList = append(fileFilterList, &fileName)
			commitList = append(commitList, &commitMsg)
		}
	}

	waitGroup.Wait()
	close(fileListChan)
	close(commitListChan)

	if len(fileFilterList) == 0 || len(commitList) == 0 {
		msg := "NO_TRACKED_FILES"
		noFileList := []*string{&msg}
		return &model.GitFolderContentResults{
			TrackedFiles:     noFileList,
			FileBasedCommits: commitList,
		}
	}

	return &model.GitFolderContentResults{
		TrackedFiles:     fileFilterList,
		FileBasedCommits: commitList,
	}
}
