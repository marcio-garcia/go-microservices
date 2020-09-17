package githubprovider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/marcio-garcia/go-microservices/api/clients/restclient"
	"github.com/marcio-garcia/go-microservices/api/domain/github"
)

const (
	headerAuthorizationKey    = "Authorization"
	headerAuthorizationFormat = "token %s"

	createRepoURL = "https://api.github.com/user/repos"
)

func getAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, accessToken)
}

// CreateRepo requests the api to create a new repository
func CreateRepo(accessToken string, request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.ErrorResponse) {
	headers := http.Header{}
	headers.Set(headerAuthorizationKey, getAuthorizationHeader(accessToken))
	response, err := restclient.RestClient.Post(createRepoURL, request, headers)
	if err != nil {
		log.Printf("Error when trying to create a new repo in github: %s", err.Error())
		errorResponse := createErrorResponse(http.StatusInternalServerError, err.Error())
		return nil, &errorResponse
	}

	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		errorResponse := createErrorResponse(http.StatusInternalServerError, "Invalid response body")
		return nil, &errorResponse
	}

	if response.StatusCode > 299 {
		var errorResponse github.ErrorResponse
		if err := json.Unmarshal(body, &errorResponse); err != nil {
			errorResponse := createErrorResponse(http.StatusInternalServerError, "Invalid json response body")
			return nil, &errorResponse
		}
		errorResponse.StatusCode = response.StatusCode
		return nil, &errorResponse
	}

	var result github.CreateRepoResponse
	if err := json.Unmarshal(body, &result); err != nil {
		log.Printf("Error when trying to unmarshal the successful json response body: %s", err.Error())
		errorResponse := createErrorResponse(http.StatusInternalServerError, "Error when trying to unmarshal the successful json response body")
		return nil, &errorResponse
	}

	return &result, nil
}

func createErrorResponse(statusCode int, message string) github.ErrorResponse {
	errorResponse := github.ErrorResponse{
		StatusCode: statusCode,
		Message:    message,
	}
	return errorResponse
}
