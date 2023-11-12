package main

import (
	"broker/internal"
	"broker/internal/api"
	"sync"
)

func main() {
	broker := internal.NewBroker()

	var wg sync.WaitGroup
	wg.Add(2)

	go initAPI(broker, &wg)
	go initBroker(broker, &wg)

	wg.Wait()
}

func initAPI(broker *internal.Broker, wg *sync.WaitGroup) {
	defer wg.Done()

	api.InitAPI(broker)
}

func initBroker(broker *internal.Broker, wg *sync.WaitGroup) {
	defer wg.Done()

	broker.Init()
}
