package messagepublisher

import (
	"broker/internal"
)

func Init(broker *internal.Broker) {
	broker.SendMessage(&internal.Message{Content: "Hello World!"})
}
