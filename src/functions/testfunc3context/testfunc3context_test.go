package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunctionFailsLocally(t *testing.T) {
	req, err := http.NewRequest("GET", "/.netlify/functions/testfunc1", nil)
	assert.NoError(t, err)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(lambdaHandler)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)

}
