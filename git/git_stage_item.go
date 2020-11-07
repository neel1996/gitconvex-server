package git

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/neel1996/gitconvex-server/global"
)

func addError(fileItem string, errMsg string) string {
	logger := global.Logger{}
	logger.Log(fmt.Sprintf("Error occurred while staging %s -> %s", fileItem, errMsg), global.StatusError)
	return "ADD_ITEM_FAILED"
}

func StageItem(repo *git.Repository, fileItem string) string {
	logger := global.Logger{}
	w, wErr := repo.Worktree()
	if wErr != nil {
		return addError(fileItem, wErr.Error())
	} else {
		itemHash, addErr := w.Add(fileItem)
		if addErr != nil {
			return addError(fileItem, addErr.Error())
		} else {
			logger.Log(fmt.Sprintf("New item -> %s added to the repo worktree --> %s", fileItem, itemHash.String()), global.StatusInfo)
			return "ADD_ITEM_SUCCESS"
		}
	}

}
