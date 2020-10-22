package api

import (
	"fmt"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/neel1996/gitconvex-server/git"
	"github.com/neel1996/gitconvex-server/global"
	"github.com/neel1996/gitconvex-server/graph/model"
	"strings"
)

func RepoStatus(repoId string) *model.GitRepoStatusResults {
	logger := global.Logger{}
	repoChan := make(chan git.RepoDetails)
	remoteChan := make(chan git.RemoteDataModel)
	branchChan := make(chan git.Branch)
	commitChan := make(chan []*object.Commit)
	trackedFileCountChan := make(chan int)

	go git.Repo(repoId, repoChan)

	var repoName *string
	r := <-repoChan
	repo := r.GitRepo

	remote := ""
	var remoteURL *string
	remoteURL = &remote
	go git.RemoteData(repo, remoteChan)
	remoteData := <-remoteChan
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

	go git.GetBranchList(repo, branchChan)
	branchList := <-branchChan
	currentBranch := &branchList.CurrentBranch
	branches := branchList.BranchList
	allBranches := branchList.AllBranchList

	logger.Log(fmt.Sprintf("Obtained branch info -- \n%v -- %v\n", branchList.CurrentBranch, branchList.BranchList), global.StatusInfo)

	var commitLength int
	var commitLengthPtr *int
	var commits []*object.Commit
	var latestCommit *string

	go git.CommitLogs(repo, commitChan)
	commits = <-commitChan
	latestCommit = &commits[0].Message
	commitLength = len(commits)
	commitLengthPtr = &commitLength

	go git.TrackedFileCount(repo, trackedFileCountChan)
	trackedFileCount := <-trackedFileCountChan

	return &model.GitRepoStatusResults{
		GitRemoteData:        remoteURL,
		GitRepoName:          repoName,
		GitBranchList:        branches,
		GitAllBranchList:     allBranches,
		GitCurrentBranch:     currentBranch,
		GitRemoteHost:        remoteData.RemoteHost,
		GitTotalCommits:      commitLengthPtr,
		GitLatestCommit:      latestCommit,
		GitTotalTrackedFiles: &trackedFileCount,
	}
}
