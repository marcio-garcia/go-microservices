package repositories

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcio-garcia/go-microservices/api/domain/repositories"
	"github.com/marcio-garcia/go-microservices/api/services"
	"github.com/marcio-garcia/go-microservices/api/utils/errors"
)

// CreateRepo is the entry point for creating a new repository
func CreateRepo(context *gin.Context) {
	var request repositories.CreateRepoRequest
	if err := context.ShouldBindJSON(&request); err != nil {
		apiError := errors.CreateBadRequestError("Invalid json body")
		context.JSON(apiError.Status(), apiError)
		return
	}

	result, err := services.RepositoryService.CreateRepo(request)
	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusCreated, result)
}
