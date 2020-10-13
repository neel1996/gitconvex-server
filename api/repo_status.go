package api

import (
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/neel1996/gitconvex-server/git"
	"github.com/neel1996/gitconvex-server/graph/model"
	"strings"
)

func RepoStatus(repoId string) *model.GitRepoStatusResults {
	repo, _ := git.Repo(repoId)
	git.RemoteData(&repo)
	git.GetBranchList(&repo)

	var cBranch string
	currentBranch := &git.GetBranchList(&repo).CurrentBranch

	cBranch = func() string {
		splitBranch := strings.Split(*currentBranch, " ")[1]
		slashSplit := strings.Split(splitBranch, "/")
		return slashSplit[len(slashSplit)-1]
	}()

	var currentBranchPtr *string
	currentBranchPtr = &cBranch

	var commitLength int
	var commitLengthPtr *int
	var commits []*object.Commit

	commits = git.CommitLogs(&repo)
	commitLength = len(commits)
	commitLengthPtr = &commitLength

	return &model.GitRepoStatusResults{
		GitRemoteData:        nil,
		GitRepoName:          nil,
		GitBranchList:        nil,
		GitAllBranchList:     nil,
		GitCurrentBranch:     currentBranchPtr,
		GitRemoteHost:        nil,
		GitTotalCommits:      commitLengthPtr,
		GitLatestCommit:      nil,
		GitTrackedFiles:      nil,
		GitFileBasedCommit:   nil,
		GitTotalTrackedFiles: nil,
	}
}
