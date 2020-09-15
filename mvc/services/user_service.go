package services

import "github.com/marcio-garcia/go-microservices/mvc/domain"

type userService struct {
}

// UserService
var (
	UserService userService
)

// GetUser returns a user by its ID
func (us *userService) GetUser(ID uint64) (*domain.User, *domain.AppError) {
	user, err := domain.UserDAO.GetUser(ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
