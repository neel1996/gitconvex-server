package search

import (
	git "github.com/libgit2/git2go/v31"
	"github.com/neel1996/gitconvex/constants"
	"github.com/neel1996/gitconvex/git/commit"
	"github.com/neel1996/gitconvex/graph/model"
	"regexp"
)

type commitAuthorSearch struct {
	commits []git.Commit
	mapper  commit.Mapper
}

func (h commitAuthorSearch) New(commits []git.Commit, mapper commit.Mapper) Search {
	return commitAuthorSearch{commits: commits, mapper: mapper}
}

func (h commitAuthorSearch) Search(searchKey string) []*model.GitCommits {
	var (
		matchingCommits []*model.GitCommits
		counter         = 0
	)

	for _, c := range h.commits {
		if h.isExceedingSearchLimit(counter) {
			break
		}

		if isMatch, _ := regexp.MatchString(searchKey, c.Author().Name); isMatch {
			commitLog := h.mapper.Map([]git.Commit{c})
			matchingCommits = append(matchingCommits, commitLog...)
		}

		counter++
	}

	return matchingCommits
}

func (h commitAuthorSearch) isExceedingSearchLimit(searchLimitCounter int) bool {
	return searchLimitCounter == constants.SearchLimit
}
