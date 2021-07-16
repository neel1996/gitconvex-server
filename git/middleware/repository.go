package middleware

import git "github.com/libgit2/git2go/v31"

type Repository interface {
	Walk() (RevWalk, error)
	Head() (Reference, error)
	LookupCommit(oid *git.Oid) (*git.Commit, error)
}

type repository struct {
	repo *git.Repository
}

func (r repository) Head() (Reference, error) {
	head, err := r.repo.Head()
	if err != nil {
		return nil, err
	}
	return NewReference(head), nil
}

func (r repository) LookupCommit(oid *git.Oid) (*git.Commit, error) {
	return r.repo.LookupCommit(oid)
}

func (r repository) Walk() (RevWalk, error) {
	walk, err := r.repo.Walk()

	return NewRevWalk(walk), err
}

func NewRepository(repo *git.Repository) Repository {
	return repository{repo: repo}
}
