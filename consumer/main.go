package main

import (
	"consumer/api"
	inter "consumer/internal"
)

func main() {
	consumer := inter.NewConsumer("127.0.0.1:8180", "127.0.0.1:7779", "/api/v1/subscribe", 30)

	api.InitAPI(consumer)
}
