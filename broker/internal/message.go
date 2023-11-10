package internal

type Message struct {
	ID      int
	Content string
}

func NewMessage(id int, content string) *Message {
	return &Message{
		ID:      id,
		Content: content,
	}
}
