package service

import (
	"cihanozhan.com/dbgo/common"
	"cihanozhan.com/dbgo/models"
	"cihanozhan.com/dbgo/repository"
	_ "github.com/lib/pq"
)

type userService struct {
	userRepository repository.UserRepository
}

type UserService interface {
	CreateUser(user *models.User) (*models.User, *common.RestErr)
	UpdateUser(id int, user *models.User) (*models.User, *common.RestErr)
	DeleteUser(id int) *common.RestErr
	GetAllUsers() ([]models.User, *common.RestErr)
	GetUserById(id int) (*models.User, *common.RestErr)
}

func NewUserService() UserService {
	return &userService{userRepository: repository.NewUserRepository()}
}

// CreateUser implements UserRepository
func (service *userService) CreateUser(user *models.User) (*models.User, *common.RestErr) {
	user, restErr := service.userRepository.Create(user)
	if restErr != nil {
		return nil, restErr
	}
	return user, nil
}

func (service *userService) UpdateUser(id int, user *models.User) (*models.User, *common.RestErr) {
	updateUser, restErr := service.userRepository.Update(id, user)
	if restErr != nil {
		return nil, restErr
	}
	return updateUser, nil
}

func (service *userService) DeleteUser(id int) *common.RestErr {
	restErr := service.userRepository.Delete(id)
	if restErr != nil {
		return restErr
	}
	return nil
}

func (service *userService) GetAllUsers() ([]models.User, *common.RestErr) {
	users, restErr := service.userRepository.GetAllUsers()
	if restErr != nil {
		return nil, restErr
	}

	return users, nil
}

func (service *userService) GetUserById(id int) (*models.User, *common.RestErr) {
	user, restErr := service.userRepository.GetById(id)
	if restErr != nil {
		return nil, restErr
	}

	return user, nil
}
