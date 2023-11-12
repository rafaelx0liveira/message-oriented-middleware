package main

import (
	pr "producer/internal"
)

func main() {
	producer := pr.NewProducer("http://localhost:8180/api/v1/publish", "application/json", "Gonheç!", 1)

	producer.SendMsg()
}
