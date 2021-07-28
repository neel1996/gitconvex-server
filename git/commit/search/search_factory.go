package search

import (
	git2go "github.com/libgit2/git2go/v31"
	"github.com/neel1996/gitconvex/git/commit"
	"github.com/neel1996/gitconvex/git/middleware"
)

type Type int

const (
	CommitHash = iota
	CommitMessage
	CommitAuthor
)

var typeMap = map[string]Type{
	"hash":    CommitHash,
	"message": CommitMessage,
	"author":  CommitAuthor,
}

func GetSearchAction(repo middleware.Repository, searchType string, commits []git2go.Commit) Search {
	mappedSearchType := typeMap[searchType]

	mapper := commit.NewMapper(repo, commit.NewFileHistory(repo))

	switch mappedSearchType {
	case CommitHash:
		return commitHashSearch{
			commits: commits,
			mapper:  mapper,
		}
	case CommitMessage:
		return commitMessageSearch{
			commits: commits,
			mapper:  mapper,
		}
	case CommitAuthor:
		return commitAuthorSearch{
			commits: commits,
			mapper:  mapper,
		}
	}

	return nil
}
