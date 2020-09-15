package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/marcio-garcia/go-microservices/mvc/domain"

	"github.com/marcio-garcia/go-microservices/mvc/services"
)

// GetUser retrieve a user
func GetUser(response http.ResponseWriter, request *http.Request) {
	userIDParam := request.URL.Query().Get("id")
	userID, err := strconv.ParseInt(userIDParam, 10, 64)
	if err != nil {
		apiError := domain.AppError{
			Message:    "User id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, _ := json.Marshal(apiError)
		response.WriteHeader(apiError.StatusCode)
		response.Write(jsonValue)
		return
	}

	user, apiError := services.UserService.GetUser(uint64(userID))

	if apiError != nil {
		jsonValue, _ := json.Marshal(apiError)
		response.WriteHeader(apiError.StatusCode)
		response.Write(jsonValue)
		return
	}

	// Return the retrieved user
	jsonValue, _ := json.Marshal(user)
	response.Write(jsonValue)
}
