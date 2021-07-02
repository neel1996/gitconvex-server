package branch

import (
	"fmt"
	git2go "github.com/libgit2/git2go/v31"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type BranchListTestSuite struct {
	suite.Suite
	repo       *git2go.Repository
	branchList List
}

func TestBranchListTestSuite(t *testing.T) {
	suite.Run(t, new(BranchListTestSuite))
}

func (suite *BranchListTestSuite) SetupTest() {
	r, err := git2go.OpenRepository(os.Getenv("GITCONVEX_TEST_REPO"))
	if err != nil {
		fmt.Println(err)
	}
	suite.repo = r
	suite.branchList = NewBranchList(suite.repo)
}

func  (suite *BranchListTestSuite)