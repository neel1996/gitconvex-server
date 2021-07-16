package middleware

import git "github.com/libgit2/git2go/v31"

type RevWalk interface {
	Iterate(iterator git.RevWalkIterator) error
	PushHead() error
}

type revWalk struct {
	walk *git.RevWalk
}

func (w *revWalk) PushHead() error {
	return w.walk.PushHead()
}

func (w *revWalk) Iterate(iterator git.RevWalkIterator) error {
	return w.walk.Iterate(iterator)
}

func NewRevWalk(walk *git.RevWalk) RevWalk {
	return &revWalk{walk: walk}
}
