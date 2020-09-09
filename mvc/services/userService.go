package services

import "github.com/marcio-garcia/go-microservices/mvc/domain"

// GetUser returns a user by its ID
func GetUser(ID uint64) (*domain.User, *domain.AppError) {
	return domain.GetUser(ID)
}
