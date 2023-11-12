package internal

import (
	"encoding/json"
	"producer/client"
	"strconv"
	"time"
)

type Producer struct {
	Client  *client.ProducerClient
	Message string
	Seconds int
}

type Message struct {
	Content string `json:"content"`
}

func NewProducer(postURL string, postEndpoint string, message string, msgTime int) *Producer {
	return &Producer{
		Client:  client.NewProducerClient(postURL, postEndpoint),
		Message: message,
		Seconds: msgTime,
	}
}

func (pr *Producer) SendMsg() {
	for i := 0; true; i++ {
		msg := Message{Content: pr.Message + " " + strconv.Itoa(i)}

		jsonData, err := json.Marshal(msg)
		if err != nil {
			panic(err)
		}

		pr.Client.PostMsg(jsonData)

		time.Sleep(time.Second * time.Duration(pr.Seconds))
	}
}
