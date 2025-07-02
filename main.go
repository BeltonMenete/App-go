/* package main

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
 */

 package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	scalar "github.com/MarceloPetrucio/go-scalar-api-reference"
)

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	}

	if err != nil {
		fmt.Println("Failed to open browser:", err)
	}
}

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	router.GET("/scalar", func(c *gin.Context) {
		htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
			SpecURL: "./docs/swagger.json", // Make sure this file exists
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

	// Open browser after a short delay so server can start
	go func() {
		time.Sleep(500 * time.Millisecond)
		openBrowser("http://localhost:8080/scalar")
	}()

	fmt.Println("Starting server at http://localhost:8080")
	router.Run(":8080")
}
