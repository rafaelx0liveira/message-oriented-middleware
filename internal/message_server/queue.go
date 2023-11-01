package message_server

import "errors"

var (
	ErrQueueEmpty = errors.New("queue empty")
	ErrQueueFull  = errors.New("queue full")
)

type Queue interface {
	Enqueue(*Message) error
	Dequeue() (*Message, error)
	Len() int
}

type SliceQueue []*Message

func NewSliceQueue() Queue {
	return &SliceQueue{}
}

func (q *SliceQueue) Len() int {
	return len(*q)
}

func (q *SliceQueue) Enqueue(value *Message) error {
	*q = append(*q, value)
	return nil
}

func (q *SliceQueue) Dequeue() (*Message, error) {
	queue := *q
	if len(*q) > 0 {
		card := queue[0]
		*q = queue[1:]
		return card, nil
	}

	var empty Message
	return &empty, ErrQueueEmpty
}
