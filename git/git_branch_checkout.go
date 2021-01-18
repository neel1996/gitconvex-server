package git

import (
	"fmt"
	git2go "github.com/libgit2/git2go/v31"
	"github.com/neel1996/gitconvex-server/global"
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

	repo := inputs.Repo
	branchName := inputs.BranchName

	if strings.Contains(branchName, "remotes/") {
		splitRef := strings.Split(branchName, "/")
		branchName = splitRef[len(splitRef)-1]
		referenceBranchName = "refs/heads/" + branchName
		isRemoteBranch = true
	} else {
		referenceBranchName = "refs/heads/" + branchName
	}

	// If the branch is a remote branch then a remote fetch will be performed and then the branch checkout will be initiated
	if isRemoteBranch {
		logger.Log(fmt.Sprintf("Branch - %s is a remote branch\nTrying with intermediate remote fetch!", branchName), global.StatusWarning)
		inputs.intermediateFetch()
		_, remoteBranchErr := repo.LookupBranch(branchName, git2go.BranchRemote)

		if remoteBranchErr != nil {
			logger.Log(remoteBranchErr.Error(), global.StatusError)
			return global.BranchCheckoutError
		}

		//TODO: Add logic to checkout remote branches
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
