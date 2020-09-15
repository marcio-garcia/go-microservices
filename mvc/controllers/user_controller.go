package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/marcio-garcia/go-microservices/mvc/domain"

	"github.com/marcio-garcia/go-microservices/mvc/services"
)

// GetUser retrieve a user
func GetUser(context *gin.Context) {
	userID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		apiError := domain.AppError{
			Message:    "User id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		context.JSON(apiError.StatusCode, apiError)
		return
	}

	user, apiError := services.UserService.GetUser(uint64(userID))

	if apiError != nil {
		context.JSON(apiError.StatusCode, apiError)
		return
	}

	context.JSON(http.StatusOK, user)
}
