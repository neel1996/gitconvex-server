package api

import (
	"encoding/json"
	"fmt"
	"github.com/neel1996/gitconvex-server/global"
	"github.com/neel1996/gitconvex-server/graph/model"
	"github.com/neel1996/gitconvex-server/utils"
)

func FetchRepo() *model.FetchRepoParams {
	var (
		repoId   []string
		repoName []string
		repoPath []string
	)

	repoData := utils.DataStoreFileReader()

	for _, repo := range repoData {
		repoId = append(repoId, repo.RepoId)
		repoName = append(repoName, repo.RepoName)
		repoPath = append(repoPath, repo.RepoPath)
	}

	logger := global.Logger{}
	jsonContent, _ := json.MarshalIndent(repoData, "", " ")
	logger.Log(fmt.Sprintf("Reading data file content \n%v", string(jsonContent)), global.StatusInfo)

	return &model.FetchRepoParams{
		RepoID:   repoId,
		RepoName: repoName,
		RepoPath: repoPath,
	}
}
