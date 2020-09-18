package repositories

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/marcio-garcia/go-microservices/api/domain/repositories"
	"github.com/marcio-garcia/go-microservices/api/services"
	"github.com/marcio-garcia/go-microservices/api/utils/errors"
	"github.com/marcio-garcia/go-microservices/api/utils/test_utils"
	"github.com/stretchr/testify/assert"
)

func TestCreateRepoInvalidJSONRequest(t *testing.T) {
	context, response := test_utils.GetMockedContext(http.MethodPost, "/repositories", strings.NewReader(""))
	CreateRepo(context)

	assert.EqualValues(t, http.StatusBadRequest, response.Code)
}

func TestCreateRepoErrorFromService(t *testing.T) {
	context, response := test_utils.GetMockedContext(http.MethodPost, "/repositories", strings.NewReader(`{"name": "testing"}`))

	services.RepositoryService = &repoServiceErrorMock{}

	CreateRepo(context)

	assert.EqualValues(t, http.StatusUnauthorized, response.Code)
}

func TestCreateRepoSuccess(t *testing.T) {
	context, response := test_utils.GetMockedContext(http.MethodPost, "/repositories", strings.NewReader(`{"name": "testing"}`))

	services.RepositoryService = &repoServiceSuccessMock{}

	CreateRepo(context)

	assert.EqualValues(t, http.StatusCreated, response.Code)

	var result repositories.CreateRepoResponse
	err := json.Unmarshal(response.Body.Bytes(), &result)
	assert.Nil(t, err)
	assert.EqualValues(t, 123, result.ID)
	assert.EqualValues(t, "testing", result.Name)
	assert.EqualValues(t, "owner", result.Owner)
}

//************
//**** Mockups
//************

type repoServiceErrorMock struct{}

func (rs *repoServiceErrorMock) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.APIError) {
	return nil, errors.CreateError(401, "Requires authorization")
}

type repoServiceSuccessMock struct{}

func (rs *repoServiceSuccessMock) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.APIError) {
	response := repositories.CreateRepoResponse{
		ID:    123,
		Name:  "testing",
		Owner: "owner",
	}
	return &response, nil
}
