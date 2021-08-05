package remote

import (
	"fmt"
	git2go "github.com/libgit2/git2go/v31"
	"github.com/neel1996/gitconvex/git/middleware"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type RemoteDeleteTestSuite struct {
	suite.Suite
	repo         middleware.Repository
	remoteName   string
	deleteRemote Delete
}

func TestRemoteDeleteTestSuite(t *testing.T) {
	suite.Run(t, new(RemoteDeleteTestSuite))
}

func (suite *RemoteDeleteTestSuite) SetupSuite() {
	r, err := git2go.OpenRepository(os.Getenv("GITCONVEX_TEST_REPO"))
	if err != nil {
		fmt.Println(err)
	}
	suite.remoteName = "new_origin"
	_ = NewAddRemote(middleware.NewRepository(r), suite.remoteName, "remote://some_url").NewRemote()
}

func (suite *RemoteDeleteTestSuite) SetupTest() {
	r, err := git2go.OpenRepository(os.Getenv("GITCONVEX_TEST_REPO"))
	if err != nil {
		fmt.Println(err)
	}
	suite.repo = middleware.NewRepository(r)
	suite.remoteName = "new_origin"
	suite.deleteRemote = NewDeleteRemote(suite.repo, suite.remoteName)
}

func (suite *RemoteDeleteTestSuite) TestDeleteNewRemote_WhenNewRemoteIsDeleted_ShouldReturnNoError() {
	err := suite.deleteRemote.DeleteRemote()

	suite.Nil(err)
}

func (suite *RemoteDeleteTestSuite) TestDeleteNewRemote_WhenRepoIsNil_ShouldReturnError() {
	suite.deleteRemote = NewDeleteRemote(nil, suite.remoteName)

	err := suite.deleteRemote.DeleteRemote()

	suite.NotNil(err)
}

func (suite *RemoteDeleteTestSuite) TestDeleteNewRemote_WhenRemoteNameIsEmpty_ShouldReturnError() {
	suite.deleteRemote = NewDeleteRemote(suite.repo, "")

	err := suite.deleteRemote.DeleteRemote()

	suite.NotNil(err)
}

func (suite *RemoteDeleteTestSuite) TestDeleteNewRemote_WhenRemoteDeletionFails_ShouldReturnError() {
	suite.deleteRemote = NewDeleteRemote(suite.repo, "new_origin")

	err := suite.deleteRemote.DeleteRemote()

	suite.NotNil(err)
}
