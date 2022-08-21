package service

import (
	"cihanozhan.com/dbgo/common"
	"cihanozhan.com/dbgo/models"
	"cihanozhan.com/dbgo/repository"
	_ "github.com/lib/pq"
)

func CreateUser(user *models.User) (*models.User, *common.RestErr) {
	user, restErr := repository.Create(user)
	if restErr != nil {
		return nil, restErr
	}
	return user, nil
}

func UpdateUser(id int, user *models.User) (*models.User, *common.RestErr) {
	updateUser, restErr := repository.Update(id, user)
	if restErr != nil {
		return nil, restErr
	}
	return updateUser, nil
}

func DeleteUser(id int) *common.RestErr {
	restErr := repository.Delete(id)
	if restErr != nil {
		return restErr
	}
	return nil
}

func GetAllUsers() ([]models.User, *common.RestErr) {
	users, restErr := repository.GetAllUsers()
	if restErr != nil {
		return nil, restErr
	}

	return users, nil
}

func GetUserById(id int) (*models.User, *common.RestErr) {
	user, restErr := repository.GetById(id)
	if restErr != nil {
		return nil, restErr
	}

	return user, nil
}
