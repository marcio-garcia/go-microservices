package domain

import (
	"fmt"
	"log"
	"net/http"
)

// UserDAO is the direct access object for the User model
var (
	users = map[uint64]*User{
		1: {1, "Walter", "White", "ww@bb.com"},
	}
	UserDAO userDAOInterface
)

func init() {
	UserDAO = &userDAO{}
}

type userDAOInterface interface {
	GetUser(uint64) (*User, *AppError)
}

type userDAO struct{}

// GetUser fetches the user from database by id
func (ud *userDAO) GetUser(ID uint64) (*User, *AppError) {

	log.Println("Getting the user from the database")

	if user := users[ID]; user != nil {
		return user, nil
	}

	appError := AppError{
		Message:    fmt.Sprintf("user %v not found", ID),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
	return nil, &appError
}
