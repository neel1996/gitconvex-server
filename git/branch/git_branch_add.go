package branch

import (
	"fmt"
	git2go "github.com/libgit2/git2go/v31"
	"github.com/neel1996/gitconvex/git/middleware"
	"github.com/neel1996/gitconvex/global"
)

type Add interface {
	AddBranch() error
}

type addBranch struct {
	repo             middleware.Repository
	branchName       string
	remoteSwitch     bool
	targetCommit     *git2go.Commit
	branchValidation Validation
}

func (a addBranch) AddBranch() error {
	err := a.validateAddBranchFields()
	if err != nil {
		logger.Log(err.Error(), global.StatusError)
		return err
	}

	head, headErr := a.repo.Head()
	if headErr != nil {
		logger.Log(fmt.Sprintf("Unable to fetch HEAD -> %s", headErr.Error()), global.StatusError)
		return headErr
	}

	targetCommit, validationErr := a.validateTargetCommit(a.targetCommit, a.repo, head)
	if validationErr != nil {
		logger.Log(validationErr.Error(), global.StatusError)
		return validationErr
	}

	logger.Log(fmt.Sprintf("Adding new branch -> %s", a.branchName), global.StatusInfo)

	_, branchErr := a.repo.CreateBranch(a.branchName, targetCommit, false)
	if branchErr != nil {
		logger.Log(fmt.Sprintf("Failed to add branch - %s - %s", a.branchName, branchErr.Error()), global.StatusError)
		return branchErr
	}

	logger.Log(fmt.Sprintf("Added new branch - %s to the repo", a.branchName), global.StatusInfo)
	return nil
}

func (a addBranch) validateAddBranchFields() error {
	err := a.branchValidation.ValidateBranchFields(a.branchName)
	if err != nil {
		return err
	}
	return nil
}

func (a addBranch) validateTargetCommit(targetCommit *git2go.Commit, repo middleware.Repository, head middleware.Reference) (*git2go.Commit, error) {
	if targetCommit != nil {
		return targetCommit, nil
	}

	headCommit, headCommitErr := repo.LookupCommit(head.Target())
	if headCommit == nil {
		return nil, headCommitErr
	}

	return headCommit, nil
}

func NewAddBranch(repo middleware.Repository, branchName string, remoteSwitch bool, targetCommit *git2go.Commit, branchValidation Validation) Add {
	return addBranch{
		repo:             repo,
		branchName:       branchName,
		remoteSwitch:     remoteSwitch,
		targetCommit:     targetCommit,
		branchValidation: branchValidation,
	}
}
