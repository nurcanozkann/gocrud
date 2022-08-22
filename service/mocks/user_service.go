package mocks

import (
	"cihanozhan.com/dbgo/common"
	"cihanozhan.com/dbgo/models"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

// CreateUser implements UserRepository
func (_m *MockUserService) CreateUser(user *models.User) (userResult *models.User, err *common.RestErr) {
	result := _m.Called(user)

	if result.Get(0) != nil {
		userResult = result.Get(0).(*models.User)
	}

	if result.Get(1) != nil {
		err = result.Get(1).(*common.RestErr)
	}

	return
}

func (_m *MockUserService) UpdateUser(id int, user *models.User) (userResult *models.User, err *common.RestErr) {
	result := _m.Called(id, user)

	if result.Get(0) != nil {
		userResult = result.Get(0).(*models.User)
	}

	if result.Get(1) != nil {
		err = result.Get(1).(*common.RestErr)
	}

	return
}

func (_m *MockUserService) DeleteUser(id int) (err *common.RestErr) {
	result := _m.Called(id)

	if result.Get(0) != nil {
		err = result.Get(1).(*common.RestErr)
	}

	return
}

func (_m *MockUserService) GetAllUsers() (users []models.User, err *common.RestErr) {
	result := _m.Called()

	if result.Get(0) != nil {
		users = result.Get(0).([]models.User)
	}

	if result.Get(1) != nil {
		err = result.Get(1).(*common.RestErr)
	}

	return
}

func (_m *MockUserService) GetUserById(id int) (user *models.User, err *common.RestErr) {
	result := _m.Called(id)

	if result.Get(0) != nil {
		user = result.Get(0).(*models.User)
	}

	if result.Get(1) != nil {
		err = result.Get(1).(*common.RestErr)
	}

	return
}
