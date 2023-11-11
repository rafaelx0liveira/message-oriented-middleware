package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateBroker_ShouldCreateBroker(t *testing.T) {
	assert := assert.New(t)

	broker := NewBroker()

	if assert.NotNil(broker) {
		assert.Equal(1, len(broker.MessageQueues))
	}
}

func TestAppendQueueToBroker_ShouldAppendQueue(t *testing.T) {
	assert := assert.New(t)

	queue := NewSliceQueue[Message]()
	broker := NewBroker()

	broker.appendMessageQueue(queue)
	len := len(broker.MessageQueues)

	assert.Equal(2, len, "Broker should have two queues")
}

func TestSendMessageToBroker_ShouldSendMessage(t *testing.T) {
	assert := assert.New(t)

	broker := NewBroker()
	message := NewMessage(1, "First message")

	err := broker.SendMessage(message)

	assert.Nil(err)
	assert.Equal(1, broker.MessageQueues[0].len())

	message = NewMessage(2, "Second message")
	err = broker.SendMessage(message)
	assert.Nil(err)

	assert.Equal(2, broker.MessageQueues[0].len())

	firstMsg, _ := broker.MessageQueues[0].dequeue()
	secondMsg, _ := broker.MessageQueues[0].dequeue()

	assert.Equal(1, firstMsg.ID)
	assert.Equal("First message", firstMsg.Content)

	assert.Equal(2, secondMsg.ID)
	assert.Equal("Second message", secondMsg.Content)
}

func TestReceiveMessageToBroker_ShouldReceiveMessage(t *testing.T) {
	assert := assert.New(t)

	broker := NewBroker()
	message := NewMessage(1, "First message")

	err := broker.MessageQueues[0].enqueue(message)

	assert.Nil(err)
	assert.Equal(1, broker.MessageQueues[0].len())

	message = NewMessage(2, "Second message")
	err = broker.MessageQueues[0].enqueue(message)
	assert.Nil(err)

	firstMsg, _ := broker.ReceiveMessage()
	secondMsg, _ := broker.ReceiveMessage()

	assert.Equal(1, firstMsg.ID)
	assert.Equal("First message", firstMsg.Content)

	assert.Equal(2, secondMsg.ID)
	assert.Equal("Second message", secondMsg.Content)
}
