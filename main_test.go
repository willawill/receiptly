package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MockUUIDGenerator struct{}

func (m MockUUIDGenerator) New() string {
	return "1234567890"
}

func TestProcessReceipt(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	mockUUIDGen := MockUUIDGenerator{}

	t.Run("ProcessReceipt returns 200 when receipt is processed", func(t *testing.T) {
		router.POST("/receipts/process", ProcessReceipt(mockUUIDGen))
		requestBody := `{"retailer": "Target", "purchaseDate": "2024-01-01", "purchaseTime": "13:01", "items": [{"shortDescription": "Pepsi - 12-oz", "price": "1.25"}]}`
		req, _ := http.NewRequest("POST", "/receipts/process", strings.NewReader(requestBody))

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)

		expectedBody := `{"id":"1234567890"}`
		assert.Equal(t, expectedBody, recorder.Body.String(), "Response body should match expected JSON")
	})

	t.Run("ProcessReceipt returns 400 when receipt body is empty", func(t *testing.T) {
		router.POST("/receipts/process", ProcessReceipt(mockUUIDGen))
		invalidBody := `{}`
		req, _ := http.NewRequest("POST", "/receipts/process", strings.NewReader(invalidBody))

		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
		assert.Contains(t, recorder.Body.String(), "error")
	})

}
