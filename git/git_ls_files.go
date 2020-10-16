package git

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/neel1996/gitconvex-server/global"
	"go/types"
	"io/ioutil"
	"strings"
)

type LsFileInfo struct {
	Content           []*string
	Commits           []*string
	TotalTrackedCount *int
}

func ListFiles(repo *git.Repository, repoPath string, lsFileChan chan *LsFileInfo) {
	logger := global.Logger{}
	var fileList []*string
	var dirList []*string
	var fileFilterList []*string
	var commitList []*string
	var totalFileCount *int

	tempCount := 0
	totalFileCount = &tempCount

	content, _ := ioutil.ReadDir(repoPath)

	for _, files := range content {
		fileName := files.Name()
		if files.IsDir() && fileName != ".git" {
			dirName := fileName
			dirList = append(dirList, &dirName)
		} else {
			fileList = append(fileList, &fileName)
		}
	}

	head, _ := repo.Head()
	hash := head.Hash()

	allCommits, _ := repo.CommitObject(hash)
	tObj, _ := allCommits.Tree()

	err := tObj.Files().ForEach(func(file *object.File) error {
		localCount := *totalFileCount + 1
		if file != nil {
			totalFileCount = &localCount
			return nil
		} else {
			return types.Error{Msg: "File from the tree is empty"}
		}
	})

	if err != nil {
		logger.Log(err.Error(), global.StatusError)
	}

	for _, file := range fileList {
		logger.Log(fmt.Sprintf("Fetching logs for - %s", *file), global.StatusInfo)
		commitItr, _ := repo.Log(&git.LogOptions{
			From:     hash,
			Order:    0,
			FileName: file,
			All:      false,
		})
		if commit, _ := commitItr.Next(); commit != nil {
			logger.Log(fmt.Sprintf("%s -- %s", *file, strings.Split(commit.Message, "\n")[0]), global.StatusInfo)
			fileStr := *file + ":File"
			fileFilterList = append(fileFilterList, &fileStr)
			commitList = append(commitList, &commit.Message)
		}
		commitItr.Close()
	}
	logger.Log(fmt.Sprintf("Total Tracked Files : %v", len(fileFilterList)), global.StatusInfo)
	lsFileChan <- &LsFileInfo{
		Content:           fileFilterList,
		Commits:           commitList,
		TotalTrackedCount: totalFileCount,
	}
}
