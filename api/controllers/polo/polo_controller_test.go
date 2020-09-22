package polo

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/marcio-garcia/go-microservices/api/utils/test_utils"
)

func TestPolo(t *testing.T) {
	context, response := test_utils.GetMockedContext(http.MethodGet, "/marco", nil)

	Marco(context)

	assert.EqualValues(t, http.StatusOK, response.Code)
	assert.EqualValues(t, "polo", response.Body.String())
}
