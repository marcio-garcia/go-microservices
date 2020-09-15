package services

import (
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/marcio-garcia/go-microservices/mvc/domain"

	"github.com/stretchr/testify/assert"
)

var (
	users = map[uint64]*domain.User{
		1: {ID: 1, FirstName: "Walter", LastName: "White", Email: "ww@bb.com"},
	}
	UserDAO userDAOMock
)

func init() {
	domain.UserDAO = &userDAOMock{}
}

type userDAOMock struct{}

func (ud *userDAOMock) GetUser(ID uint64) (*domain.User, *domain.AppError) {
	log.Println("Getting the user from the mock")

	if user := users[ID]; user != nil {
		return user, nil
	}

	appError := domain.AppError{
		Message:    fmt.Sprintf("user %v not found", ID),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
	return nil, &appError
}

func TestGetUserNotFoundInDB(t *testing.T) {
	user, err := UserService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user 0 not found", err.Message)
	assert.EqualValues(t, "not_found", err.Code)
}

func TestGetUserSuccess(t *testing.T) {
	user, err := UserService.GetUser(1)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 1, user.ID)
	assert.EqualValues(t, "Walter", user.FirstName)
	assert.EqualValues(t, "White", user.LastName)
	assert.EqualValues(t, "ww@bb.com", user.Email)
}
