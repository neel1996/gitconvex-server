package tests

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	git2 "github.com/neel1996/gitconvex-server/git"
	"os"
	"path"
	"testing"
)

func TestCommitLogs(t *testing.T) {
	var repoPath string
	var r *git.Repository
	currentEnv := os.Getenv("GOTESTENV")
	fmt.Println("Environment : " + currentEnv)

	if currentEnv == "ci" {
		repoPath = "/home/runner/work/gitconvex-go-server/starfleet"
		r, _ = git.PlainOpen(repoPath)
	} else {
		cwd, _ := os.Getwd()
		r, _ = git.PlainOpen(path.Join(cwd, ".."))
	}
	logChan := make(chan git2.AllCommitData)

	type args struct {
		repo       *git.Repository
		commitChan chan git2.AllCommitData
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "Git logs test case", args: struct {
			repo       *git.Repository
			commitChan chan git2.AllCommitData
		}{repo: r, commitChan: logChan}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go git2.AllCommits(tt.args.repo, tt.args.commitChan)
			commits := <-logChan
			commitLength := commits.TotalCommits

			fmt.Printf("Total commits : %v", commitLength)

			if commitLength == 0 {
				t.Error("No commit logs received!")
			}
		})
	}
}
