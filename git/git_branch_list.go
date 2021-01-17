package git

import (
	"fmt"
	git2go "github.com/libgit2/git2go/v31"
	"github.com/neel1996/gitconvex-server/global"
	"go/types"
	"strings"
)

type BranchListInterface interface {
	GetBranchList(branchChan chan Branch)
}

type BranchListInputs struct {
	Repo *git2go.Repository
}

type Branch struct {
	CurrentBranch string
	BranchList    []*string
	AllBranchList []*string
}

// isBranchNameValid function checks if the branch name is valid to be return as an eligible branch
//
// The function filters out Tags and stashes returned as references
func isBranchNameValid(branchName string) bool {
	var branchNameValid bool
	branchNameValid = true

	if branchName == "HEAD" {
		branchNameValid = false
	}

	if !strings.Contains(branchName, "refs/") {
		branchNameValid = false
	}

	if strings.Contains(branchName, "tags/") {
		branchNameValid = false
	}

	return branchNameValid
}

// GetBranchList fetches all the branches from the target repository
// The result will be returned as a struct with the current branch and all the available branches
func (inputs BranchListInputs) GetBranchList(branchChan chan Branch) {
	var (
		branches      []*string
		allBranchList []*string
	)
	var currentBranch string
	repo := inputs.Repo
	logger := global.Logger{}

	if repo != nil {
		head, headErr := repo.Head()

		if headErr != nil {
			logger.Log(fmt.Sprintf("Repo head is invalid -> %s", headErr.Error()), global.StatusError)
			nilMsg := []string{"No Branches available", "Repo HEAD is nil", "No Branches Present"}
			branchChan <- Branch{
				BranchList:    []*string{&nilMsg[0]},
				CurrentBranch: nilMsg[1],
				AllBranchList: []*string{&nilMsg[2]},
			}
		} else {
			currentBranch = head.Name()
			splitCurrentBranch := strings.Split(currentBranch, "/")
			currentBranch = splitCurrentBranch[len(splitCurrentBranch)-1]

			allBranchList = append(allBranchList, &currentBranch)
			branches = append(branches, &currentBranch)

			localBranchIterator, itrErr := repo.NewBranchIterator(git2go.BranchAll)

			if itrErr == nil {
				_ = localBranchIterator.ForEach(func(branch *git2go.Branch, branchType git2go.BranchType) error {
					branchName, nameErr := branch.Name()
					if branch == nil {
						return types.Error{Msg: "branch iterator end"}
					}
					if nameErr == nil && !branch.IsTag() && !branch.IsNote() && branchName != currentBranch {
						if branch.IsRemote() && strings.Contains(branchName, "/") {
							splitString := strings.Split(branchName, "/")
							splitBranch := splitString[len(splitString)-1]
							if splitBranch != "HEAD" && splitBranch != currentBranch {
								concatRemote := "remotes/" + strings.Join(splitString, "/")
								allBranchList = append(allBranchList, &concatRemote)
							}
						} else {
							allBranchList = append(allBranchList, &branchName)
						}
					}
					return nil
				})
			}
			branchChan <- Branch{
				BranchList:    branches,
				CurrentBranch: currentBranch,
				AllBranchList: allBranchList,
			}
		}
	}

	close(branchChan)
}
