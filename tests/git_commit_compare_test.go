package tests

import (
	"github.com/libgit2/git2go/v31"
	git2 "github.com/neel1996/gitconvex/git"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCompareCommit(t *testing.T) {
	r, _ := git.OpenRepository(os.Getenv("GITCONVEX_TEST_REPO"))

	sampleCommits := git2.CommitLogStruct{
		Repo:            r,
		ReferenceCommit: "",
	}.CommitLogs()

	type args struct {
		repo                *git.Repository
		baseCommitString    string
		compareCommitString string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "Git commit compare test case", args: struct {
			repo                *git.Repository
			baseCommitString    string
			compareCommitString string
		}{repo: r, baseCommitString: *sampleCommits.Commits[0].Hash, compareCommitString: *sampleCommits.Commits[1].Hash}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var testObj git2.CompareCommitInterface
			testObj = git2.CompareCommitStruct{
				Repo:                tt.args.repo,
				BaseCommitString:    tt.args.baseCommitString,
				CompareCommitString: tt.args.compareCommitString,
			}

			got := testObj.CompareCommit()

			assert.NotZero(t, len(got))
		})
	}
}
