package remote

import (
	"errors"
	git2go "github.com/libgit2/git2go/v31"
)

type Validation interface {
	ValidateRemoteFields() error
}

type validation struct {
	repo *git2go.Repository
}

func (v validation) ValidateRemoteFields() error {
	validateRepoErr := v.validateRepo()
	if validateRepoErr != nil {
		return validateRepoErr
	}

	validateRemotesErr := v.validateRemoteCollection()
	if validateRemotesErr != nil {
		return validateRemotesErr
	}

	return nil
}

func (v validation) validateRepo() error {
	if v.repo == nil {
		return errors.New("repo is nil")
	}

	return nil
}

func (v validation) validateRemoteCollection() error {
	if v.repo.Remotes == (git2go.RemoteCollection{}) {
		return errors.New("remote collection is nil")
	}

	return nil
}

func NewRemoteValidation(repo *git2go.Repository) Validation {
	return validation{
		repo: repo,
	}
}
