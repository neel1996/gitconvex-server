package git

import (
	"fmt"
	git "github.com/go-git/go-git/v5"
	"github.com/neel1996/gitconvex-server/global"
	"go/types"
)

func InitHandler(repoPath string) (*ResponseModel, error) {
	_, err := git.PlainInit(repoPath, false)

	if err != nil {
		logger := &global.Logger{}
		logger.Log(fmt.Sprintf("Git repo init failed \n%v", err), global.StatusError)
		return nil, types.Error{
			Msg: fmt.Sprintf("Git repo clone failed \n%v", err),
		}
	}

	return &ResponseModel{
		Status:    "success",
		Message:   "Git repo has been initialized",
		HasFailed: false,
	}, nil

}
