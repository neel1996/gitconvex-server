package commit

import (
	"errors"
	"fmt"
	git2go "github.com/libgit2/git2go/v31"
	"github.com/neel1996/gitconvex/constants"
	"github.com/neel1996/gitconvex/git/middleware"
	"github.com/neel1996/gitconvex/global"
	"github.com/neel1996/gitconvex/graph/model"
)

type FileHistory interface {
	Get(*git2go.Commit) ([]*model.GitCommitFileResult, error)
}

type fileHistory struct {
	repo middleware.Repository
}

func (f fileHistory) Get(commit *git2go.Commit) ([]*model.GitCommitFileResult, error) {
	commitHash := commit.Id().String()

	logger.Log(fmt.Sprintf("Fetching file details for commit %v", commitHash), global.StatusInfo)

	prevTree, commitTree, treeErr := f.treesOf(commit)
	if treeErr != nil {
		return fileHistoryError(treeErr)
	}

	return f.diffBetweenTrees(prevTree, commitTree)
}

func (f fileHistory) diffBetweenTrees(previousTree *git2go.Tree, currentTree *git2go.Tree) ([]*model.GitCommitFileResult, error) {
	var history []*model.GitCommitFileResult

	logger.Log("Checking diff between trees", global.StatusInfo)
	diff, diffErr := f.repo.DiffTreeToTree(previousTree, currentTree, nil)
	if diffErr != nil {
		return fileHistoryError(diffErr)
	}

	numDelta, numDeltaErr := diff.NumDeltas()
	if numDeltaErr != nil {
		return fileHistoryError(numDeltaErr)
	}

	for d := 0; d < numDelta; d++ {
		delta, _ := diff.Delta(d)
		status := delta.Status.String()
		history = append(history, &model.GitCommitFileResult{
			Type:     status[0:1],
			FileName: delta.NewFile.Path,
		})
	}

	fmt.Println(numDelta, history)
	return history, nil
}

func (f fileHistory) treesOf(commit *git2go.Commit) (*git2go.Tree, *git2go.Tree, error) {
	var treeErr error
	logger.Log("Getting current and previous trees", global.StatusInfo)

	parents := commit.ParentCount()
	if parents == 0 {
		return nil, nil, errors.New("commit has no parents")
	}

	previousCommit := commit.Parent(0)

	previousTree, treeErr := previousCommit.Tree()
	currentTree, treeErr := commit.Tree()
	if treeErr != nil {
		return nil, nil, treeErr
	}

	return previousTree, currentTree, nil
}

func fileHistoryError(err error) ([]*model.GitCommitFileResult, error) {
	logger.Log(err.Error(), global.StatusError)
	return nil, constants.CommitFileHistoryError
}

func NewFileHistory(repo middleware.Repository) FileHistory {
	return fileHistory{repo: repo}
}
