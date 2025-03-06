package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestProcessReceipt(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.POST("/receipts/process", ProcessReceipt)
	req, _ := http.NewRequest("POST", "/receipts/process", nil)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	expectedBody := `{"id":"1234567890"}`
	assert.Equal(t, expectedBody, recorder.Body.String(), "Response body should match expected JSON")
}
