package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	return router
}

func main() {
	router := setupRouter()
	router.Run(":8080")
}

func ProcessReceipt(c *gin.Context) {
	response := map[string]string{"id": "1234567890"}
	c.JSON(http.StatusOK, response)
}
