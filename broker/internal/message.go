package internal

type Message struct {
	ID      string
	Content string
}

func NewMessage(id string, content string) *Message {
	return &Message{
		ID:      id,
		Content: content,
	}
}
