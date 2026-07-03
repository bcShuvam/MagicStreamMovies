package main

import (
	"fmt"

	// Give it the 'controller' alias explicitly right here
	controller "MagicStreamMoviesServer/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, MagicStreamMovies!")
	})

	// This now matches the 'controller' alias perfectly
	router.GET("/movies", controller.GetMovies())

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Failed to start server", err)
	}
}
