package utils

import (
	"encoding/json"
	"github.com/neel1996/gitconvex-server/global"
	"io/ioutil"
)

type RepoData struct {
	RepoId   string `json:"id"`
	RepoName string `json:"repoName"`
	RepoPath string `json:"repoPath"`
}

func DataStoreFileReader() []RepoData {
	logger := global.Logger{}

	envConfig := EnvConfigFileReader()
	dbFile := envConfig.DataBaseFile
	dbFileContent, err := ioutil.ReadFile(dbFile)

	if err != nil {
		logger.Log(err.Error(), global.StatusError)
	} else {
		var repoData []RepoData
		if unmarshallErr := json.Unmarshal(dbFileContent, &repoData); unmarshallErr != nil {
			logger.Log(unmarshallErr.Error(), global.StatusError)
		}
		return repoData
	}
	return nil
}
