package middleware

import git "github.com/libgit2/git2go/v31"

type Reference interface {
	Target() *git.Oid
}

type reference struct {
	ref *git.Reference
}

func (r reference) Target() *git.Oid {
	return r.ref.Target()
}

func NewReference(ref *git.Reference) Reference {
	return reference{ref: ref}
}
