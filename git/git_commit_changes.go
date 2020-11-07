package git

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/neel1996/gitconvex-server/global"
)

func CommitChanges(repo *git.Repository, commitMessage string) string {
	logger := global.Logger{}
	w, wErr := repo.Worktree()

	if wErr != nil {
		logger.Log(fmt.Sprintf("Error occurred while fetching repo status -> %s", wErr.Error()), global.StatusError)
		return "COMMIT_FAILED"
	} else {
		hash, err := w.Commit(commitMessage, &git.CommitOptions{
			All: false,
		})
		if err != nil {
			logger.Log(fmt.Sprintf("Error occurred while committing changes -> %s", err.Error()), global.StatusError)
			return "COMMIT_FAILED"
		} else {
			logger.Log(fmt.Sprintf("Staged changes have been comitted - %s", hash.String()), global.StatusInfo)
			return "COMMIT_DONE"
		}
	}
}
