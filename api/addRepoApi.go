package api

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/neel1996/gitconvex-server/global"
	"github.com/neel1996/gitconvex-server/graph/model"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"
)

type repoData struct {
	Id        string `json:"id"`
	RepoName  string `json:"repoName"`
	RepoPath  string `json:"repoPath"`
	TimeStamp string `json:"timestamp"`
}

var logger *global.Logger

func repoIdGenerator(c chan string) {
	newUUID, _ := uuid.NewUUID()
	repoId := strings.Split(newUUID.String(), "-")[0]
	c <- repoId
}

func repoDataFileWriter(repoId string, repoName string, repoPath string) {
	rArray := make([]repoData, 1)

	rArray[0] = repoData{
		Id:        repoId,
		RepoName:  repoName,
		RepoPath:  repoPath,
		TimeStamp: time.Now().String(),
	}

	repoDataJSON, _ := json.Marshal(rArray)

	cwd, _ := os.Getwd()
	dbDir := path.Join(cwd, "/database/")
	dbFile := dbDir + "/" + "repo-datastore.json"

	fileStat, fileOpenErr := os.Open(dbFile)

	if fileOpenErr != nil {
		logger.Log(fmt.Sprintf("Error occurred while opening repo data JSON file \n%v", fileOpenErr), global.StatusError)
		dirErr := os.Mkdir(dbDir, 0644)
		if dirErr != nil {
			logger.Log(fmt.Sprintf("Error occurred creating database directory \n%v", dirErr), global.StatusError)
			panic(dirErr)
		}
	} else {
		repoDataContent, readErr := ioutil.ReadAll(fileStat)
		if readErr != nil {
			logger.Log(fmt.Sprintf("Error occurred while reading data file \n%v", readErr), global.StatusError)
			panic(readErr)
		} else {
			var rArray []repoData
			parseErr := json.Unmarshal(repoDataContent, &rArray)

			if parseErr != nil {
				logger.Log(fmt.Sprintf("%v", readErr), global.StatusError)

				err := ioutil.WriteFile(dbFile, repoDataJSON, 0644)
				if err != nil {
					logger.Log(fmt.Sprintf("Error occurred while writing repo data JSON file \n%v", err), global.StatusError)
				} else {
					logger.Log(fmt.Sprintf("New repo details added to data file \n%v", rArray), global.StatusInfo)
				}
			} else {
				newRepoData := append(rArray, repoData{
					Id:        repoId,
					RepoName:  repoName,
					RepoPath:  repoPath,
					TimeStamp: time.Now().String(),
				})

				newRepoDataJSON, _ := json.MarshalIndent(newRepoData, "", " ")
				err := ioutil.WriteFile(dbFile, newRepoDataJSON, 0644)
				if err != nil {
					logger := global.Logger{Message: fmt.Sprintf("Error occurred while writing repo data JSON file \n%v", err)}
					logger.LogError()
				}
			}
		}
	}

}

func AddRepo(repoName string, repoPath string) *model.AddRepoParams {
	var repoIdChannel = make(chan string)

	go repoIdGenerator(repoIdChannel)
	repoId := <-repoIdChannel
	go repoDataFileWriter(repoId, repoName, repoPath)

	close(repoIdChannel)

	return &model.AddRepoParams{
		RepoID:  repoId,
		Status:  "Repo Added",
		Message: "The new repository has been added to Gitconvex",
	}
}
