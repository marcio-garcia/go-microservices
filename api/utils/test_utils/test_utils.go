package test_utils

import (
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

// GetMockedContext -
func GetMockedContext(httpMethod string, url string, body io.Reader) (*gin.Context, *httptest.ResponseRecorder) {
	response := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(response)
	request, _ := http.NewRequest(httpMethod, url, body)
	context.Request = request
	return context, response
}
