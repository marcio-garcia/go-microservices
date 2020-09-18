package githubprovider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/marcio-garcia/go-microservices/api/clients/restclient"
	"github.com/marcio-garcia/go-microservices/api/domain/github"
	"github.com/marcio-garcia/go-microservices/api/domain/repositories"
	"github.com/marcio-garcia/go-microservices/api/providers"
	"github.com/marcio-garcia/go-microservices/api/utils/errors"
)

const (
	headerAuthorizationKey    = "Authorization"
	headerAuthorizationFormat = "token %s"

	createRepoURL = "https://api.github.com/user/repos"
)

// GithubProvider implements the ProviderInterface
var (
	GithubProvider providers.ProviderInterface
)

func init() {
	GithubProvider = &githubProvider{}
}

type githubProvider struct{}

// CreateRepo requests the api to create a new repository
func (p *githubProvider) CreateRepo(accessToken string, request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, *errors.APIError) {
	//func CreateRepo(accessToken string, request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.ErrorResponse) {
	headers := http.Header{}
	headers.Set(headerAuthorizationKey, getAuthorizationHeader(accessToken))

	githubRequest := github.CreateRepoRequest{
		Name:        request.Name,
		Description: request.Description,
		Private:     false,
	}

	response, err := restclient.RestClient.Post(createRepoURL, githubRequest, headers)
	if err != nil {
		log.Printf("Error when trying to create a new repo in github: %s", err.Error())
		responseError := errors.CreateError(http.StatusInternalServerError, err.Error())
		return nil, &responseError
	}

	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		responseError := errors.CreateInternalServerError("Invalid response body")
		return nil, &responseError
	}

	if response.StatusCode > 299 {
		var gitResponseError github.ErrorResponse
		if err := json.Unmarshal(body, &gitResponseError); err != nil {
			responseError := errors.CreateInternalServerError("Invalid json response body")
			return nil, &responseError
		}
		responseError := errors.CreateError(response.StatusCode, gitResponseError.Message)
		return nil, &responseError
	}

	var githubResult github.CreateRepoResponse
	if err := json.Unmarshal(body, &githubResult); err != nil {
		log.Printf("Error when trying to unmarshal the successful json response body: %s", err.Error())
		responseError := errors.CreateInternalServerError("Error when trying to unmarshal the successful json response body")
		return nil, &responseError
	}

	result := repositories.CreateRepoResponse{
		ID:    githubResult.ID,
		Name:  githubResult.Name,
		Owner: githubResult.Owner.Login,
	}

	return &result, nil
}

func getAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, accessToken)
}
