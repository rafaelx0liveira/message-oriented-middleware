package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rafaelx0liveira/message-oriented-middleware/internal/message_server/api/api/config"
	"github.com/rafaelx0liveira/message-oriented-middleware/internal/message_server/api/api/router"
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