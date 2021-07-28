package search

import (
	git2go "github.com/libgit2/git2go/v31"
	"github.com/neel1996/gitconvex/git/commit"
	"github.com/neel1996/gitconvex/graph/model"
)

type Search interface {
	Search(string) []*model.GitCommits
	New([]git2go.Commit, commit.Mapper) Search
}
