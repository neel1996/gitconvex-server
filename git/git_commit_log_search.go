package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/neel1996/gitconvex-server/graph/model"
	"regexp"
)

func SearchCommitLogs(repo *git.Repository, searchType string, searchKey string) []*model.GitCommits {
	var searchResult []*model.GitCommits

	commitLogs, _ := repo.Log(&git.LogOptions{
		Order: git.LogOrderDefault,
		All:   true,
	})

	_ = commitLogs.ForEach(func(commit *object.Commit) error {
		if len(searchResult) > 10 {
			return nil
		}

		switch searchType {
		case "message":
			if isMatch, _ := regexp.MatchString(searchKey, commit.Message); isMatch {
				commitLog := commitOrganizer([]object.Commit{*commit})
				searchResult = append(searchResult, commitLog...)
			}
			break
		case "hash":
			if isMatch, _ := regexp.MatchString(searchKey, commit.Hash.String()); isMatch {
				commitLog := commitOrganizer([]object.Commit{*commit})
				searchResult = append(searchResult, commitLog...)
			}
			break
		case "user":
			if isMatch, _ := regexp.MatchString(searchKey, commit.Author.Name); isMatch {
				commitLog := commitOrganizer([]object.Commit{*commit})
				searchResult = append(searchResult, commitLog...)
			}
			break
		}
		return nil
	})

	commitLogs.Close()
	return searchResult
}
