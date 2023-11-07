package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rafaelx0liveira/message-oriented-middleware/cmd/api/config"
	"github.com/rafaelx0liveira/message-oriented-middleware/cmd/api/router"
)

// Define a global variable to store the logger
var (
	logger *config.Logger
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	logger = config.GetLogger("main")

	// Print a message informing that the application is starting
	logger.Info("Starting the application...")
	
	// Initialize the Initialize function from router package
	// In Go you don't import the file, you import the package
	router.Initialize()
}