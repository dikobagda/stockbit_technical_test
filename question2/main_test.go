package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Server(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/movies/batman/page/1", nil)
	response := httptest.NewRecorder()
	setupRouter().ServeHTTP(response, request)

	statusCode := http.StatusOK
	assert.Equal(t, statusCode, response.Code, "Invalid response code")
	t.Log(response.Body.String())
}
