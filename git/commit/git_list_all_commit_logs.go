package commit

import (
	"fmt"
	git2go "github.com/libgit2/git2go/v31"
	"github.com/neel1996/gitconvex/git/middleware"
	"github.com/neel1996/gitconvex/global"
)

type ListAllLogs interface {
	Get() ([]git2go.Commit, error)
}

type listAllLogs struct {
	repo middleware.Repository
}

func (l listAllLogs) Get() ([]git2go.Commit, error) {
	repo := l.repo

	logItr, itrErr := repo.Walk()
	if itrErr != nil {
		logger.Log(fmt.Sprintf("Repo has no logs -> %s", itrErr.Error()), global.StatusError)
		return nil, itrErr
	}

	commits, err := l.allCommitLogs(logItr)
	if err != nil {
		logger.Log(fmt.Sprintf("Unable to obtain commits for the repo"), global.StatusError)
		return nil, err
	}

	return commits, nil
}

func (l listAllLogs) allCommitLogs(logItr middleware.RevWalk) ([]git2go.Commit, error) {
	var c commitType
	_ = logItr.PushHead()

	err := logItr.Iterate(revIterator(&c))

	return c.commits, err
}

func revIterator(c *commitType) git2go.RevWalkIterator {
	return func(commit *git2go.Commit) bool {
		if commit != nil {
			c.commits = append(c.commits, *commit)
			return true
		}

		return false
	}
}

func NewListAllLogs(repo middleware.Repository) ListAllLogs {
	return listAllLogs{
		repo: repo,
	}
}
