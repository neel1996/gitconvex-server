package branch

import (
	"github.com/golang/mock/gomock"
	"github.com/neel1996/gitconvex/mocks"
	"github.com/stretchr/testify/suite"
	"testing"
)

type BranchFieldValidationTestSuite struct {
	suite.Suite
	mockController *gomock.Controller
	mockRepo       *mocks.MockRepository
	branchNames    []string
	validation     Validation
}

func TestBranchFieldValidationTestSuite(t *testing.T) {
	suite.Run(t, new(BranchFieldValidationTestSuite))
}

func (suite *BranchFieldValidationTestSuite) SetupTest() {
	suite.mockController = gomock.NewController(suite.T())
	suite.mockRepo = mocks.NewMockRepository(suite.mockController)
	suite.branchNames = []string{"test_branch_1", "test_branch_2"}
	suite.validation = NewBranchFieldsValidation(suite.mockRepo, suite.branchNames[0], suite.branchNames[1])
}

func (suite *BranchFieldValidationTestSuite) TestValidateBranchFields_WhenAllFieldsAreValid_ShouldReturnNil() {
	err := suite.validation.ValidateBranchFields()

	suite.Nil(err)
}

func (suite *BranchFieldValidationTestSuite) TestValidateBranchFields_WhenRepoIsNil_ShouldReturnError() {
	suite.validation = NewBranchFieldsValidation(nil, "test_branch")
	err := suite.validation.ValidateBranchFields()

	suite.NotNil(err)
	suite.Equal(NilRepoError, err)
}

func (suite *BranchFieldValidationTestSuite) TestValidateBranchFields_WhenBranchNameIsEmpty_ShouldReturnError() {
	suite.validation = NewBranchFieldsValidation(suite.mockRepo, "")

	err := suite.validation.ValidateBranchFields()

	suite.NotNil(err)
	suite.Equal(EmptyBranchNameError, err)
}
