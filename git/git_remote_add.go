package git

import (
	"fmt"
	git2go "github.com/libgit2/git2go/v31"
	"github.com/neel1996/gitconvex-server/global"
)

type AddRemoteInterface interface {
	AddRemote() string
}

type AddRemoteStruct struct {
	Repo       *git2go.Repository
	RemoteName string
	RemoteURL  string
}

// AddRemote adds a new remote to the target git repo
func (a AddRemoteStruct) AddRemote() string {
	repo := a.Repo
	remoteName := a.RemoteName
	remoteURL := a.RemoteURL

	remote, err := repo.Remotes.Create(remoteName, remoteURL)

	if err == nil {
		logger.Log(fmt.Sprintf("New remote %s added to the repo", remote.Name()), global.StatusInfo)
		return global.RemoteAddSuccess
	} else {
		logger.Log("Remote addition Failed -> "+err.Error(), global.StatusError)
		return global.RemoteAddError
	}
}
