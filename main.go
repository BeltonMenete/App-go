package main

import (
	"fmt"
	"net/http"

	scalar "github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Simple ping route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "test-gin",
		})
	})

	// Scalar reference UI route
	router.GET("/reference", func(c *gin.Context) {
		htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
			SpecURL: "./docs/swagger.json", // Or a public URL
			CustomOptions: scalar.CustomOptions{
				PageTitle: "Simple API",
			},
			DarkMode: true,
		})
		if err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %v", err))
			return
		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlContent))
	})

	// Start server
	fmt.Println("Starting server on :8080")
	router.Run(":8080")
}
