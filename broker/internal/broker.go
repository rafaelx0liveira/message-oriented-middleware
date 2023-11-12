package internal

import (
	"sync"
	"time"
)

type Broker struct {
	MessageQueues   []Queue[Message]
	SubscriberQueue Queue[Subscriber]
	mu              sync.Mutex
}

// Public:
func NewBroker() *Broker {
	broker := &Broker{}
	broker.appendMessageQueue(NewSliceQueue[Message]())
	broker.SubscriberQueue = NewSliceQueue[Subscriber]()
	return broker
}

func (b *Broker) Init() {
	for {
		b.mu.Lock()

		hasSubscriber := b.subscriberQueueLength() > 0
		hasMessages := b.messageQueueLength() > 0

		if hasSubscriber && hasMessages {

			subscriber, err := b.receiveSubscriber()
			if err != nil {
				panic(err)
			}

			message, err := b.receiveMessage()
			if err != nil {
				panic(err)
			}

			PostMessage(subscriber, message)
		}

		b.mu.Unlock()
		time.Sleep(10 * time.Millisecond)
	}
}

func (b *Broker) SendMessage(message *Message) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	return b.MessageQueues[0].enqueue(message)
}

func (b *Broker) RegisterSubscriber(subscriber *Subscriber) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	return b.SubscriberQueue.enqueue(subscriber)
}

// Private:
func (b *Broker) appendMessageQueue(q Queue[Message]) {
	b.MessageQueues = append(b.MessageQueues, q)
}

func (b *Broker) receiveMessage() (*Message, error) {
	return b.MessageQueues[0].dequeue()
}

func (b *Broker) receiveSubscriber() (*Subscriber, error) {
	return b.SubscriberQueue.dequeue()
}

func (b *Broker) messageQueueLength() int {
	return b.MessageQueues[0].len()
}

func (b *Broker) subscriberQueueLength() int {
	return b.SubscriberQueue.len()
}
