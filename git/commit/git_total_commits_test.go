package commit

import (
	"fmt"
	git2go "github.com/libgit2/git2go/v31"
	"github.com/stretchr/testify/suite"
	"os"
	"path/filepath"
	"testing"
)

type TotalCommitsTestSuite struct {
	suite.Suite
	total      Total
	repo       *git2go.Repository
	noHeadRepo *git2go.Repository
}

func TestTotalCommitsTestSuite(t *testing.T) {
	suite.Run(t, new(TotalCommitsTestSuite))
}

func (suite *TotalCommitsTestSuite) SetupTest() {
	r, err := git2go.OpenRepository(os.Getenv("GITCONVEX_TEST_REPO"))
	if err != nil {
		fmt.Println(err)
	}

	noHeadPath := os.Getenv("GITCONVEX_TEST_REPO") + string(filepath.Separator) + "no_head"
	noHeadRepo, _ := git2go.OpenRepository(noHeadPath)

	suite.repo = r
	suite.noHeadRepo = noHeadRepo
	suite.total = NewTotalCommits(suite.repo)
}

func (suite *TotalCommitsTestSuite) TestGet_WhenLogsAreAvailable_ShouldReturnTotal() {
	got := suite.total.Get()

	suite.NotZero(got)
}

func (suite *TotalCommitsTestSuite) TestGet_WhenRepoHasNoLogs_ShouldReturnZero() {
	suite.total = NewTotalCommits(suite.noHeadRepo)

	got := suite.total.Get()

	suite.Zero(got)
}
