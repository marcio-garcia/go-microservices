package app

import (
	"github.com/marcio-garcia/go-microservices/api/controllers/polo"
	"github.com/marcio-garcia/go-microservices/api/controllers/repositories"
)

func createRoutes() {
	router.GET("/marco", polo.Marco)
	router.POST("/repositories", repositories.CreateRepo)
}
