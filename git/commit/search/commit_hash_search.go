package search

import (
	git "github.com/libgit2/git2go/v31"
	"github.com/neel1996/gitconvex/constants"
	"github.com/neel1996/gitconvex/git/commit"
	"github.com/neel1996/gitconvex/graph/model"
	"regexp"
)

type commitHashSearch struct {
	commits []git.Commit
	mapper  commit.Mapper
}

func (h commitHashSearch) New(commits []git.Commit, mapper commit.Mapper) Search {
	return commitHashSearch{commits: commits, mapper: mapper}
}

func (h commitHashSearch) Search(searchKey string) []*model.GitCommits {
	var (
		matchingCommits []*model.GitCommits
		counter         = 0
	)

	for _, c := range h.commits {
		if h.isExceedingSearchLimit(counter) {
			break
		}

		if isMatch, _ := regexp.MatchString(searchKey, c.Id().String()); isMatch {
			commitLog := h.mapper.Map([]git.Commit{c})
			matchingCommits = append(matchingCommits, commitLog...)
		}

		counter++
	}

	return matchingCommits
}

func (h commitHashSearch) isExceedingSearchLimit(searchLimitCounter int) bool {
	return searchLimitCounter == constants.SearchLimit
}
