package search

import (
	git "github.com/libgit2/git2go/v31"
	"github.com/neel1996/gitconvex/constants"
	"regexp"
	"strings"
)

type commitAuthorSearch struct {
	commits []git.Commit
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

		if isMatch, _ := regexp.MatchString(h.toLower(searchKey), h.toLower(commit.Author().Name)); isMatch {
			matchingCommits = append(matchingCommits, commit)
			counter++
		}
	}

	return matchingCommits
}

func (h commitAuthorSearch) toLower(value string) string {
	return strings.ToLower(strings.TrimSpace(value))
}

func (h commitAuthorSearch) isExceedingSearchLimit(searchLimitCounter int) bool {
	return searchLimitCounter == constants.SearchLimit
}
