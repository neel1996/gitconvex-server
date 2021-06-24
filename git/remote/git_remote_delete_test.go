package remote

import (
	"fmt"
	git2go "github.com/libgit2/git2go/v31"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type RemoteDeleteTestSuite struct {
	suite.Suite
	deleteRemote Delete
}

func TestRemoteDeleteTestSuite(t *testing.T) {
	suite.Run(t, new(RemoteDeleteTestSuite))
}

func (suite *RemoteDeleteTestSuite) SetupTest() {
	r, err := git2go.OpenRepository(os.Getenv("GITCONVEX_TEST_REPO"))
	if err != nil {
		fmt.Println(err)
	}
	suite.deleteRemote = NewDeleteRemote(r, "new_origin")
}

func (suite *RemoteDeleteTestSuite) TestDeleteNewRemote_WhenNewRemoteIsDeleted_ShouldReturnNoError() {
	err := suite.deleteRemote.DeleteRemote()

	suite.Nil(err)
}

func (suite *RemoteDeleteTestSuite) TestDeleteNewRemote_WhenRequiredFieldsAreEmpty_ShouldReturnError() {
	suite.deleteRemote = NewDeleteRemote(nil, "")

	err := suite.deleteRemote.DeleteRemote()

	suite.NotNil(err)
}

func (suite *RemoteDeleteTestSuite) TestDeleteNewRemote_WhenRemoteDeletionFails_ShouldReturnError() {
	r, _ := git2go.OpenRepository(os.Getenv("GITCONVEX_TEST_REPO"))

	suite.deleteRemote = NewDeleteRemote(r, "new_origin")

	err := suite.deleteRemote.DeleteRemote()

	suite.NotNil(err)
}