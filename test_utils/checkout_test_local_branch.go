package test_utils

import (
	"github.com/neel1996/gitconvex/git/branch/checkout"
	"github.com/neel1996/gitconvex/git/middleware"
	"log"
)

func CheckoutTestLocalBranch(repo middleware.Repository, branchName string) {
	err := checkout.NewCheckOutLocalBranch(repo, branchName).CheckoutBranch()
	if err != nil {
		log.Println(err)
		return
	}
}
