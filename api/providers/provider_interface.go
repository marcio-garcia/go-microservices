package providers

import (
	"github.com/marcio-garcia/go-microservices/api/domain/repositories"
	"github.com/marcio-garcia/go-microservices/api/utils/errors"
)

// ProviderInterface declares the interface for the repository API provider
type ProviderInterface interface {
	CreateRepo(accessToken string, request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, *errors.APIError)
}
