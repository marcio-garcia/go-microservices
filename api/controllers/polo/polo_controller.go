package polo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	polo = "polo"
)

// Marco -
func Marco(context *gin.Context) {
	context.String(http.StatusOK, polo)
}
