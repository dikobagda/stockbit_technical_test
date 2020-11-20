package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Server(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()

	Server().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "Invalid response code")
	t.Log(response.Body.String())
}
