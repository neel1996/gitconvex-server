package git

import (
	"fmt"
	git2go "github.com/libgit2/git2go/v31"
	"github.com/neel1996/gitconvex-server/global"
	"github.com/neel1996/gitconvex-server/graph/model"
)

type PullInterface interface {
	PullFromRemote() *model.PullResult
}

type PullStruct struct {
	Repo         *git2go.Repository
	RemoteURL    string
	RemoteBranch string
	RemoteName   string
	RepoName     string
	RepoPath     string
	AuthOption   string
	UserName     string
	Password     string
	SSHKeyPath   string
}

func returnPullErr(msg string) *model.PullResult {
	logger.Log(msg, global.StatusError)
	return &model.PullResult{
		Status:      global.PullFromRemoteError,
		PulledItems: nil,
	}
}

// PullFromRemote pulls the changes from the remote repository using the remote URL and branch name received
func (p PullStruct) PullFromRemote() *model.PullResult {
	repo := p.Repo
	remoteBranch := p.RemoteBranch
	remoteURL := p.RemoteURL

	var remoteDataObject RemoteDataInterface
	remoteDataObject = RemoteDataStruct{
		Repo:      repo,
		RemoteURL: remoteURL,
	}

	remoteName := remoteDataObject.GetRemoteName()
	targetRefPsec := "refs/remotes/" + remoteName + "/" + remoteBranch
	targetRemote, _ := repo.Remotes.Lookup(remoteName)

	if targetRemote == nil {
		return returnPullErr("Target remote is unavailable")
	}

	var fetchObject FetchInterface
	fetchObject = FetchStruct{
		Repo:         repo,
		RemoteName:   remoteName,
		RepoPath:     p.RepoPath,
		RemoteURL:    p.RemoteURL,
		RemoteBranch: p.RemoteBranch,
		RepoName:     p.RepoName,
		AuthOption:   p.AuthOption,
		UserName:     p.UserName,
		Password:     p.Password,
		SSHKeyPath:   p.SSHKeyPath,
	}
	fetchResult := fetchObject.FetchFromRemote()
	if fetchResult.Status == global.FetchFromRemoteError {
		return returnPullErr("Remote fetch failed!")
	}

	remoteRef, remoteRefErr := repo.References.Lookup(targetRefPsec)

	if remoteRefErr == nil {
		remoteCommit, _ := repo.LookupCommit(remoteRef.Target())
		fmt.Println(remoteRef.Name())
		fmt.Println(remoteCommit.Message())

		annotatedCommit, _ := repo.AnnotatedCommitFromRef(remoteRef)
		if annotatedCommit != nil {
			mergeAnalysis, _, mergeErr := repo.MergeAnalysis([]*git2go.AnnotatedCommit{annotatedCommit})
			if mergeErr != nil {
				return returnPullErr("Pull failed - " + mergeErr.Error())
			} else {
				if mergeAnalysis&git2go.MergeAnalysisUpToDate != 0 {
					logger.Log("No new changes to pull from remote", global.StatusWarning)
					msg := "No new changes to pull from remote"
					return &model.PullResult{
						Status:      global.PullNoNewChanges,
						PulledItems: []*string{&msg},
					}
				} else {
					err := repo.Merge([]*git2go.AnnotatedCommit{annotatedCommit}, nil, &git2go.CheckoutOptions{
						Strategy: git2go.CheckoutAllowConflicts,
					})
					if err != nil {
						return returnPullErr("Annotated merge failed : " + err.Error())
					} else {
						repoIndex, _ := repo.Index()
						if repoIndex.HasConflicts() {
							return returnPullErr("Conflicts encountered while pulling changes")
						}

						indexTree, indexTreeErr := repoIndex.WriteTree()
						if indexTreeErr != nil {
							return returnPullErr("Index Tree Error : " + indexTreeErr.Error())
						}
						remoteTree, treeErr := repo.LookupTree(indexTree)
						if treeErr != nil {
							return returnPullErr("Tree Error : " + treeErr.Error())
						}
						checkoutErr := repo.CheckoutTree(remoteTree, nil)
						if checkoutErr != nil {
							return returnPullErr("Tree checkout error : " + checkoutErr.Error())
						}

						localRef, localRefErr := repo.LookupBranch(remoteBranch, git2go.BranchLocal)
						if localRefErr != nil {
							return returnPullErr("Local Reference lookup error :" + localRefErr.Error())
						}
						_, targetErr := localRef.SetTarget(remoteRef.Target(), "")
						if targetErr != nil {
							return returnPullErr("Target set error : " + targetErr.Error())
						}
						head, _ := repo.Head()
						if head == nil {
							return returnPullErr("HEAD is nil")
						}
						headTarget, _ := head.SetTarget(remoteRef.Target(), "")
						if headTarget == nil {
							return returnPullErr("Unable to set target to HEAD")
						}

						logger.Log("New changes pulled from remote -> "+targetRemote.Name(), global.StatusInfo)
						msg := "New changed pulled from remote"
						return &model.PullResult{
							Status:      global.PullFromRemoteSuccess,
							PulledItems: []*string{&msg},
						}
					}
				}
			}
		}
	} else {
		return returnPullErr(remoteRefErr.Error())
	}
	return returnPullErr("")
}
