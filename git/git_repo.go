package git

import (
	git "github.com/go-git/go-git/v5"
	"github.com/neel1996/gitconvex-server/global"
	"github.com/neel1996/gitconvex-server/utils"
)

type RepoDetails struct {
	RepoId   string
	RepoPath string
	GitRepo  *git.Repository
}

func Repo(repoId string) (*RepoDetails, error) {
	var repoData []utils.RepoData
	var repoPath string
	logger := global.Logger{}

	repoData = utils.DataStoreFileReader()

	for _, repo := range repoData {
		if repo.RepoId == repoId {
			repoPath = repo.RepoPath
			break
		}
	}

	repository, err := git.PlainOpen(repoPath)

	if err != nil {
		logger.Log(err.Error(), global.StatusError)
		return nil, err
	} else {
		return &RepoDetails{
			RepoId:   repoId,
			RepoPath: repoPath,
			GitRepo:  repository,
		}, nil
	}
}
