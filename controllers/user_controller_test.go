package controllers

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"cihanozhan.com/dbgo/common"
	"cihanozhan.com/dbgo/models"
	"cihanozhan.com/dbgo/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetUserById(t *testing.T) {
	mockUserService := &mocks.MockUserService{}
	userController := &userController{
		userService: mockUserService,
	}

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)

	ctx.Params = append(ctx.Params, gin.Param{
		Key:   "id",
		Value: "1",
	})

	expectedUser := &models.User{
		Name: "Nurcan",
	}

	mockUserService.On("GetUserById", 1).Return(expectedUser, nil)
	userController.GetUserById(ctx)

	response := &models.User{}
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, expectedUser, response)
}

func TestGetUserByIdWhenIdParamDoesNotExist(t *testing.T) {
	mockUserService := &mocks.MockUserService{}
	userController := &userController{
		userService: mockUserService,
	}

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)

	userController.GetUserById(ctx)
	mockUserService.AssertNumberOfCalls(t, "GetUserById", 0)
	mockUserService.AssertNotCalled(t, "GetUserById", mock.Anything)

	response := &common.RestErr{}
	err := json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, 400, response.Status)
	assert.Equal(t, "Bad request", response.Error)
}
