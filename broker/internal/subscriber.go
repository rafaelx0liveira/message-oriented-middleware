package internal

type Subscriber struct {
	ID          int
	EventType   string
	ConsumerUrl string
}

func NewSubscriber(id int, eventType string, url string) *Subscriber {
	return &Subscriber{
		ID:          id,
		EventType:   eventType,
		ConsumerUrl: url,
	}
}
