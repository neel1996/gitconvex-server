package api

import (
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/neel1996/gitconvex-server/git"
	"github.com/neel1996/gitconvex-server/graph/model"
	"strings"
)

func RepoStatus(repoId string) *model.GitRepoStatusResults {
	var repoName *string
	r, _ := git.Repo(repoId)
	repo := r.GitRepo

	remote := ""
	var remoteURL *string
	remoteURL = &remote
	remotes := git.RemoteData(repo)

	sRemote := strings.Split(remotes[0], "/")
	repoName = &sRemote[len(sRemote)-1]

	if len(remotes) > 1 {
		*remoteURL = strings.Join(remotes, "||")
	} else {
		*remoteURL = remotes[0]
	}

	currentBranch := &git.GetBranchList(repo).CurrentBranch
	branches := git.GetBranchList(repo).BranchList

	var commitLength int
	var commitLengthPtr *int
	var commits []*object.Commit

	commits = git.CommitLogs(repo)
	commitLength = len(commits)
	commitLengthPtr = &commitLength

	trackedFileList := git.ListFiles(repo, r.RepoPath)
	trackedFileCount := len(trackedFileList)
	var trackedFileCountPtr *int
	trackedFileCountPtr = &trackedFileCount

	return &model.GitRepoStatusResults{
		GitRemoteData:        remoteURL,
		GitRepoName:          repoName,
		GitBranchList:        branches,
		GitAllBranchList:     nil,
		GitCurrentBranch:     currentBranch,
		GitRemoteHost:        nil,
		GitTotalCommits:      commitLengthPtr,
		GitLatestCommit:      nil,
		GitTrackedFiles:      trackedFileList,
		GitFileBasedCommit:   nil,
		GitTotalTrackedFiles: trackedFileCountPtr,
	}
}
