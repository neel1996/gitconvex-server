package git

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/neel1996/gitconvex-server/global"
	"go/types"
)

func ListFiles(repo *git.Repository) *[]string {
	logger := global.Logger{}
	var fileList []string

	head, _ := repo.Head()
	commit, _ := repo.CommitObject(head.Hash())
	tObj, err := commit.Tree()

	if err != nil {
		logger.Log(err.Error(), global.StatusError)
	} else {
		err := tObj.Files().ForEach(func(file *object.File) error {
			if file != nil {
				fileList = append(fileList, file.Name)
				return nil
			}
			return types.Error{Msg: "File pointer is empty"}
		})
		if err != nil {
			logger.Log(err.Error(), global.StatusError)
		}
	}

	logger.Log(fmt.Sprintf("Total Tracked Files : %v", len(fileList)), global.StatusInfo)
	return &fileList
}
