package internal

type Broker struct {
	Queues []Queue
}

// Public:
func NewBroker() *Broker {
	broker := &Broker{}
	broker.appendQueue(NewSliceQueue())
	return broker
}

func (b *Broker) SendMessage(message *Message) error {
	return b.Queues[0].enqueue(message)
}

func (b *Broker) ReceiveMessage() (*Message, error) {
	return b.Queues[0].dequeue()
}

// Private:
func (b *Broker) appendQueue(q Queue) {
	b.Queues = append(b.Queues, q)
}
