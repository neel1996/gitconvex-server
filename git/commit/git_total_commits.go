package commit

import (
	"fmt"
	git2go "github.com/libgit2/git2go/v31"
	"github.com/neel1996/gitconvex/git/middleware"
	"github.com/neel1996/gitconvex/global"
)

type Total interface {
	Get() int
}

type totalCommits struct {
	repo    middleware.Repository
	revWalk middleware.RevWalk
}

type commitType struct {
	commits []git2go.Commit
}

// Get function returns the total number of commits from the repo and commit message of the most recent commit
func (t totalCommits) Get() int {
	var total = 0
	repo := t.repo

	allLogs := NewListAllLogs(repo)
	commits, err := allLogs.Get()
	if err != nil {
		logger.Log(fmt.Sprintf("Unable to obtain commits for the repo : %v", err.Error()), global.StatusError)
		return total
	}

	total = len(commits)
	if total == 0 {
		logger.Log("Repo has no commit logs", global.StatusError)
		return total
	}

	logger.Log(fmt.Sprintf("Total commits in the repo -> %v", total), global.StatusInfo)
	return total
}

func NewTotalCommits(repo middleware.Repository) Total {
	return totalCommits{repo: repo}
}
