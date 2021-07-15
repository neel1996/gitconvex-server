package commit

import (
	"fmt"
	git2go "github.com/libgit2/git2go/v31"
	libgit_interface "github.com/neel1996/gitconvex/git/interface"
	"github.com/neel1996/gitconvex/global"
)

type Total interface {
	Get() int
}

type totalCommits struct {
	repo    libgit_interface.Repository
	revWalk libgit_interface.RevWalk
}

// Get function returns the total number of commits from the repo and commit message of the most recent commit
func (t totalCommits) Get() int {
	var total = 0
	repo := t.repo

	logItr, itrErr := repo.Walk()
	if itrErr != nil {
		logger.Log(fmt.Sprintf("Repo has no logs -> %s", itrErr.Error()), global.StatusError)
		return total
	}

	commits, err := t.allCommitLogs(logItr)
	if err != nil {
		logger.Log(fmt.Sprintf("Unable to obtain commits for the repo"), global.StatusError)
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

func (t totalCommits) allCommitLogs(logItr libgit_interface.RevWalk) ([]git2go.Commit, error) {
	var commits []git2go.Commit
	_ = logItr.PushHead()

	err := logItr.Iterate(revIterator(commits))

	return commits, err
}

func revIterator(commits []git2go.Commit) git2go.RevWalkIterator {
	return func(commit *git2go.Commit) bool {
		if commit != nil {
			commits = append(commits, *commit)
			return true
		} else {
			return false
		}
	}
}

func NewTotalCommits(repo libgit_interface.Repository) Total {
	return totalCommits{repo: repo}
}
