package git

import (
	"fmt"
	git2go "github.com/libgit2/git2go/v31"
	"github.com/neel1996/gitconvex-server/global"
	"go/types"
	"strings"
)

type BranchCheckoutInterface interface {
	intermediateFetch()
	CheckoutBranch() string
}

type BranchCheckoutInputs struct {
	Repo       *git2go.Repository
	BranchName string
}

func returnCheckoutError(err error) string {
	logger.Log(err.Error(), global.StatusError)
	return global.BranchCheckoutError
}

// intermediateFetch performs a remote fetch if the selected checkout branch is a remote branch
func (inputs BranchCheckoutInputs) intermediateFetch() {
}

// CheckoutBranch checks out the branchName received as argument
func (inputs BranchCheckoutInputs) CheckoutBranch() string {
	var isRemoteBranch bool
	var referenceBranchName string
	var remoteBranchName string

	repo := inputs.Repo
	branchName := inputs.BranchName

	if strings.Contains(branchName, "remotes/") {
		singleBranchName := strings.Split(branchName, "/")
		referenceBranchName = "refs/heads/" + singleBranchName[len(singleBranchName)-1]
		remoteBranchSplit := strings.Split(branchName, "remotes/")
		remoteBranchName = remoteBranchSplit[len(remoteBranchSplit)-1]
		branchName = singleBranchName[len(singleBranchName)-1]
		isRemoteBranch = true
	} else {
		referenceBranchName = "refs/heads/" + branchName
	}

	if isRemoteBranch {
		logger.Log(fmt.Sprintf("Branch - %s is a remote branch. Trying with intermediate remote fetch!", branchName), global.StatusWarning)

		remoteBranch, remoteBranchErr := repo.LookupBranch(remoteBranchName, git2go.BranchRemote)
		if remoteBranchErr != nil {
			return returnCheckoutError(remoteBranchErr)
		} else {
			remoteHead := remoteBranch.Target()
			remoteCommit, remoteCommitErr := repo.LookupCommit(remoteHead)
			if remoteCommitErr != nil {
				return returnCheckoutError(remoteCommitErr)
			} else {
				remoteTree, remoteTreeErr := remoteCommit.Tree()
				if remoteTree != nil {
					checkoutErr := repo.CheckoutTree(remoteTree, &git2go.CheckoutOptions{Strategy: git2go.CheckoutSafe})
					if checkoutErr != nil {
						return returnCheckoutError(checkoutErr)
					} else {
						_, localLookupErr := repo.LookupBranch(branchName, git2go.BranchLocal)
						if localLookupErr != nil {
							logger.Log(localLookupErr.Error(), global.StatusError)

							var addBranchObject AddBranchInterface
							addBranchObject = AddBranchInput{
								Repo:         repo,
								BranchName:   branchName,
								RemoteSwitch: false,
								TargetCommit: remoteCommit,
							}

							branchCreateStatus := addBranchObject.AddBranch()
							if branchCreateStatus != "BRANCH_ADD_FAILED" {
								err := repo.SetHead(referenceBranchName)
								if err != nil {
									return returnCheckoutError(err)
								} else {
									return fmt.Sprintf("Head checked out to branch - %v", branchName)
								}
							}else{
								return returnCheckoutError(types.Error{Msg: "Branch creation failed"})
							}
						} else {
							err := repo.SetHead(referenceBranchName)
							if err != nil {
								return returnCheckoutError(err)
							} else {
								return fmt.Sprintf("Head checked out to branch - %v", branchName)
							}
						}
					}
				} else {
					return returnCheckoutError(remoteTreeErr)
				}
			}
		}
	}

	branch, branchErr := repo.LookupBranch(branchName, git2go.BranchLocal)
	if branchErr != nil {
		return returnCheckoutError(branchErr)
	}

	topCommit, _ := repo.LookupCommit(branch.Target())
	if topCommit != nil {
		tree, treeErr := topCommit.Tree()
		if treeErr != nil {
			return returnCheckoutError(treeErr)
		}

		checkoutErr := repo.CheckoutTree(tree, &git2go.CheckoutOptions{
			Strategy:       git2go.CheckoutSafe,
			DisableFilters: false,
		})

		if checkoutErr != nil {
			return returnCheckoutError(checkoutErr)
		}

		err := repo.SetHead(referenceBranchName)
		if err != nil {
			return returnCheckoutError(err)
		}
	}

	logger.Log(fmt.Sprintf("Current branch checked out to -> %s", branchName), global.StatusInfo)
	return fmt.Sprintf("Head checked out to branch - %v", branchName)
}
