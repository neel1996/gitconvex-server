package commit

import (
	"github.com/neel1996/gitconvex/global"
	"github.com/neel1996/gitconvex/graph/model"
)

var logger global.Logger

type Commit interface {
	GitCommitChange() (string, error)
}

type Operation struct {
	Changes     Changes
	Total       Total
	ListAllLogs ListAllLogs
	FileHistory FileHistory
	Mapper      Mapper
}

func (c Operation) GitCommitChange() (string, error) {
	err := c.Changes.Add()

	if err != nil {
		logger.Log(err.Error(), global.StatusError)
		return "", ChangeError
	}

	return global.CommitChangeSuccess, nil
}

func (c Operation) GitTotalCommits() int {
	return c.Total.Get()
}

func (c Operation) GitCommitLogs() ([]*model.GitCommits, error) {
	commits, logsErr := c.ListAllLogs.Get()
	if logsErr != nil {
		logger.Log(logsErr.Error(), global.StatusError)
		return nil, LogsError
	}

	return c.Mapper.Map(commits), nil
}
