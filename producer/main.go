package main

import (
	pr "producer/internal"
)

func main() {
	producer := pr.NewProducer("http://localhost:8080", "application/json", "Hello World!")

	producer.SendMsg()
}