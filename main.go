package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thaijdk/finalexam/customer"
	"github.com/thaijdk/finalexam/database"
)

func authMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token != "token2019" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": http.StatusUnauthorized})
		c.Abort()
		return
	}
	c.Next()
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(authMiddleware)
	r.POST("/customers", customer.PostHandler)
	r.GET("/customers/:id", customer.GetByIdHandler)
	r.GET("/customers", customer.GetHandler)
	r.PUT("/customers/:id", customer.UpdateHandler)
	r.DELETE("/customers/:id", customer.DeleteByIdHandler)
	return r
}

func main() {
	database.SetDB()
	r := setupRouter()
	r.Run(":2019")
}
