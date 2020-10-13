package git

import (
	"fmt"
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/neel1996/gitconvex-server/global"
	"go/types"
)

func CommitLogs(repo *git.Repository) []*object.Commit {
	logIter, _ := repo.Log(&git.LogOptions{})
	logger := global.Logger{}
	var commits []*object.Commit

	err := logIter.ForEach(func(commit *object.Commit) error {
		if commit != nil {
			commits = append(commits, commit)
			return nil
		} else {
			return types.Error{Msg: "Empty commit"}
		}
	})

	if err != nil {
		logger.Log(fmt.Sprintf("Unable to obtain commits for the repo"), global.StatusError)
		return nil
	} else {
		return commits
	}
}
