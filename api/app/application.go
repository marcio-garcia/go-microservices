package app

import (
	"github.com/gin-gonic/gin"
	"github.com/marcio-garcia/go-microservices/api/log"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

// Start entry point of the app
func Start() {
	log.Info("Start mapping the routes", "step:1", "status:pending")
	createRoutes()
	log.Info("Routes successfully mapped", "step:2", "status:success")

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
