package message_server

type Broker struct {
	Queues []Queue
}

func NewBroker() *Broker {
	return &Broker{}
}

func (b *Broker) AppendQueue(q Queue) {
	b.Queues = append(b.Queues, q)
}
