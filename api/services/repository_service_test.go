package services

import (
	"testing"

	"github.com/marcio-garcia/go-microservices/api/providers/githubprovider"

	"github.com/stretchr/testify/assert"

	"github.com/marcio-garcia/go-microservices/api/domain/repositories"
	"github.com/marcio-garcia/go-microservices/api/utils/errors"
)

func TestCreateRepoInvalidInput(t *testing.T) {
	request := repositories.CreateRepoRequest{
		Name:        "",
		Description: "",
	}
	result, err := RepositoryService.CreateRepo(request)
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, 400, err.Status())
	assert.EqualValues(t, "Invalid repository name", err.Message())
	assert.EqualValues(t, "", err.Error())
}

func TestCreateRepoProviderError(t *testing.T) {
	githubprovider.GithubProvider = &githubProviderErrorMock{}
	request := repositories.CreateRepoRequest{
		Name:        "go-example",
		Description: "Example description",
	}
	result, err := RepositoryService.CreateRepo(request)
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualValues(t, 400, err.Status())
	assert.EqualValues(t, "Bad request error", err.Message())
}

func TestCreateRepoSuccess(t *testing.T) {
	githubprovider.GithubProvider = &githubProviderSuccessMock{}
	request := repositories.CreateRepoRequest{
		Name:        "go-example",
		Description: "Example description",
	}
	result, err := RepositoryService.CreateRepo(request)
	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.EqualValues(t, 1, result.ID)
	assert.EqualValues(t, "go-example", result.Name)
	assert.EqualValues(t, "test-user", result.Owner)
}

//************
//**** Mockups
//************

type githubProviderErrorMock struct{}

func (p *githubProviderErrorMock) CreateRepo(accessToken string, request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, *errors.APIError) {
	err := errors.CreateBadRequestError("Bad request error")
	return nil, &err
}

type githubProviderSuccessMock struct{}

func (p *githubProviderSuccessMock) CreateRepo(accessToken string, request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, *errors.APIError) {
	response := repositories.CreateRepoResponse{
		ID:    1,
		Name:  "go-example",
		Owner: "test-user",
	}
	return &response, nil
}
