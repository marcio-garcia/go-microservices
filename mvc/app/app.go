package app

import (
	"net/http"

	"github.com/marcio-garcia/go-microservices/mvc/controllers"
)

// Start entry point of the app
func Start() {
	http.HandleFunc("/users", controllers.GetUser)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}
