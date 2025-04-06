package service

import (
	"go-api-template/model/commonerrors"
	"go-api-template/pkg/random"
	"go-api-template/repository"
	repositorymodel "go-api-template/repository/model"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/suite"
)

type UserTestSuite struct {
	suite.Suite

	ctx                *gin.Context
	service            IUser
	userRepositoryMock *repository.UserMock
	userID             strfmt.UUID4
	failedError        error
}

func (suite *UserTestSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)
}

func (suite *UserTestSuite) SetupTest() {
	suite.ctx, _ = gin.CreateTestContext(nil)
	suite.userRepositoryMock = &repository.UserMock{}
	suite.service = NewUser(suite.userRepositoryMock)
	suite.userID = random.UUID4()
	suite.failedError = commonerrors.ErrFailed
}

func (suite *UserTestSuite) TearDownTest() {
	suite.userRepositoryMock.AssertExpectations(suite.T())
}

func Test_User_TestSuite(t *testing.T) {
	t.Parallel() // Enable parallel execution
	suite.Run(t, &UserTestSuite{})
}

func (suite *UserTestSuite) Test_UserByID_ReturnsError_InCaseOfSelectUserByFilterFailed() {
	// Arrange
	filter := repositorymodel.UsersFilter{
		ID: &suite.userID,
	}

	suite.userRepositoryMock.
		On("SelectUserByFilter", filter).
		Return(nil, suite.failedError).
		Once()

	// Act
	response, err := suite.service.UserByID(suite.ctx, suite.userID)

	// Assert
	suite.Require().Nil(response)

	if suite.Error(err) {
		suite.Equal(suite.failedError, err)
	}
}

func (suite *UserTestSuite) Test_UserByID_ReturnsError_InCaseOfUserNotFound() {
	// Arrange
	filter := repositorymodel.UsersFilter{
		ID: &suite.userID,
	}

	suite.userRepositoryMock.
		On("SelectUserByFilter", filter).
		Return(nil, commonerrors.ErrUserNotFound).
		Once()

	// Act
	response, err := suite.service.UserByID(suite.ctx, suite.userID)

	// Assert
	suite.Require().Nil(response)

	if suite.Error(err) {
		suite.Equal(commonerrors.ErrUserNotFound, err)
	}
}

func (suite *UserTestSuite) Test_UserByID_ReturnsUser_InCaseOfSuccess() {
	// Arrange
	user := &repositorymodel.User{
		ID: suite.userID,
	}

	filter := repositorymodel.UsersFilter{
		ID: &suite.userID,
	}

	suite.userRepositoryMock.
		On("SelectUserByFilter", filter).
		Return(user, nil).
		Once()

	// Act
	response, err := suite.service.UserByID(suite.ctx, suite.userID)

	// Assert
	suite.Require().NoError(err)

	if suite.NotNil(response) {
		suite.Equal(user.ID, response.User.ID)
	}
}
