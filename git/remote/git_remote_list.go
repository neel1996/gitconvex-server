package remote

import (
	"fmt"
	"github.com/neel1996/gitconvex/git/middleware"
	"github.com/neel1996/gitconvex/global"
	"github.com/neel1996/gitconvex/graph/model"
)

type List interface {
	GetAllRemotes() []*model.RemoteDetails
}

type listRemotes struct {
	repo             middleware.Repository
	remoteValidation Validation
}

func (l listRemotes) GetAllRemotes() []*model.RemoteDetails {
	var remoteList []*model.RemoteDetails
	repo := l.repo

	if validationErr := l.remoteValidation.ValidateRemoteFields(); validationErr != nil {
		logger.Log(validationErr.Error(), global.StatusError)
		return nil
	}

	list, listErr := repo.Remotes().List()
	if listErr != nil {
		logger.Log(listErr.Error(), global.StatusError)
		return nil
	}

	for _, remoteEntry := range list {
		remote, remoteErr := repo.Remotes().Lookup(remoteEntry)
		if remoteErr != nil {
			logger.Log(remoteErr.Error(), global.StatusError)
			continue
		}

		data := model.RemoteDetails{
			RemoteName: remote.Name(),
			RemoteURL:  remote.Url(),
		}
		remoteList = append(remoteList, &data)
	}

	if len(remoteList) == 0 {
		logger.Log(fmt.Sprintf("No remotes available for the repo"), global.StatusWarning)
		return nil
	}

	logger.Log(fmt.Sprintf("Remote data fetched => %+v", remoteList), global.StatusInfo)

	return remoteList
}

func NewRemoteList(repo middleware.Repository, remoteValidation Validation) List {
	return listRemotes{repo: repo, remoteValidation: remoteValidation}
}
