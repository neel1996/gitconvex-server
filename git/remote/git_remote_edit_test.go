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

type RemoteEditTestSuite struct {
	suite.Suite
	repo           middleware.Repository
	mockController *gomock.Controller
	mockRepo       *mocks.MockRepository
	remoteName     string
	remoteUrl      string
	validation     Validation
	editRemote     Edit
}

func TestRemoteEditTestSuite(t *testing.T) {
	suite.Run(t, new(RemoteEditTestSuite))
}

func (suite *RemoteEditTestSuite) SetupTest() {
	r, err := git2go.OpenRepository(os.Getenv("GITCONVEX_TEST_REPO"))
	if err != nil {
		fmt.Println(err)
	}

	suite.repo = middleware.NewRepository(r)
	suite.mockController = gomock.NewController(suite.T())
	suite.mockRepo = mocks.NewMockRepository(suite.mockController)
	suite.remoteName = "origin"
	suite.remoteUrl = "https://github.com/neel1996/gitconvex-test.git"
	suite.editRemote = NewEditRemote(suite.mockRepo, suite.remoteName, suite.remoteUrl)
}

func (suite *RemoteEditTestSuite) TestEditRemote_WhenRemoteIsEdited_ShouldReturnNil() {
	suite.editRemote = NewEditRemote(suite.repo, suite.remoteName, suite.remoteUrl)

	wantErr := suite.editRemote.EditRemote()

	suite.Nil(wantErr)
}

func (suite *RemoteEditTestSuite) TestEditRemote_WhenRepoIsNil_ShouldReturnError() {
	suite.editRemote = NewEditRemote(nil, suite.remoteName, suite.remoteUrl)

	wantErr := suite.editRemote.EditRemote()

	suite.NotNil(wantErr)
}

func (suite *RemoteEditTestSuite) TestEditRemote_WhenRemoteCollectionIsNil_ShouldReturnError() {
	suite.editRemote = NewEditRemote(suite.mockRepo, suite.remoteName, suite.remoteUrl)

	wantErr := suite.editRemote.EditRemote()

	suite.NotNil(wantErr)
}

func (suite *RemoteEditTestSuite) TestEditRemote_WhenRemoteNameIsEmpty_ShouldReturnError() {
	suite.editRemote = NewEditRemote(suite.mockRepo, "", suite.remoteUrl)

	wantErr := suite.editRemote.EditRemote()

	suite.NotNil(wantErr)
}

func (suite *RemoteEditTestSuite) TestEditRemote_WhenRemoteUrlIsEmpty_ShouldReturnError() {
	suite.editRemote = NewEditRemote(suite.mockRepo, suite.remoteName, "")

	wantErr := suite.editRemote.EditRemote()

	suite.NotNil(wantErr)
}

func (suite *RemoteEditTestSuite) TestEditRemote_WhenRepoHasNoRemotes_ShouldReturnError() {
	suite.editRemote = NewEditRemote(suite.mockRepo, suite.remoteName, suite.remoteUrl)

	wantErr := suite.editRemote.EditRemote()

	suite.NotNil(wantErr)
}

func (suite *RemoteEditTestSuite) TestEditRemote_WhenRemoteIsNotPresent_ShouldReturnError() {
	suite.editRemote = NewEditRemote(suite.mockRepo, "no_exists_remote", suite.remoteUrl)

	wantErr := suite.editRemote.EditRemote()

	suite.NotNil(wantErr)
}
