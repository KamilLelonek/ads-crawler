package publisher

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPublisherRoutes(t *testing.T) {
	router := prepareRouter()
	recorder := performRequest(router, http.MethodGet, "/v1/publisher/random")
	response, err := readJson(recorder)
	status, exists := response["publisher"]

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, "random", status)
}

func prepareRouter() http.Handler {
	router := gin.Default()
	v1 := router.Group("/v1")

	MountRoutes(v1)

	return router
}

func performRequest(handler http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, req)

	return recorder
}

func readJson(recorder *httptest.ResponseRecorder) (map[string]string, error) {
	var response map[string]string

	err := json.Unmarshal([]byte(recorder.Body.String()), &response)

	return response, err
}
