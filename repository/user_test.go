package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-api-template/model/commonerrors"
	"go-api-template/pkg/random"
	"go-api-template/repository/model"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-openapi/strfmt"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/suite"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

type UserTestSuite struct {
	suite.Suite

	db          *sqlx.DB
	dbMock      sqlmock.Sqlmock
	repository  IUser
	userID      strfmt.UUID4
	failedError error
}

func (suite *UserTestSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)
}

func (suite *UserTestSuite) SetupTest() {
	db, dbMock, _ := sqlmock.New()
	suite.dbMock = dbMock
	suite.db = sqlx.NewDb(db, "sqlmock")
	suite.repository = NewUser(suite.db)
	suite.userID = random.UUID4()
	suite.failedError = errors.New("failed")
}

func (suite *UserTestSuite) TearDownTest() {
	suite.db.Close()
}

func Test_User_TestSuite(t *testing.T) {
	suite.Run(t, &UserTestSuite{})
}

func (suite *UserTestSuite) Test_Begin_ReturnsError_InCaseOfBeginFailed() {
	// Arrange
	suite.dbMock.ExpectBegin().WillReturnError(errors.New("failed"))

	// Act
	_, err := suite.repository.Begin()

	// Assert
	suite.Error(err)
}

func (suite *UserTestSuite) Test_Begin_ReturnsTransaction_InCaseOfSuccess() {
	// Arrange
	suite.dbMock.ExpectBegin().WillReturnError(nil)

	// Act
	tx, err := suite.repository.Begin()

	// Assert
	suite.NoError(err)
	suite.NotNil(tx)
}

func (suite *UserTestSuite) Test_SelectUserByFilter_ReturnsError_InCaseOfSelectFailed() {
	// Arrange
	usersFilter := model.UsersFilter{
		ID: &suite.userID,
	}

	suite.dbMock.ExpectQuery("SELECT").WillReturnError(errors.New("failed"))

	// Act
	_, err := suite.repository.SelectUserByFilter(context.Background(), usersFilter)

	// Assert
	suite.Error(err)
}

func (suite *UserTestSuite) Test_SelectUserByFilter_ReturnsError_InCaseOfUserNotFound() {
	// Arrange
	usersFilter := model.UsersFilter{
		ID: &suite.userID,
	}

	suite.dbMock.ExpectQuery("SELECT").WillReturnError(sql.ErrNoRows)

	// Act
	_, err := suite.repository.SelectUserByFilter(context.Background(), usersFilter)

	// Assert
	suite.Error(err)
	suite.Equal(commonerrors.ErrUserNotFound, err)
}

func (suite *UserTestSuite) Test_SelectUserByFilter_ReturnsUser_InCaseOfSuccess() {
	// Arrange
	usersFilter := model.UsersFilter{
		ID: &suite.userID,
	}

	rows := sqlmock.NewRows([]string{"id"}).AddRow(string(suite.userID))
	suite.dbMock.ExpectQuery("SELECT").WillReturnRows(rows)

	// Act
	user, err := suite.repository.SelectUserByFilter(context.Background(), usersFilter)

	// Assert
	suite.NoError(err)
	suite.NotNil(user)
	suite.Equal(suite.userID, user.ID)
}
