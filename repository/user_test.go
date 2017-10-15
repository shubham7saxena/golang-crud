package repository

import (
	"crud/appcontext"
	"crud/config"
	"crud/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type userTestSuite struct {
	suite.Suite
	repository *userRepository
}

func (suite *userTestSuite) SetupTest() {
	config.Load()
	appcontext.Initiate()
	suite.repository = NewUserRepository()
}

func (suite *userTestSuite) TearDownTest() {
	cleanupDB()
}

func (suite *userTestSuite) TestUserInsert() {
	userID := 1
	userAge := 20
	userName := "Saxena"
	testUser := newUser(userID, userAge, userName)
	err := suite.repository.Insert(testUser)
	assert.Nil(suite.T(), err)
}

func (suite *userTestSuite) TestGetUser() {
	userID := 1
	userAge := 20
	userName := "Saxena"
	testUser := newUser(userID, userAge, userName)
	_ = suite.repository.Insert(testUser)
	user, err := suite.repository.GetUser(testUser.Id)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), testUser, user)
}

func (suite *userTestSuite) TestDeleteUser() {
	userID := 1
	userAge := 20
	userName := "Saxena"
	testUser := newUser(userID, userAge, userName)
	_ = suite.repository.Insert(testUser)
	err := suite.repository.DeleteUser(testUser.Id)
	assert.Nil(suite.T(), err)
}

func (suite *userTestSuite) TestUserUpdate() {
	userID := 1
	userAge := 20
	userName := "Saxena"
	testUser := newUser(userID, userAge, userName)
	_ = suite.repository.Insert(testUser)
	testUserUpdated := newUser(userID, userAge, "Shubham")
	err := suite.repository.UpdateUser(testUserUpdated)
	assert.Nil(suite.T(), err)
}

func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(userTestSuite))
}

func newUser(userID int, age int, name string) *domain.User {
	return &domain.User{
		Id:   userID,
		Age:  age,
		Name: name,
	}
}

func cleanupDB() {
	db := appcontext.GetDB()
	db.Exec("truncate users")
	db.Exec("alter sequence users_id_seq restart with 1")
}
