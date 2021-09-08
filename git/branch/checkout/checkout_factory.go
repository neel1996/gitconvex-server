package checkout

import (
	"fmt"
	"github.com/neel1996/gitconvex/git/middleware"
	"github.com/neel1996/gitconvex/global"
	"github.com/neel1996/gitconvex/validator"
	"strings"
)

type Factory interface {
	GetCheckoutAction() Checkout
}

type factory struct {
	repo            middleware.Repository
	branchName      string
	repoValidator   validator.Validator
	branchValidator validator.ValidatorWithStringFields
}

func (f factory) GetCheckoutAction() Checkout {
	logger.Log(fmt.Sprintf("Received branch %s", f.branchName), global.StatusInfo)
	if repoValidationErr := f.repoValidator.Validate(); repoValidationErr != nil {
		logger.Log(repoValidationErr.Error(), global.StatusError)
		return nil
	}

	if validationErr := f.validateBranchFields(); validationErr != nil {
		return nil
	}

	if strings.Contains(f.branchName, "remotes/") {
		return NewCheckoutRemoteBranch(f.repo, f.branchName, nil)
	} else {
		return NewCheckOutLocalBranch(f.repo, f.branchName)
	}
}

func (f factory) validateBranchFields() error {
	logger.Log("Validating branch fields before checkout", global.StatusInfo)

	if err := f.branchValidator.ValidateWithFields(f.branchName); err != nil {
		logger.Log(err.Error(), global.StatusError)
		return err
	}
	return nil
}

func NewCheckoutFactory(repo middleware.Repository, branchName string, repoValidator validator.Validator, branchValidator validator.ValidatorWithStringFields) Factory {
	return factory{
		repo:            repo,
		branchName:      branchName,
		repoValidator:   repoValidator,
		branchValidator: branchValidator,
	}
}
