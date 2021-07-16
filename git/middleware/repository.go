package libgit_interface

import git "github.com/libgit2/git2go/v31"

type Repository interface {
	Walk() (RevWalk, error)
}

type repository struct {
	repo *git.Repository
}

func (r repository) Walk() (RevWalk, error) {
	walk, err := r.repo.Walk()

	return NewRevWalk(walk), err
}

func NewRepository(repo *git.Repository) Repository {
	return repository{repo: repo}
}
