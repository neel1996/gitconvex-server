package git

import (
	"fmt"
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/neel1996/gitconvex-server/global"
	"go/types"
	"strings"
)

type Branch struct {
	CurrentBranch string
	BranchList    []*string
	AllBranchList []*string
}

// GetBranchList fetches all the branches from the target repository
// The result will be returned as a struct with the current branch and all the available branches

func GetBranchList(repo *git.Repository) *Branch {
	var (
		branchList    *Branch
		branches      []*string
		allBranchList []*string
	)
	var currentBranch string

	logger := global.Logger{}

	if repo != nil {
		head, _ := repo.Head()
		currentBranch = head.Name().String()
		splitCurrentBranch := strings.Split(currentBranch, "/")
		currentBranch = splitCurrentBranch[len(splitCurrentBranch)-1]

		bIter, _ := repo.Branches()

		ref, _ := repo.References()
		_ = ref.ForEach(func(reference *plumbing.Reference) error {
			var (
				//refName    string
				refNamePtr *string
			)

			if reference.Name().String() != "HEAD" && strings.Contains(reference.Name().String(), "refs/") {
				refNameSplit := strings.Split(reference.Name().String(), "refs/")
				if len(refNameSplit) == 2 {
					logger.Log(fmt.Sprintf("Available Branch : %v", refNameSplit[1]), global.StatusInfo)
					if strings.Contains(refNameSplit[1], "heads/") {
						headBranch := strings.Split(refNameSplit[1], "heads/")[1]
						refNamePtr = &headBranch
					} else {
						refNamePtr = &refNameSplit[1]
					}
					allBranchList = append(allBranchList, refNamePtr)
				}
			}

			return nil
		})

		_ = bIter.ForEach(func(reference *plumbing.Reference) error {
			if reference != nil {
				localBranch := reference.String()
				splitBranch := strings.Split(localBranch, "/")
				localBranch = splitBranch[len(splitBranch)-1]

				branches = append(branches, &localBranch)
				return nil
			} else {
				return types.Error{Msg: "Empty reference"}
			}
		})
		bIter.Close()
	}

	branchList = &Branch{
		BranchList:    branches,
		CurrentBranch: currentBranch,
		AllBranchList: allBranchList,
	}

	logger.Log(fmt.Sprintf("Obtained branch info -- \n%v -- %v\n", branchList.CurrentBranch, branchList.BranchList), global.StatusInfo)
	return branchList
}
