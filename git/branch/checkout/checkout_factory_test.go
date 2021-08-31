package checkout

import (
	"github.com/golang/mock/gomock"
	"github.com/neel1996/gitconvex/git/branch"
	branchMocks "github.com/neel1996/gitconvex/git/branch/mocks"
	"github.com/neel1996/gitconvex/git/middleware"
	"github.com/neel1996/gitconvex/mocks"
	"github.com/stretchr/testify/suite"
	"testing"
)

type CheckoutFactoryTestSuite struct {
	suite.Suite
	repo                 middleware.Repository
	branchName           string
	remoteBranchName     string
	branchValidation     branch.Validation
	mockController       *gomock.Controller
	mockRepo             *mocks.MockRepository
	mockBranchValidation *branchMocks.MockValidation
	checkoutFactory      Factory
}

func TestCheckoutFactoryTestSuite(t *testing.T) {
	suite.Run(t, new(CheckoutFactoryTestSuite))
}

func (suite *CheckoutFactoryTestSuite) SetupTest() {
	suite.mockController = gomock.NewController(suite.T())
	suite.mockRepo = mocks.NewMockRepository(suite.mockController)
	suite.mockBranchValidation = branchMocks.NewMockValidation(suite.mockController)
	suite.branchName = "test_branch"
	suite.remoteBranchName = "remotes/origin/test_branch"
}

func (suite *CheckoutFactoryTestSuite) TearDownTest() {
	suite.mockController.Finish()
}

func (suite *CheckoutFactoryTestSuite) TestGetCheckoutAction_WhenBranchIsLocal_ShouldReturnLocalCheckoutAction() {
	suite.checkoutFactory = NewCheckoutFactory(suite.mockRepo, suite.branchName, suite.mockBranchValidation)

	suite.mockBranchValidation.EXPECT().ValidateBranchFields(suite.branchName).Return(nil)

	wantAction := NewCheckOutLocalBranch(suite.mockRepo, suite.branchName)
	gotAction := suite.checkoutFactory.GetCheckoutAction()

	suite.Equal(wantAction, gotAction)
}

func (suite *CheckoutFactoryTestSuite) TestGetCheckoutAction_WhenBranchIsRemote_ShouldReturnRemoteCheckoutAction() {
	suite.checkoutFactory = NewCheckoutFactory(suite.mockRepo, suite.remoteBranchName, suite.mockBranchValidation)

	suite.mockBranchValidation.EXPECT().ValidateBranchFields(suite.remoteBranchName).Return(nil)

	wantAction := NewCheckoutRemoteBranch(suite.mockRepo, suite.remoteBranchName)
	gotAction := suite.checkoutFactory.GetCheckoutAction()

	suite.Equal(wantAction, gotAction)
}

func (suite *CheckoutFactoryTestSuite) TestGetCheckoutAction_WhenBranchValidationFails_ShouldReturnNil() {
	suite.checkoutFactory = NewCheckoutFactory(suite.mockRepo, "", suite.mockBranchValidation)

	suite.mockBranchValidation.EXPECT().ValidateBranchFields("").Return(branch.EmptyBranchNameError)

	gotAction := suite.checkoutFactory.GetCheckoutAction()

	suite.Nil(gotAction)
}
