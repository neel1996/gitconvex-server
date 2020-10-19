package git

import (
	"fmt"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/neel1996/gitconvex-server/global"
)

func AddBranch(repoId string, branchName string) string {
	logger := global.Logger{}
	repo := GetRepo(repoId)

	headRef, _ := repo.Head()
	ref := plumbing.NewHashReference(plumbing.ReferenceName(fmt.Sprintf("refs/heads/%v", branchName)), headRef.Hash())
	branchErr := repo.Storer.SetReference(ref)

	if branchErr != nil {
		logger.Log(branchErr.Error(), global.StatusError)
		return ""
	}

	return fmt.Sprintf("Branch %v created", branchName)
}
