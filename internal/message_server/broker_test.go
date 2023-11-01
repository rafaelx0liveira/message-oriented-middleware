package message_server

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateBroker_ShouldCreateBroker(t *testing.T) {
	assert := assert.New(t)

	broker := NewBroker()

	if assert.NotNil(broker) {
		assert.Equal(0, len(broker.Queues))
	}
}

func TestAppendQueueToBroker_ShouldAppendQueue(t *testing.T) {
	assert := assert.New(t)

	queue := NewSliceQueue()
	broker := NewBroker()

	broker.AppendQueue(queue)
	len := len(broker.Queues)

	assert.Equal(1, len, "Broker should have one queue")
}
