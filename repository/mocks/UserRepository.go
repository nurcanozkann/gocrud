// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	"cihanozhan.com/dbgo/common"
	models "cihanozhan.com/dbgo/models"

	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type MockUserRepository struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: user
func (_m *MockUserRepository) Create(user *models.User) (userResult *models.User, err *common.RestErr) {
	result := _m.Called(user)

	if result.Get(0) != nil {
		userResult = result.Get(0).(*models.User)
	}

	if result.Get(1) != nil {
		err = result.Get(1).(*common.RestErr)
	}

	return
}

// DeleteUser provides a mock function with given fields: delUserEntity
func (_m *MockUserRepository) Delete(id int) (err *common.RestErr)  {
	result := _m.Called(id)

	
	if result.Get(0) != nil {
		err = result.Get(1).(*common.RestErr)
	}

	return
}

// GetAll provides a mock function with given fields:
func (_m *MockUserRepository) GetAllUsers() (userResult []models.User, err *common.RestErr) {
	result := _m.Called()

	if result.Get(0) != nil {
		userResult = result.Get(0).([]models.User)
	}

	if result.Get(1) != nil {
		err = result.Get(1).(*common.RestErr)
	}

	return
}

// GetByEmail provides a mock function with given fields: email
func (_m *MockUserRepository) GetByEmail(email string) (userResult *models.User, err *common.RestErr) {
	result := _m.Called(email)

	if result.Get(0) != nil {
		userResult = result.Get(0).(*models.User)
	}

	if result.Get(1) != nil {
		err = result.Get(1).(*common.RestErr)
	}

	return
}

// GetById provides a mock function with given fields: id
func (_m *MockUserRepository) GetById(id int) (user *models.User, err *common.RestErr) {
	result := _m.Called(id)

	if result.Get(0) != nil {
		user = result.Get(0).(*models.User)
	} 

	if result.Get(1) != nil {
		err = result.Get(1).(*common.RestErr)
	}

	return
}

// UpdateUser provides a mock function with given fields: updateUserEntity
func (_m *MockUserRepository) Update(id int, user *models.User) (userResult *models.User, err *common.RestErr) {
	result := _m.Called(id, user)

	if result.Get(0) != nil {
		userResult = result.Get(0).(*models.User)
	}

	if result.Get(1) != nil {
		err = result.Get(1).(*common.RestErr)
	}

	return
}