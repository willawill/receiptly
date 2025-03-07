package main

import (
	"net/http"
	"receiptly/domain/entities"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	return router
}

func main() {
	router := setupRouter()
	router.Run(":8080")
}

type UUIDGenerator interface {
	New() string
}

type RealUUIDGenerator struct{}

func (g RealUUIDGenerator) New() string {
	return uuid.New().String()
}

func ProcessReceipt(uuidGenerator UUIDGenerator) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newReceipt entities.Receipt
		if err := c.ShouldBindJSON(&newReceipt); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newReceipt.ID = uuidGenerator.New()
		c.JSON(http.StatusOK, newReceipt)
	}
}

type RealUUIDGenerator struct{}

func (g RealUUIDGenerator) New() string {
	return uuid.New().String()
}

func ProcessReceipt(uuidGenerator UUIDGenerator) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newReceipt entities.Receipt
		if err := c.ShouldBindJSON(&newReceipt); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newReceipt.ID = uuidGenerator.New()
		c.JSON(http.StatusOK, newReceipt)
	}
}


