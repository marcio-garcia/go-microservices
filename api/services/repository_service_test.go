package services

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/marcio-garcia/go-microservices/api/domain/repositories"
)

func TestCreateRepoInvalidInput(t *testing.T) {
	request := repositories.CreateRepoRequest{}
	result, err := RepositoryService.CreateRepo(request)
	assert.Nil(t, result)
	assert.NotNil(t, err)
}

func TestCreateRepoProviderError(t *testing.T) {

}

func TestCreateRepoSuccess(t *testing.T) {

}
