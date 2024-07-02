package test

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func HelloHandler(c *gin.Context) {
	// Start a new goroutine to handle the request concurrently
	go func() {
		// Simulate some processing time
		time.Sleep(8 * time.Second)

		// Respond to the client's request once the processing is done
		fmt.Println("Hello, world!")
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, world!",
		})
	}()
}
