package tests

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	git2 "github.com/neel1996/gitconvex-server/git"
	"os"
	"path"
	"testing"
)

func TestListFiles(t *testing.T) {
	lsFileChan := make(chan *git2.LsFileInfo)
	var repoPath string

	cwd, _ := os.Getwd()
	repoPath = path.Join(cwd, "..")

	currentEnv := os.Getenv("GOTESTENV")
	fmt.Println("Environment : " + currentEnv)

	if currentEnv == "ci" {
		repoPath = "/home/runner/work/gitconvex-go-server/starfleet"
	}
	r, _ := git.PlainOpen(repoPath)

	type args struct {
		repo       *git.Repository
		repoPath   string
		lsFileChan chan *git2.LsFileInfo
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "Git ls files test case", args: struct {
			repo       *git.Repository
			repoPath   string
			lsFileChan chan *git2.LsFileInfo
		}{repo: r, repoPath: repoPath, lsFileChan: lsFileChan}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repoContent := git2.ListFiles(tt.args.repo, tt.args.repoPath)

			trackedFiles := repoContent.TrackedFiles
			commits := repoContent.FileBasedCommits

			if len(trackedFiles) == 0 || len(commits) == 0 {
				t.Error("Expected repo file data not received")
			}
		})
	}
}
