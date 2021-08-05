package remote

import (
	"fmt"
	"github.com/golang/mock/gomock"
	git2go "github.com/libgit2/git2go/v31"
	"github.com/neel1996/gitconvex/git/middleware"
	"github.com/neel1996/gitconvex/graph/model"
	"github.com/neel1996/gitconvex/mocks"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type ListRemoteTestSuite struct {
	suite.Suite
	repo           middleware.Repository
	mockController *gomock.Controller
	mockRepo       *mocks.MockRepository
	listRemote     List
}

func (suite *ListRemoteTestSuite) SetupTest() {
	r, err := git2go.OpenRepository(os.Getenv("GITCONVEX_TEST_REPO"))
	if err != nil {
		fmt.Println(err)
	}

	suite.mockController = gomock.NewController(suite.T())
	suite.mockRepo = mocks.NewMockRepository(suite.mockController)
	suite.repo = middleware.NewRepository(r)
	suite.listRemote = NewRemoteList(suite.repo)
}

func TestListRemoteTestSuite(t *testing.T) {
	suite.Run(t, new(ListRemoteTestSuite))
}

func (suite *ListRemoteTestSuite) TestGetAllRemotes_WhenRepoIsValid_ShouldReturnAllRemotes() {
	expectedRemotes := []*model.RemoteDetails{{
		RemoteName: "origin",
		RemoteURL:  "https://github.com/neel1996/gitconvex-test.git",
	}}

	remoteList := suite.listRemote.GetAllRemotes()

	suite.Equal(len(expectedRemotes), len(remoteList))
	suite.Equal(expectedRemotes[0].RemoteName, remoteList[0].RemoteName)
	suite.Equal(expectedRemotes[0].RemoteURL, remoteList[0].RemoteURL)
}

func (suite *ListRemoteTestSuite) TestGetAllRemotes_WhenRepoIsNil_ShouldReturnNil() {
	suite.listRemote = NewRemoteList(nil)

	remoteList := suite.listRemote.GetAllRemotes()

	suite.Nil(remoteList)
}

func (suite *ListRemoteTestSuite) TestGetAllRemotes_WhenRepoHasNoRemotes_ShouldReturnNil() {
	suite.listRemote = NewRemoteList(suite.mockRepo)

	remoteList := suite.listRemote.GetAllRemotes()

	suite.Nil(remoteList)
}
