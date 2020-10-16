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
	remoteData := git.RemoteData(repo)
	remotes := remoteData.RemoteURL

	sRemote := strings.Split(*remotes[0], "/")
	repoName = &sRemote[len(sRemote)-1]

	if len(remotes) > 1 {
		var tempRemoteArray []string
		for _, ptrRemote := range remotes {
			tempRemoteArray = append(tempRemoteArray, *ptrRemote)
		}
		*remoteURL = strings.Join(tempRemoteArray, "||")
	} else {
		*remoteURL = *remotes[0]
	}

	branchList := git.GetBranchList(repo)
	currentBranch := &branchList.CurrentBranch
	branches := branchList.BranchList
	allBranches := branchList.AllBranchList

	var commitLength int
	var commitLengthPtr *int
	var commits []*object.Commit
	var latestCommit *string

	commits = git.CommitLogs(repo)
	latestCommit = &commits[0].Message
	commitLength = len(commits)
	commitLengthPtr = &commitLength

	lsFileInfo := git.ListFiles(repo, r.RepoPath)
	trackedFileList := lsFileInfo.Content
	trackedFileCount := len(lsFileInfo.Content)
	trackedFileCommits := lsFileInfo.Commits

	var trackedFileCountPtr *int
	trackedFileCountPtr = &trackedFileCount

	return &model.GitRepoStatusResults{
		GitRemoteData:        remoteURL,
		GitRepoName:          repoName,
		GitBranchList:        branches,
		GitAllBranchList:     allBranches,
		GitCurrentBranch:     currentBranch,
		GitRemoteHost:        remoteData.RemoteHost,
		GitTotalCommits:      commitLengthPtr,
		GitLatestCommit:      latestCommit,
		GitTrackedFiles:      trackedFileList,
		GitFileBasedCommit:   trackedFileCommits,
		GitTotalTrackedFiles: trackedFileCountPtr,
	}
}
