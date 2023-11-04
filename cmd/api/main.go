package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// const for the port number
	const portNumber = ":8080"

	// Initialize the router 
	r := gin.Default()

	// Define the route for the endpoint /ping
	// The context c is the object that allows us to interact with the request and response
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(portNumber) // Run the server 
}