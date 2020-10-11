package api

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/neel1996/gitconvex-server/git"
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

func localLogger(message string, status string) {
	logger := &global.Logger{Message: message}
	logger.Log(logger.Message, status)
}

// repoIdGenerator generates a unique ID for the newly added repo

func repoIdGenerator(c chan string) {
	newUUID, _ := uuid.NewUUID()
	repoId := strings.Split(newUUID.String(), "-")[0]
	c <- repoId
}

// repoDataFileWriter writes the new repo details to the repo_datastore.json file

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
	var isDbDirCreated bool

	if fileOpenErr != nil {
		localLogger(fmt.Sprintf("Error occurred while opening repo data JSON file \n%v", fileOpenErr), global.StatusError)

		dirErr := os.Mkdir(dbDir, 0755)
		_, err := os.Create(dbFile)

		if err != nil {
			localLogger(err.Error(), global.StatusError)
		}
		if dirErr != nil {
			localLogger(fmt.Sprintf("Error occurred creating database directory \n%v", dirErr), global.StatusError)
			panic(dirErr)
		}
		fileStat, _ = os.Open(dbFile)
		isDbDirCreated = true
	}

	if isDbDirCreated {
		repoDataContent, readErr := ioutil.ReadAll(fileStat)
		if readErr != nil {
			localLogger(fmt.Sprintf("Error occurred while reading data file \n%v", readErr), global.StatusError)
			panic(readErr)
		} else {
			var rArray []repoData
			parseErr := json.Unmarshal(repoDataContent, &rArray)

			if parseErr != nil {
				localLogger(fmt.Sprintf("%v", readErr), global.StatusError)

				err := ioutil.WriteFile(dbFile, repoDataJSON, 0644)
				if err != nil {
					localLogger(fmt.Sprintf("Error occurred while writing repo data JSON file \n%v", err), global.StatusError)
				} else {
					localLogger(fmt.Sprintf("New repo details added to data file \n%v", rArray), global.StatusError)
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
					localLogger(fmt.Sprintf("Error occurred while writing repo data JSON file \n%v", err), global.StatusError)
				}
			}
		}
	}

}

// AddRepo function gets the repository details and includes a record to the gitconvex repo datastore file
// If initSwitch is 'true' then the git repo init function will be invoked to initialize a new repo
// If cloneSwitch is 'true' then the repo will be cloned to the file system using the repoURL field

func AddRepo(repoName string, repoPath string, cloneSwitch bool, repoURL *string, initSwitch bool) *model.AddRepoParams {
	var repoIdChannel = make(chan string)

	_, invalidRepoErr := git.RepoValidator(repoPath)

	if invalidRepoErr != nil {
		localLogger(fmt.Sprintf("The repo is not a valid git repo\n%v", invalidRepoErr), global.StatusError)

		return &model.AddRepoParams{
			RepoID:  "",
			Status:  "Failed",
			Message: invalidRepoErr.Error(),
		}
	}

	go repoIdGenerator(repoIdChannel)
	repoId := <-repoIdChannel
	go repoDataFileWriter(repoId, repoName, repoPath)
	close(repoIdChannel)

	if cloneSwitch && len(*repoURL) > 0 {
		_, err := git.CloneHandler(repoPath, *repoURL)

		if err != nil {
			localLogger(fmt.Sprintf("%v", err), global.StatusError)

			return &model.AddRepoParams{
				RepoID:  "",
				Status:  "Failed",
				Message: err.Error(),
			}
		}
	}
	if initSwitch {
		_, err := git.InitHandler(repoPath)

		if err != nil {
			localLogger(fmt.Sprintf("%v", err), global.StatusError)

			return &model.AddRepoParams{
				RepoID:  "",
				Status:  "Failed",
				Message: err.Error(),
			}
		}
	}

	localLogger("New repo added to the datastore", global.StatusInfo)

	return &model.AddRepoParams{
		RepoID:  repoId,
		Status:  "Repo Added",
		Message: "The new repository has been added to Gitconvex",
	}
}
