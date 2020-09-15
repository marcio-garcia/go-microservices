package services

import (
	"net/http"

	"github.com/marcio-garcia/go-microservices/mvc/domain"
)

type itemService struct{}

// ItemService
var (
	ItemService itemService
)

// GetItem returns a item by its ID
func (is *itemService) GetItem(ID uint64) (*domain.Item, *domain.AppError) {
	return nil, &domain.AppError{
		Message:    "Not implemented",
		StatusCode: http.StatusBadRequest,
		Code:       "bad_request",
	}
}
