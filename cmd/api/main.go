package main

import (
	"github.com/rafaelx0liveira/message-oriented-middleware/cmd/api/router"
)

func main() {
	// Initialize the Initialize function from router package
	// In Go you don't import the file, you import the package
	router.Initialize()
}