package main

import (
	"broker/internal/api/config"
	"broker/internal/api/router"

	"github.com/gin-gonic/gin"
)

// Define a global variable to store the logger
var (
	logger *config.Logger
)

func main() {
	// Set the Gin to release mode
	gin.SetMode(gin.ReleaseMode)

	// Initialize the logger
	logger = config.GetLogger("main")

	// Print a message informing that the application is starting
	logger.Info("Starting the application...")
	
	// Initialize the Initialize function from router package
	// In Go you don't import the file, you import the package
	router.Initialize()
}
