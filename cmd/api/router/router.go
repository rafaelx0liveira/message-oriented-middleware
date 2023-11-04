package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize(){
		// const for the port number
		const portNumber = ":8080"

		// Initialize the router 
		router := gin.Default()
	
		// Define the route for the endpoint /ping
		// The context c is the object that allows us to interact with the request and response
		router.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		router.Run(portNumber) // Run the server 
}