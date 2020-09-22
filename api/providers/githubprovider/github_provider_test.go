package githubprovider

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/marcio-garcia/go-microservices/api/clients/restclient"

	"github.com/marcio-garcia/go-microservices/api/domain/repositories"

	"github.com/stretchr/testify/assert"
)

func TestConstants(t *testing.T) {
	assert.EqualValues(t, "Authorization", headerAuthorizationKey)
	assert.EqualValues(t, "token %s", headerAuthorizationFormat)
	assert.EqualValues(t, "https://api.github.com/user/repos", createRepoURL)
}

func TestGetAuthorizationHeader(t *testing.T) {
	header := getAuthorizationHeader("aaaa")
	assert.EqualValues(t, "token aaaa", header)
}

func TestCreateRepoRestClientError(t *testing.T) {
	restclient.RestClient = &restClientRestClientErrorMock{}
	response, err := GithubProvider.CreateRepo("", repositories.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, (*err).Status())
	assert.EqualValues(t, "Error returned by the REST client", (*err).Message())
}

func TestCreateRepoInvalidResponseBody(t *testing.T) {
	restclient.RestClient = &restClientInvalidResponseBodyMock{}
	response, err := GithubProvider.CreateRepo("", repositories.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, (*err).Status())
	assert.EqualValues(t, "Invalid response body", (*err).Message())
}

func TestCreateRepoStatusCodeErrorInvalidResponseBody(t *testing.T) {
	restclient.RestClient = &restClientStatusCodeErrorInvalidResponseBodyMock{}
	response, err := GithubProvider.CreateRepo("", repositories.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, (*err).Status())
	assert.EqualValues(t, "Invalid json response body", (*err).Message())
}

func TestCreateRepoStatusCodeError(t *testing.T) {
	restclient.RestClient = &restClientStatusCodeErrorMock{}
	response, err := GithubProvider.CreateRepo("", repositories.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, (*err).Status())
	assert.EqualValues(t, "Requires authentication", (*err).Message())
}

func TestCreateRepoSuccessInvalidResponseBody(t *testing.T) {
	restclient.RestClient = &restClientSuccessInvalidResponseBodyMock{}
	response, err := GithubProvider.CreateRepo("", repositories.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, (*err).Status())
	assert.EqualValues(t, "Error when trying to unmarshal the successful json response body", (*err).Message())
}

func TestCreateRepoSuccess(t *testing.T) {
	restclient.RestClient = &restClientSuccessMock{}
	response, err := GithubProvider.CreateRepo("", repositories.CreateRepoRequest{})
	assert.NotNil(t, response)
	assert.Nil(t, err)
	assert.EqualValues(t, 1, response.ID)
	assert.EqualValues(t, "go-test", response.Name)
}

// **************
// ****** Mockups
// **************

type restClientRestClientErrorMock struct{}

// Post makes POST requests
func (rc *restClientRestClientErrorMock) Post(URL string, body interface{}, header http.Header) (*http.Response, error) {
	err := errors.New("Error returned by the REST client")
	return nil, err
}

type restClientInvalidResponseBodyMock struct{}

// Post makes POST requests
func (rc *restClientInvalidResponseBodyMock) Post(URL string, body interface{}, header http.Header) (*http.Response, error) {
	invalidCloser, _ := os.Open("-asf3")
	response := http.Response{
		StatusCode: http.StatusCreated,
		Body:       invalidCloser,
	}

	return &response, nil
}

type restClientStatusCodeErrorInvalidResponseBodyMock struct{}

// Post makes POST requests
func (rc *restClientStatusCodeErrorInvalidResponseBodyMock) Post(URL string, body interface{}, header http.Header) (*http.Response, error) {
	response := http.Response{
		StatusCode: http.StatusNotFound,
		Body:       ioutil.NopCloser(strings.NewReader(`{"message": 1}`)),
	}

	return &response, nil
}

type restClientStatusCodeErrorMock struct{}

// Post makes POST requests
func (rc *restClientStatusCodeErrorMock) Post(URL string, body interface{}, header http.Header) (*http.Response, error) {
	response := http.Response{
		StatusCode: http.StatusUnauthorized,
		Body:       ioutil.NopCloser(strings.NewReader(`{"message": "Requires authentication","documentation_url": "https://link.to.doc/"}`)),
	}

	return &response, nil
}

type restClientSuccessInvalidResponseBodyMock struct{}

// Post makes POST requests
func (rc *restClientSuccessInvalidResponseBodyMock) Post(URL string, body interface{}, header http.Header) (*http.Response, error) {
	response := http.Response{
		StatusCode: http.StatusCreated,
		Body:       ioutil.NopCloser(strings.NewReader(`{ID:"", Name:123, FullName:"", Owner:github.RepoOwner{ID:0, Login:"", URL:"", HTMLURL:""}, Permissions:github.RepoPermissions{Admin:false, Push:false, Pull:false}}`)),
	}

	return &response, nil
}

type restClientSuccessMock struct{}

// Post makes POST requests
func (rc *restClientSuccessMock) Post(URL string, body interface{}, header http.Header) (*http.Response, error) {
	response := http.Response{
		StatusCode: http.StatusCreated,
		Body:       ioutil.NopCloser(strings.NewReader(`{"id":1, "name":"go-test", "full_name":"", "owner":{"id":0, "login":"", "url":"", "html_url":""}, "permissions":{"admin":false, "push":false, "pull":false}}`)),
	}

	return &response, nil
}
