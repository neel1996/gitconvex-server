package remote

import (
	"fmt"
	"github.com/neel1996/gitconvex/git/middleware"
	"github.com/neel1996/gitconvex/global"
)

type Delete interface {
	DeleteRemote() error
}

type deleteRemote struct {
	repo       middleware.Repository
	remoteName string
}

// DeleteRemote deletes the remote based on the specified remoteName
func (d deleteRemote) DeleteRemote() error {
	validationError := NewRemoteValidation(d.repo, d.remoteName).ValidateRemoteFields()
	if validationError != nil {
		return validationError
	}

	err := d.deleteSelectedRemote(d.remoteName)
	if err != nil {
		logger.Log(fmt.Sprintf("Remote => %s cannot be found in the repo", d.remoteName), global.StatusError)
		return err
	}

	return nil
}

func (d *deleteRemote) deleteSelectedRemote(remoteEntry string) error {
	err := d.repo.Remotes().Delete(remoteEntry)
	if err != nil {
		logger.Log(err.Error(), global.StatusError)
		return err
	}

	logger.Log(fmt.Sprintf("The remote => %s deleted from repo", d.remoteName), global.StatusInfo)
	return nil
}

func NewDeleteRemote(repo middleware.Repository, remoteName string) Delete {
	return deleteRemote{
		repo:       repo,
		remoteName: remoteName,
	}
}
