package router

import (
	con "consumer/internal"

	"github.com/gin-gonic/gin"
)

// Initialize is used to initialize the routes
func Initialize(consumer *con.Consumer) {
	// const for the port number
	const portNumber = ":8181"

	// Initialize the router
	router := gin.Default()

	// Initialize the routes
	InitializeRoutes(router, consumer)

	router.Run(portNumber) // Run the server
}
