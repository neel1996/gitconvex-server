package middleware

import git "github.com/libgit2/git2go/v31"

type Remotes interface {
	Create(string, string) (*git.Remote, error)
	Get() git.RemoteCollection
	Delete(string) error
}

type remotes struct {
	git.RemoteCollection
}

func (r remotes) Create(name string, url string) (*git.Remote, error) {
	return r.RemoteCollection.Create(name, url)
}

func (r remotes) Get() git.RemoteCollection {
	return r.RemoteCollection
}

func (r remotes) Delete(name string) error {
	return r.RemoteCollection.Delete(name)
}

func NewRemotes(remoteCollection git.RemoteCollection) Remotes {
	return remotes{
		RemoteCollection: remoteCollection,
	}
}
