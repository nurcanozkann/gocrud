package service

import (
	"testing"

	"cihanozhan.com/dbgo/models"
	"cihanozhan.com/dbgo/repository/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetUserById(t *testing.T) {
	mockUserRepository := &mocks.MockUserRepository{}
	userService := &userService{
		userRepository: mockUserRepository,
	}

	expectedUser := &models.User{}

	mockUserRepository.On("GetById", 1).Return(expectedUser, nil)
	actualUser, err := userService.GetUserById(1)
	assert.Equal(t, expectedUser, actualUser)
	assert.Nil(t, err, "err must be nil")
}

func TestCreateUser(t *testing.T) {
	mockUserRepository := &mocks.MockUserRepository{}
	userService := &userService{
		userRepository: mockUserRepository,
	}

	user := &models.User{
		Name:     "Ela",
		Email:    "ela@gmail.com",
		Password: "ela123",
	}

	expectedUser := &models.User{
		Name:     "Ela",
		Email:    "ela@gmail.com",
		Password: "ela123",
	}
	expectedUser.ID = 1

	mockUserRepository.On("Create", user).Return(expectedUser, nil)
	savedUser, err := userService.CreateUser(user)
	assert.Equal(t, expectedUser, savedUser)
	assert.Nil(t, err, "err must be nil")
}

func TestUpdateUser(t *testing.T) {
	mockUserRepository := &mocks.MockUserRepository{}
	userService := &userService{
		userRepository: mockUserRepository,
	}

	user := &models.User{
		Name:     "Ela",
		Email:    "ela@gmail.com",
		Password: "ela123",
	}

	expectedUser := &models.User{
		Name:     "Ela",
		Email:    "ela@gmail.com",
		Password: "ela123",
	}
	expectedUser.ID = 1

	mockUserRepository.On("Update", 1, user).Return(expectedUser, nil)
	savedUser, err := userService.UpdateUser(1, user)
	assert.Equal(t, expectedUser, savedUser)
	assert.Nil(t, err, "err must be nil")
}

func TestDeleteUser(t *testing.T) {
	mockUserRepository := &mocks.MockUserRepository{}
	userService := &userService{
		userRepository: mockUserRepository,
	}

	mockUserRepository.On("Delete", 1).Return(nil)
	err := userService.DeleteUser(1)
	assert.Nil(t, err, "err must be nil")
}

func TestGetAllUsers(t *testing.T) {
	mockUserRepository := &mocks.MockUserRepository{}
	userService := &userService{
		userRepository: mockUserRepository,
	}

	expectedUsers := []models.User{}

	mockUserRepository.On("GetAllUsers").Return(expectedUsers, nil)
	actualUsers, err := userService.GetAllUsers()
	assert.Equal(t, expectedUsers, actualUsers)
	assert.Nil(t, err, "err must be nil")
}
