package internal

import "errors"

var (
	ErrQueueEmpty = errors.New("queue empty")
	ErrQueueFull  = errors.New("queue full")
)

type Queue[T any] interface {
	enqueue(*T) error
	dequeue() (*T, error)
	len() int
}

type SliceQueue[T any] []*T

func NewSliceQueue[T any]() Queue[T] {
	return &SliceQueue[T]{}
}

func (q *SliceQueue[T]) len() int {
	return len(*q)
}

func (q *SliceQueue[T]) enqueue(value *T) error {
	*q = append(*q, value)
	return nil
}

func (q *SliceQueue[T]) dequeue() (*T, error) {
	queue := *q
	if len(*q) > 0 {
		element := queue[0]
		*q = queue[1:]
		return element, nil
	}

	var empty T
	return &empty, ErrQueueEmpty
}
