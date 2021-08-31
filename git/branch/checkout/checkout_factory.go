package checkout

import (
	"fmt"
	"github.com/neel1996/gitconvex/git/branch"
	"github.com/neel1996/gitconvex/git/middleware"
	"github.com/neel1996/gitconvex/global"
	"strings"
)

type Factory interface {
	GetCheckoutAction() Checkout
}

type factory struct {
	repo             middleware.Repository
	branchName       string
	branchValidation branch.Validation
}

func (f factory) GetCheckoutAction() Checkout {
	logger.Log(fmt.Sprintf("Received branch %s", f.branchName), global.StatusInfo)
	if validationErr := f.validateBranchFields(); validationErr != nil {
		return nil
	}

	if strings.Contains(f.branchName, "remotes/") {
		return NewCheckoutRemoteBranch(f.repo, f.branchName)
	} else {
		return NewCheckOutLocalBranch(f.repo, f.branchName)
	}
}

func (f factory) validateBranchFields() error {
	logger.Log("Validating branch fields before checkout", global.StatusInfo)

	if err := f.branchValidation.ValidateBranchFields(f.branchName); err != nil {
		logger.Log(err.Error(), global.StatusError)
		return err
	}
	return nil
}

func NewCheckoutFactory(repo middleware.Repository, branchName string, branchValidation branch.Validation) Factory {
	return factory{
		repo:             repo,
		branchName:       branchName,
		branchValidation: branchValidation,
	}
}
