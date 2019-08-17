package publisher

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var message = ""

func TestHttpGet(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	response, err := HttpGet(server.URL)
	assert.Nil(t, err)

	buf := new(bytes.Buffer)
	buf.ReadFrom(response)
	result := buf.String()

	assert.Equal(t, message, result)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Print(w, message)
}
