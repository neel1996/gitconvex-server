package git

import (
	git "github.com/go-git/go-git/v5"
	"github.com/neel1996/gitconvex-server/global"
	"github.com/neel1996/gitconvex-server/utils"
)

func Repo(repoId string) (git.Repository, error) {
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
		return git.Repository{}, err
	} else {
		return *repository, nil
	}
}
