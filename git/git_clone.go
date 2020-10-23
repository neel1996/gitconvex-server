package git

import (
	"fmt"
	git "github.com/go-git/go-git/v5"
	"github.com/neel1996/gitconvex-server/global"
	"go/types"
)

func CloneHandler(repoPath string, repoURL string) (*ResponseModel, error){
	_, err := git.PlainClone(repoPath, false, &git.CloneOptions{
		URL: repoURL,
	})
	if err != nil {
		logger := &global.Logger{}
		logger.Log(fmt.Sprintf("Error occurred while cloning repo \n%v", err), global.StatusError)
		return nil, types.Error{Msg: "Git repo clone failed"}
	}

	return &ResponseModel{
		Status:    "success",
		Message:   "Git clone completed",
		HasFailed: false,
	}, nil
}
