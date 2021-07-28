package search

import (
	git "github.com/libgit2/git2go/v31"
	"github.com/neel1996/gitconvex/constants"
	"regexp"
)

type commitAuthorSearch struct {
	commits []git.Commit
}

func (h commitAuthorSearch) New(commits []git.Commit) Search {
	return commitAuthorSearch{commits: commits}
}

func (h commitAuthorSearch) Search(searchKey string) []git.Commit {
	var (
		matchingCommits []git.Commit
		counter         = 0
	)

	for _, commit := range h.commits {
		if h.isExceedingSearchLimit(counter) {
			break
		}

		if isMatch, _ := regexp.MatchString(searchKey, commit.Author().Name); isMatch {
			matchingCommits = append(matchingCommits, commit)
		}

		counter++
	}

	return matchingCommits
}

func (h commitAuthorSearch) isExceedingSearchLimit(searchLimitCounter int) bool {
	return searchLimitCounter == constants.SearchLimit
}