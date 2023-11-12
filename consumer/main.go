package main

import (
	"consumer/api"
	inter "consumer/internal"
	"sync"
)

func main() {
	consumer := inter.NewConsumer("http://127.0.0.1:8181/api/v1/webhook", "http://127.0.0.1:8180/api/v1/subscribe", "application/json", 1000)

	var wg sync.WaitGroup
	wg.Add(2)

	go InitAPI(consumer, &wg)
	go InitConsumer(consumer, &wg)

	wg.Wait()
}

func InitAPI(consumer *inter.Consumer, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	api.InitAPI(consumer)
}

func InitConsumer(consumer *inter.Consumer, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	consumer.LaunchHooksPeriodically(1000)
}
