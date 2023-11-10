package internal

import "errors"

var (
	ErrQueueEmpty = errors.New("queue empty")
	ErrQueueFull  = errors.New("queue full")
)

type Queue interface {
	enqueue(*Message) error
	dequeue() (*Message, error)
	len() int
}

type SliceQueue []*Message

func NewSliceQueue() Queue {
	return &SliceQueue{}
}

func (q *SliceQueue) len() int {
	return len(*q)
}

func (q *SliceQueue) enqueue(value *Message) error {
	*q = append(*q, value)
	return nil
}

func (q *SliceQueue) dequeue() (*Message, error) {
	queue := *q
	if len(*q) > 0 {
		card := queue[0]
		*q = queue[1:]
		return card, nil
	}

	var empty Message
	return &empty, ErrQueueEmpty
}
