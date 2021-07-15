package commit

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	git2go "github.com/libgit2/git2go/v31"
	"github.com/neel1996/gitconvex/mocks"
	"github.com/stretchr/testify/suite"
	"os"
	"path/filepath"
	"testing"
)

type TotalCommitsTestSuite struct {
	suite.Suite
	mockController *gomock.Controller
	total          Total
	repo           *git2go.Repository
	mockRepo       *mocks.MockRepository
	mockWalker     *mocks.MockRevWalk
	noHeadRepo     *git2go.Repository
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

	suite.mockController = gomock.NewController(suite.T())
	suite.repo = r
	suite.noHeadRepo = noHeadRepo
	suite.mockWalker = mocks.NewMockRevWalk(suite.mockController)
	suite.mockRepo = mocks.NewMockRepository(suite.mockController)
	suite.total = NewTotalCommits(suite.mockRepo)
}

func (suite *TotalCommitsTestSuite) TestGet_WhenLogsAreAvailable_ShouldReturnTotal() {
	got := suite.total.Get()

	suite.NotZero(got)
}

func (suite *TotalCommitsTestSuite) TestGet_WhenRepoHasNoLogs_ShouldReturnZero() {
	suite.total = NewTotalCommits(suite.mockRepo)
	suite.mockRepo.EXPECT().Walk().Return(nil, errors.New("WALKER_ERROR"))

	got := suite.total.Get()

	suite.Zero(got)
}
