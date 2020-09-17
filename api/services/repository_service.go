package services

import (
	"strings"

	"github.com/marcio-garcia/go-microservices/api/config"
	"github.com/marcio-garcia/go-microservices/api/domain/github"
	"github.com/marcio-garcia/go-microservices/api/providers/githubprovider"

	"github.com/marcio-garcia/go-microservices/api/domain/repositories"
	"github.com/marcio-garcia/go-microservices/api/utils/errors"
)

type repoServiceInterface interface {
	CreateRepo(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.APIError)
}

// RepositoryService manages the repositories throught different providers
var (
	RepositoryService repoServiceInterface
)

func init() {
	RepositoryService = &repoService{}
}

type repoService struct{}

func (rs *repoService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.APIError) {
	input.Name = strings.TrimSpace(input.Name)
	if input.Name == "" {
		return nil, errors.CreateBadRequestError("Invalid repository name")
	}

	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
	}

	response, err := githubprovider.CreateRepo(config.GetGithubAccessToken(), request)

	if err != nil {
		return nil, errors.CreateError(err.StatusCode, err.Message)
	}

	result := repositories.CreateRepoResponse{
		ID:    response.ID,
		Name:  response.Name,
		Owner: response.Owner.Login,
	}

	return &result, nil
}
