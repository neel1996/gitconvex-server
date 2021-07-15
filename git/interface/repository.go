package libgit_interface

import git "github.com/libgit2/git2go/v31"

type Repository interface {
	Walk() (*git.RevWalk, error)
}

type repository struct {
	repo *git.Repository
}

func (r repository) Walk() (*git.RevWalk, error) {
	return r.repo.Walk()
}

func NewRepository(repo *git.Repository) Repository {
	return repository{repo: repo}
}
