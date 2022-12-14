package repository

import (
	"cihanozhan.com/dbgo/common"
	"cihanozhan.com/dbgo/config"
	"cihanozhan.com/dbgo/models"
	"github.com/jinzhu/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

//go:generate mockery --name=UserRepository
type UserRepository interface {
	Create(user *models.User) (*models.User, *common.RestErr)
	Update(id int, user *models.User) (*models.User, *common.RestErr)
	GetByEmail(email string) (*models.User, *common.RestErr)
	GetById(id int) (*models.User, *common.RestErr)
	Delete(id int) *common.RestErr
	GetAllUsers() ([]models.User, *common.RestErr)
}

func NewUserRepository() UserRepository {
	return &userRepository{DB: config.ConnDB()}
}

func (repo *userRepository) Create(user *models.User) (*models.User, *common.RestErr) {
	existingUser, _ := repo.GetByEmail(user.Email)
	if existingUser.Email == user.Email {
		restErr := common.Forbidden("User already exists")
		return nil, restErr
	}

	result := config.DB.Create(&user)

	if result.Error != nil {
		restErr := common.InternalErr("server error.")
		return nil, restErr
	}
	return user, nil
}

func (repo *userRepository) Update(id int, user *models.User) (*models.User, *common.RestErr) {
	if user.Email == "" {
		restErr := common.NotFound("User email is empty.")
		return user, restErr
	}
	if user.Name == "" {
		restErr := common.NotFound("User name is empty.")
		return user, restErr
	}
	if user.Password == "" {
		restErr := common.NotFound("User password is empty.")
		return user, restErr
	}

	userEntity, restErr := repo.GetById(id)
	if restErr != nil {
		restErr := common.NotFound("User with that id does not exist.")
		return nil, restErr
	}

	userEntity.Name = user.Name
	userEntity.Email = user.Email
	userEntity.Password = user.Password

	result := config.DB.Save(userEntity)
	if result.Error != nil {
		restErr := common.InternalErr("server error.")
		return nil, restErr
	}
	return user, nil
}

func (repo *userRepository) GetByEmail(email string) (*models.User, *common.RestErr) {
	var getUser models.User
	res := config.DB.Where("email = ?", email).Take(&getUser)
	if res.Error != nil {
		restErr := common.NotFound("User Not Found.")
		return &getUser, restErr
	}
	return &getUser, nil
}

func (repo *userRepository) GetById(id int) (*models.User, *common.RestErr) {
	var user models.User
	res := config.DB.Where("id = ?", id).Take(&user)
	if res.Error != nil {
		restErr := common.NotFound("User Not Found.")
		return &user, restErr
	}
	return &user, nil
}

func (repo *userRepository) Delete(id int) *common.RestErr {
	userEntity, restErr := repo.GetById(id)
	if restErr != nil {
		restErr := common.NotFound("User with that id does not exist.")
		return restErr
	}

	result := config.DB.Delete(&userEntity)

	if result.Error != nil {
		restErr := common.InternalErr("server error.")
		return restErr
	}

	return nil
}

func (repo *userRepository) GetAllUsers() ([]models.User, *common.RestErr) {
	users := []models.User{}
	result := config.DB.Find(&users)
	if result.Error != nil {
		restErr := common.NotFound("Users Not Found.")
		return nil, restErr
	}

	return users, nil
}
