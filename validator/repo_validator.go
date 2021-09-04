package validator

import "github.com/neel1996/gitconvex/git/middleware"

type repoValidator struct {
	repo middleware.Repository
}

func (v repoValidator) Validate() error {
	if v.repo == nil {
		return NilRepoError
	}
	return nil
}

func NewRepoValidator(repo middleware.Repository) Validator {
	return repoValidator{repo: repo}
}
