package domain_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/marcio-garcia/go-microservices/mvc/domain"
	"github.com/stretchr/testify/assert"
)

func TestGetUserNotFound(t *testing.T) {
	var userID uint64 = 0
	errorMessage := fmt.Sprintf("user %v not found", userID)
	user, err := domain.UserDAO.GetUser(userID)

	assert.Nil(t, user, "we are not expecting an user with id 0")
	assert.NotNil(t, err, "we are expecting an error when user id is 0")
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, errorMessage, err.Message)
}

func TestGetUserSuccess(t *testing.T) {
	var userID uint64 = 1
	user, err := domain.UserDAO.GetUser(userID)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, userID, user.ID)
	assert.EqualValues(t, "Walter", user.FirstName)
	assert.EqualValues(t, "White", user.LastName)
	assert.EqualValues(t, "ww@bb.com", user.Email)
}
