package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlerSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/.netlify/functions/testfunc1", nil)
	assert.NoError(t, err)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(helloFunc)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	expected := "This is go1"
	assert.Equal(t, expected, rr.Body.String())
}
