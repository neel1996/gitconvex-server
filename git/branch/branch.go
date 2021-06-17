package branch

import (
	"errors"
	"github.com/neel1996/gitconvex/global"
	"github.com/neel1996/gitconvex/graph/model"
)

var logger global.Logger

type Branch interface {
	GitAddBranch() (string, error)
	GitCheckoutBranch() (string, error)
	GitCompareBranches() ([]*model.BranchCompareResults, error)
}

type Operation struct {
	Add      Add
	Checkout Checkout
	Compare  Compare
}

func (b Operation) GitAddBranch() (string, error) {
	addBranchResult := b.Add.AddBranch()

	if addBranchResult == global.BranchAddError {
		return "", errors.New(global.BranchAddError)
	}

	return addBranchResult, nil
}

func (b Operation) GitCheckoutBranch() (string, error) {
	checkoutBranchResult := b.Checkout.CheckoutBranch()

	if checkoutBranchResult == global.BranchCheckoutError {
		return "", errors.New(global.BranchCheckoutError)
	}

	return checkoutBranchResult, nil
}

func (b Operation) GitCompareBranches() ([]*model.BranchCompareResults, error) {
	branchDiff := b.Compare.CompareBranch()

	if len(branchDiff) == 0 {
		return []*model.BranchCompareResults{}, errors.New("no difference between the two branches")
	}

	return branchDiff, nil
}
