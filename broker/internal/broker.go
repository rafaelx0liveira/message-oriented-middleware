package internal

import "fmt"

type Broker struct {
	MessageQueues   []Queue[Message]
	SubscriberQueue Queue[Subscriber]
}

// Public:
func NewBroker() *Broker {
	broker := &Broker{}
	broker.appendMessageQueue(NewSliceQueue[Message]())
	broker.SubscriberQueue = NewSliceQueue[Subscriber]()
	return broker
}

func (b *Broker) SendMessage(message *Message) error {
	return b.MessageQueues[0].enqueue(message)
}

func (b *Broker) ReceiveMessage() (*Message, error) {
	return b.MessageQueues[0].dequeue()
}

func (b *Broker) RegisterSubscriber(subscriber *Subscriber) error {
	fmt.Print((b.SubscriberQueue.len()))
	return b.SubscriberQueue.enqueue(subscriber)
}

// Private:
func (b *Broker) appendMessageQueue(q Queue[Message]) {
	b.MessageQueues = append(b.MessageQueues, q)
}
