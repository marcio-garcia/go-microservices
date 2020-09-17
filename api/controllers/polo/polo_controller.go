package polo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	polo = "polo"
)

// Polo -
func Polo(context *gin.Context) {
	context.String(http.StatusOK, polo)
}
