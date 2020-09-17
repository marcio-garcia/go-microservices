package app

import (
	"github.com/marcio-garcia/go-microservices/api/controllers/polo"
	"github.com/marcio-garcia/go-microservices/api/controllers/repositories"
)

func createRoutes() {
	router.GET("/marco", polo.Polo)
	router.POST("/repositories", repositories.CreateRepo)
}
