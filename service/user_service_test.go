package service

import (
	"testing"
)

func TestGetAllUserss(t *testing.T) {
	_, err := GetAllUsers()

	if err != nil {
		t.Error(err.Message)
	}
}
