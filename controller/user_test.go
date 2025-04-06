package controller

import (
	"encoding/json"
	"go-api-template/model"
	"go-api-template/model/commonerrors"
	"go-api-template/pkg/random"
	"go-api-template/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserTestSuite struct {
	suite.Suite

	ctx              *gin.Context
	responseRecorder *httptest.ResponseRecorder
	controller       IUser
	userServiceMock  *service.UserMock
	userID           strfmt.UUID4
}

func (suite *UserTestSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)
}

func (suite *UserTestSuite) SetupTest() {
	suite.responseRecorder = httptest.NewRecorder()
	suite.ctx, _ = gin.CreateTestContext(suite.responseRecorder)
	suite.userServiceMock = &service.UserMock{}
	suite.controller = NewUser(suite.userServiceMock)
	suite.userID = random.UUID4()
}

func (suite *UserTestSuite) TearDownTest() {
	suite.userServiceMock.AssertExpectations(suite.T())
}

func Test_User_TestSuite(t *testing.T) {
	t.Parallel() // Enable parallel execution
	suite.Run(t, &UserTestSuite{})
}

func (suite *UserTestSuite) Test_UserByID_ReturnsBadRequest_InCaseOfUserIDIsNotValid() {
	// Arrange
	userID := random.String(10)

	suite.ctx.Request = httptest.NewRequest(http.MethodGet, "/users/"+userID, nil)
	suite.ctx.Params = append(suite.ctx.Params, gin.Param{Key: "user_id", Value: userID})

	// Act
	suite.controller.UserByID(suite.ctx)

	// Assert
	suite.Require().Equal(http.StatusBadRequest, suite.responseRecorder.Code)

	var actualResponse model.BadRequestResponse
	err := json.NewDecoder(suite.responseRecorder.Body).Decode(&actualResponse)
	suite.Require().NoError(err)
	suite.Equal(commonerrors.ErrInvalidUserID.Error(), actualResponse.Message)
}

func (suite *UserTestSuite) Test_UserByID_ReturnsUnprocessableEntity_InCaseOfUserNotFound() {
	// Arrange
	suite.userServiceMock.
		On("UserByID", mock.Anything, suite.userID).
		Return(nil, commonerrors.ErrUserNotFound).
		Once()

	suite.ctx.Request = httptest.NewRequest(http.MethodGet, "/users/"+suite.userID.String(), nil)
	suite.ctx.Params = append(suite.ctx.Params, gin.Param{Key: "user_id", Value: suite.userID.String()})

	// Act
	suite.controller.UserByID(suite.ctx)

	// Assert
	suite.Require().Equal(http.StatusUnprocessableEntity, suite.responseRecorder.Code)

	var actualResponse model.UnprocessableEntityResponse
	err := json.NewDecoder(suite.responseRecorder.Body).Decode(&actualResponse)
	suite.Require().NoError(err)
	suite.Equal(commonerrors.ErrUserNotFound.Error(), actualResponse.Message)
}

func (suite *UserTestSuite) Test_UserByID_ReturnsResponse_InCaseOfSuccess() {
	// Arrange
	expectedResponse := &model.UserByIDResponse{
		User: &model.User{
			ID: suite.userID,
		},
	}

	suite.userServiceMock.
		On("UserByID", mock.Anything, suite.userID).
		Return(expectedResponse, nil).
		Once()

	suite.ctx.Request = httptest.NewRequest(http.MethodGet, "/users/"+suite.userID.String(), nil)
	suite.ctx.Params = append(suite.ctx.Params, gin.Param{Key: "user_id", Value: suite.userID.String()})

	// Act
	suite.controller.UserByID(suite.ctx)

	// Assert
	suite.Require().Equal(http.StatusOK, suite.responseRecorder.Code)

	var actualResponse model.UserByIDResponse
	err := json.NewDecoder(suite.responseRecorder.Body).Decode(&actualResponse)
	suite.Require().NoError(err)
	suite.Equal(expectedResponse, &actualResponse)
}
