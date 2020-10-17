package tests

import (
	"github.com/go-git/go-git/v5"
	git2 "github.com/neel1996/gitconvex-server/git"
	"os"
	"path"
	"testing"
)

func TestListFiles(t *testing.T) {
	lsFileChan := make(chan *git2.LsFileInfo)
	cwd, _ := os.Getwd()
	repoPath := path.Join(cwd, "..")
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
			go git2.ListFiles(tt.args.repo, tt.args.repoPath, lsFileChan)
			lsFiles := <-lsFileChan

			content := lsFiles.Content
			commits := lsFiles.Commits
			totalCommits := lsFiles.TotalTrackedCount

			if len(content) == 0 || len(commits) == 0 || *totalCommits == 0 {
				t.Error("Expected repo file data not received")
			}
		})
	}
}
