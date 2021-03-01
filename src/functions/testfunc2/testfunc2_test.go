package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/.netlify/functions/testfunc2", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(lambdaHandler)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	expected := "Looks like you made it to testfunc2!"
	assert.Equal(t, expected, rr.Body.String())
}
