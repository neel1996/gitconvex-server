package commit

import (
	"fmt"
	"github.com/golang/mock/gomock"
	git2go "github.com/libgit2/git2go/v31"
	"github.com/neel1996/gitconvex/git/middleware"
	"github.com/neel1996/gitconvex/mocks"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type FileHistoryTestSuite struct {
	suite.Suite
	mockController *gomock.Controller
	repo           middleware.Repository
	mockRepo       *mocks.MockRepository
	fileHistory    FileHistory
}

func TestFileHistoryTestSuite(t *testing.T) {
	suite.Run(t, new(FileHistoryTestSuite))
}

func (suite *FileHistoryTestSuite) SetupTest() {
	r, err := git2go.OpenRepository(os.Getenv("GITCONVEX_TEST_REPO"))
	if err != nil {
		fmt.Println(err)
	}

	suite.mockController = gomock.NewController(suite.T())
	suite.repo = middleware.NewRepository(r)
	suite.mockRepo = mocks.NewMockRepository(suite.mockController)
	suite.fileHistory = NewFileHistory(suite.mockRepo)
}

func (suite *FileHistoryTestSuite) TestGet_WhenRepoHasMoreThanOneCommit_ShouldReturnDiffForChosenCommit() {
	suite.fileHistory = NewFileHistory(suite.repo)

	head, _ := suite.repo.Head()
	commit, _ := suite.repo.LookupCommit(head.Target())

	gotHistory, err := suite.fileHistory.Get(commit)

	suite.Nil(err)
	suite.NotZero(len(gotHistory))
}
