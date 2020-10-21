package git

import (
	"fmt"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/neel1996/gitconvex-server/global"
	"github.com/neel1996/gitconvex-server/graph/model"
)

func DeleteBranch(repoId string, branchName string, forceFlag bool) *model.BranchDeleteStatus {
	var branchErr error
	logger := global.Logger{}
	repo := GetRepo(repoId)

	headRef, _ := repo.Head()
	ref := plumbing.NewHashReference(plumbing.ReferenceName(fmt.Sprintf("refs/heads/%v", branchName)), headRef.Hash())

	if forceFlag {
		logger.Log("Deleting branch "+branchName+" forcefully", global.StatusInfo)
		branchErr = repo.Storer.RemoveReference(ref.Name())
	} else {
		b, bErr := repo.Branch(branchName)
		if bErr != nil {
			fmt.Println(bErr.Error())
		} else {
			logger.Log("Deleting branch "+b.Name, global.StatusInfo)
			branchErr = repo.Storer.RemoveReference(ref.Name())
		}
	}

	if branchErr != nil {
		logger.Log(branchErr.Error(), global.StatusError)
		return &model.BranchDeleteStatus{Status: "BRANCH_DELETE_FAILED"}
	}

	return &model.BranchDeleteStatus{
		Status: "BRANCH_DELETE_SUCCESS",
	}
}
