package message_server

import "testing"

func TestCreateBroker_ShouldCreateBroker(t *testing.T) {
	broker := NewBroker()

	if broker == nil {
		t.Error("Broker should not be nil")
	}
}

func TestAppendQueueToBroker_ShouldAppendQueue(t *testing.T) {
	queue := NewSliceQueue()
	broker := NewBroker()

	broker.AppendQueue(queue)
	len := len(broker.Queues)

	if len != 1 {
		t.Error("Broker should have one queue")
	}
}
