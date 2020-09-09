package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/hello", HelloHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello world!"))
}
