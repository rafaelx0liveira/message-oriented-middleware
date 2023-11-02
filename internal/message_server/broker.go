package message_server

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
	return b.Queues[0].Enqueue(message)
}

func (b *Broker) ReceiveMessage() (*Message, error) {
	return b.Queues[0].Dequeue()
}

// Private:
func (b *Broker) appendQueue(q Queue) {
	b.Queues = append(b.Queues, q)
}
