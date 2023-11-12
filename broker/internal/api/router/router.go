package router

import (
	"github.com/gin-gonic/gin"
	"broker/internal"
)

// Initialize is used to initialize the routes
func Initialize(broker *internal.Broker){
		// const for the port number
		const portNumber = ":8080"

		// Initialize the router 
		router := gin.Default()
	
		// Initialize the routes
		InitializeRoutes(router, broker)

		router.Run(portNumber) // Run the server 
}