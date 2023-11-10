package message_server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBroker_ShouldCreateBroker(t *testing.T) {
	assert := assert.New(t)

	broker := NewBroker()

	if assert.NotNil(broker) {
		assert.Equal(1, len(broker.Queues))
	}
}

func TestAppendQueueToBroker_ShouldAppendQueue(t *testing.T) {
	assert := assert.New(t)

	queue := NewSliceQueue()
	broker := NewBroker()

	broker.appendQueue(queue)
	len := len(broker.Queues)

	assert.Equal(2, len, "Broker should have two queues")
}

func TestSendMessageToBroker_ShouldSendMessage(t *testing.T) {
	assert := assert.New(t)

	broker := NewBroker()
	message := NewMessage(1, "First message")

	err := broker.SendMessage(message)

	assert.Nil(err)
	assert.Equal(1, broker.Queues[0].len())

	message = NewMessage(2, "Second message")
	err = broker.SendMessage(message)
	assert.Nil(err)

	assert.Equal(2, broker.Queues[0].len())

	firstMsg, _ := broker.Queues[0].dequeue()
	secondMsg, _ := broker.Queues[0].dequeue()

	assert.Equal(1, firstMsg.ID)
	assert.Equal("First message", firstMsg.Content)

	assert.Equal(2, secondMsg.ID)
	assert.Equal("Second message", secondMsg.Content)
}

func TestReceiveMessageToBroker_ShouldReceiveMessage(t *testing.T) {
	assert := assert.New(t)

	broker := NewBroker()
	message := NewMessage(1, "First message")

	err := broker.Queues[0].enqueue(message)

	assert.Nil(err)
	assert.Equal(1, broker.Queues[0].len())

	message = NewMessage(2, "Second message")
	err = broker.Queues[0].enqueue(message)
	assert.Nil(err)

	firstMsg, _ := broker.ReceiveMessage()
	secondMsg, _ := broker.ReceiveMessage()

	assert.Equal(1, firstMsg.ID)
	assert.Equal("First message", firstMsg.Content)

	assert.Equal(2, secondMsg.ID)
	assert.Equal("Second message", secondMsg.Content)
}
