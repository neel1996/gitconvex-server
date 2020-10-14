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

	currentBranch := &git.GetBranchList(&repo).CurrentBranch
	splitCurrentBranch := strings.Split(*currentBranch, "/")
	currentBranch = &splitCurrentBranch[len(splitCurrentBranch)-1]

	var commitLength int
	var commitLengthPtr *int
	var commits []*object.Commit

	commits = git.CommitLogs(&repo)
	commitLength = len(commits)
	commitLengthPtr = &commitLength

	trackedFileList := git.ListFiles(&repo)
	trackedFileCount := len(*trackedFileList)
	var trackedFileCountPtr *int
	trackedFileCountPtr = &trackedFileCount

	return &model.GitRepoStatusResults{
		GitRemoteData:        nil,
		GitRepoName:          nil,
		GitBranchList:        nil,
		GitAllBranchList:     nil,
		GitCurrentBranch:     currentBranch,
		GitRemoteHost:        nil,
		GitTotalCommits:      commitLengthPtr,
		GitLatestCommit:      nil,
		GitTrackedFiles:      nil,
		GitFileBasedCommit:   nil,
		GitTotalTrackedFiles: trackedFileCountPtr,
	}
}
