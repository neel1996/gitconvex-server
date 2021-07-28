package search

import (
	git "github.com/libgit2/git2go/v31"
	"github.com/neel1996/gitconvex/constants"
	"github.com/neel1996/gitconvex/git/commit"
	"github.com/neel1996/gitconvex/graph/model"
	"regexp"
)

type commitMessageSearch struct {
	commits []git.Commit
	mapper  commit.Mapper
}

func (m commitMessageSearch) New(commits []git.Commit, mapper commit.Mapper) Search {
	return commitMessageSearch{commits: commits, mapper: mapper}
}

func (m commitMessageSearch) Search(searchKey string) []*model.GitCommits {
	var (
		matchingCommits []*model.GitCommits
		counter         = 0
	)

	for _, c := range m.commits {
		if m.isExceedingSearchLimit(counter) {
			break
		}

		if isMatch, _ := regexp.MatchString(searchKey, c.Message()); isMatch {
			commitLog := m.mapper.Map([]git.Commit{c})
			matchingCommits = append(matchingCommits, commitLog...)
		}

		counter++
	}

	return matchingCommits
}

func (m commitMessageSearch) isExceedingSearchLimit(searchLimitCounter int) bool {
	return searchLimitCounter == constants.SearchLimit
}
