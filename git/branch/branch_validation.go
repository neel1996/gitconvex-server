package branch

import (
	"github.com/neel1996/gitconvex/git/middleware"
)

type Validation interface {
	ValidateBranchFields() error
}

type validateBranch struct {
	repo       middleware.Repository
	branchName []string
}

func (v validateBranch) ValidateBranchFields() error {
	if v.repo == nil {
		return NilRepoError
	}

	for _, branchName := range v.branchName {
		if branchName == "" {
			return EmptyBranchNameError
		}
	}

	return nil
}

func NewBranchFieldsValidation(repo middleware.Repository, branchName ...string) Validation {
	return validateBranch{
		repo:       repo,
		branchName: branchName,
	}
}
