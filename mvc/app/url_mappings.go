package app

import (
	"github.com/marcio-garcia/go-microservices/mvc/controllers"
)

func mapURL() {
	router.GET("/users/:id", controllers.GetUser)
}
