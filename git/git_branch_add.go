package git

import (
	"fmt"
	git2go "github.com/libgit2/git2go/v31"
	"github.com/neel1996/gitconvex-server/global"
)

type AddBranchInterface interface {
	AddBranch() string
}

type AddBranchInput struct {
	Repo       *git2go.Repository
	BranchName string
}

// AddBranch adds a new branch to the target repo
func (input AddBranchInput) AddBranch() string {
	logger := global.Logger{}

	repo := input.Repo
	branchName := input.BranchName
	head, headErr := repo.Head()

	logger.Log(fmt.Sprintf("Adding new branch -> %s", branchName), global.StatusInfo)
	if headErr != nil {
		logger.Log(fmt.Sprintf("Unable to fetch HEAD -> %s", headErr.Error()), global.StatusError)
		return global.BranchAddError
	} else {
		targetCommit, _ := repo.LookupCommit(head.Target())
		_, branchErr := repo.CreateBranch(branchName, targetCommit, false)

		if branchErr != nil {
			logger.Log(fmt.Sprintf("Failed to add branch - %s - %s", branchName, branchErr.Error()), global.StatusError)
			return global.BranchAddError
		}

		logger.Log(fmt.Sprintf("Added new branch - %s to the repo", branchName), global.StatusInfo)
		return global.BranchAddSuccess
	}
}
