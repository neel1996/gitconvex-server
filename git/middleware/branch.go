package middleware

import git2go "github.com/libgit2/git2go/v31"

type Branch interface {
	Target() *git2go.Oid
}

type branch struct {
	branch *git2go.Branch
}

func (b branch) Target() *git2go.Oid {
	return b.branch.Target()
}

func NewBranch(gitBranch *git2go.Branch) Branch {
	return branch{
		branch: gitBranch,
	}
}
