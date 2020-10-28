package git

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/neel1996/gitconvex-server/global"
	"github.com/neel1996/gitconvex-server/graph/model"
)

func CommitFileList(repo *git.Repository, commitHash string) []*model.GitCommitFileResult {
	logger := global.Logger{}
	var res []*model.GitCommitFileResult

	logger.Log(fmt.Sprintf("Fetching file details for commit %v", commitHash), global.StatusInfo)

	hash := plumbing.NewHash(commitHash)
	commit, err := repo.CommitObject(hash)

	if err != nil {
		logger.Log(err.Error(), global.StatusError)
		res = append(res, &model.GitCommitFileResult{
			Type:     "",
			FileName: "",
		})

		return res
	}

	currentTree, _ := commit.Tree()
	prev, parentErr := commit.Parents().Next()

	if parentErr != nil {
		logger.Log(parentErr.Error(), global.StatusError)
		res = append(res, &model.GitCommitFileResult{
			Type:     "",
			FileName: "",
		})

		return res
	}
	prevTree, _ := prev.Tree()
	//diff, _ := currentTree.Diff(prevTree)

	diff, diffErr := prevTree.Diff(currentTree)

	if diffErr != nil {
		logger.Log(diffErr.Error(), global.StatusError)
		res = append(res, &model.GitCommitFileResult{
			Type:     "",
			FileName: "",
		})
		return res
	} else {
		for _, change := range diff {
			action, _ := change.Action()
			_, file, _ := change.Files()

			actionType := action.String()

			if actionType == "Insert" {
				actionType = "Add"
			}

			res = append(res, &model.GitCommitFileResult{
				Type:     actionType[:1],
				FileName: file.Name,
			})
		}
		return res
	}

}
