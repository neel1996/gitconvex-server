package remote

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

type RemoteUrlDataTestSuite struct {
	suite.Suite
	repo           middleware.Repository
	mockController *gomock.Controller
	mockRepo       *mocks.MockRepository
	listRemoteUrl  ListRemoteUrl
}

func TestRemoteUrlDataTestSuite(t *testing.T) {
	suite.Run(t, new(RemoteUrlDataTestSuite))
}

func (suite *RemoteUrlDataTestSuite) SetupTest() {
	r, err := git2go.OpenRepository(os.Getenv("GITCONVEX_TEST_REPO"))
	if err != nil {
		fmt.Println(err)
	}

	suite.repo = middleware.NewRepository(r)
	suite.mockController = gomock.NewController(suite.T())
	suite.mockRepo = mocks.NewMockRepository(suite.mockController)
	suite.listRemoteUrl = NewRemoteUrlData(suite.mockRepo)
}

func (suite *RemoteUrlDataTestSuite) TestGetAllRemoteUrl_WhenRemotesArePresent_ShouldReturnRemoteUrlList() {
	suite.listRemoteUrl = NewRemoteUrlData(suite.repo)

	urlList := suite.listRemoteUrl.GetAllRemoteUrl()

	suite.NotZero(len(urlList))
	suite.Equal("https://github.com/neel1996/gitconvex-test.git", *urlList[0])
}

func (suite *RemoteUrlDataTestSuite) TestGetAllRemoteUrl_WhenRepoIsNil_ShouldReturnNil() {
	suite.listRemoteUrl = NewRemoteUrlData(nil)

	urlList := suite.listRemoteUrl.GetAllRemoteUrl()

	suite.Nil(urlList)
}

func (suite *RemoteUrlDataTestSuite) TestGetAllRemoteUrl_WhenRepoHasNoRemotes_ShouldReturnNil() {
	suite.listRemoteUrl = NewRemoteUrlData(suite.mockRepo)

	urlList := suite.listRemoteUrl.GetAllRemoteUrl()

	suite.Nil(urlList)
}

func (suite *RemoteUrlDataTestSuite) TestGetAllRemoteUrl_WhenRemotesAreNil_ShouldReturnNil() {
	suite.listRemoteUrl = NewRemoteUrlData(suite.mockRepo)

	urlList := suite.listRemoteUrl.GetAllRemoteUrl()

	suite.Nil(urlList)
}
