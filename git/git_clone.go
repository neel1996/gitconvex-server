package git

import (
	"fmt"
	git2go "github.com/libgit2/git2go/v31"
	"github.com/neel1996/gitconvex-server/global"
	"github.com/neel1996/gitconvex-server/graph/model"
	"github.com/neel1996/gitconvex-server/utils"
	"go/types"
)

type CloneInterface interface {
	CloneHandler() (*model.ResponseModel, error)
	fallbackClone() (*model.ResponseModel, error)
}

type CloneStruct struct {
	RepoName   string
	RepoPath   string
	RepoURL    string
	AuthOption string
	UserName   string
	Password   string
	SSHKeyPath string
}

// fallbackClone performs a git clone using the native git client
// If the go-git based clone fails due to an authentication issue, then this function will be invoked to perform a clone
func (c CloneStruct) fallbackClone() (*model.ResponseModel, error) {
	repoPath := c.RepoPath
	repoURL := c.RepoURL

	args := []string{"clone", repoURL, repoPath}
	cmd := utils.GetGitClient(".", args)
	cmdStr, cmdErr := cmd.Output()

	if cmdErr != nil {
		logger.Log(fmt.Sprintf("Fallback clone failed -> %s", cmdErr.Error()), global.StatusError)
		return nil, cmdErr
	} else {
		logger.Log(fmt.Sprintf("New repo has been cloned to -> %s -> %s", repoPath, cmdStr), global.StatusInfo)
		return &model.ResponseModel{
			Status:    "success",
			Message:   "Git clone completed",
			HasFailed: false,
		}, nil
	}
}

// CloneHandler clones the remote repo to the target directory
// It supports options for SSH and HTTPS authentications
func (c CloneStruct) CloneHandler() (*model.ResponseModel, error) {
	authOption := c.AuthOption
	repoPath := c.RepoPath
	repoURL := c.RepoURL
	userName := c.UserName
	password := c.Password
	sshKeyPath := c.SSHKeyPath

	logger.Log(fmt.Sprintf("Initiating repo clone with - %v auth option", authOption), global.StatusInfo)
	var err error
	var r *git2go.Repository

	var remoteCBObject RemoteCallbackInterface
	remoteCBObject = &RemoteCallbackStruct{
		RepoName:   c.RepoName,
		UserName:   userName,
		Password:   password,
		AuthOption: authOption,
		SSHKeyPath: sshKeyPath,
	}
	var remoteCallbacks git2go.RemoteCallbacks
	remoteCallbacks = remoteCBObject.RemoteCallbackSelector()

	r, err = git2go.Clone(repoURL, repoPath, &git2go.CloneOptions{
		FetchOptions: &git2go.FetchOptions{
			RemoteCallbacks: remoteCallbacks,
		},
	})

	if err != nil {
		logger.Log(fmt.Sprintf("Error occurred while cloning repo \n%v", err), global.StatusError)
		return nil, types.Error{Msg: "Git repo clone failed"}
	}

	logger.Log(fmt.Sprintf("Repo %v - Cloned to target directory - %s", c.RepoName, r.Path()), global.StatusInfo)

	return &model.ResponseModel{
		Status:    "success",
		Message:   "Git clone completed",
		HasFailed: false,
	}, nil
}
