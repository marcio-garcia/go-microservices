package domain

import (
	"fmt"
	"net/http"
)

var (
	users = map[uint64]*User{
		1: {1, "Waler", "White", "ww@bb.com"},
	}
)

// GetUser fetch user from database
func GetUser(ID uint64) (*User, *AppError) {
	user := users[ID]
	if user == nil {
		appError := AppError{
			Message:    fmt.Sprintf("user %v not found", ID),
			StatusCode: http.StatusNotFound,
			Code:       "not_found",
		}
		return nil, &appError
	}
	return user, nil
}
