package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateQueue_ShouldCreateQueue(t *testing.T) {
	assert := assert.New(t)

	queue := NewSliceQueue[Message]()

	assert.NotNil(queue)
}

func TestEnqueue_ShouldEnqueueMessage(t *testing.T) {
	assert := assert.New(t)

	queue := &SliceQueue[Message]{}
	queue.enqueue(NewMessage(1, "First message"))

	assert.Equal(1, (*queue)[0].ID)
	assert.Equal("First message", (*queue)[0].Content)
	assert.Equal(1, len(*queue))

	queue.enqueue(NewMessage(2, "Second message"))

	assert.Equal(2, (*queue)[1].ID)
	assert.Equal("Second message", (*queue)[1].Content)
	assert.Equal(2, len(*queue))
}

func TestDequeue_ShouldDequeueMessage(t *testing.T) {
	assert := assert.New(t)

	queue := &SliceQueue[Message]{}

	queue.enqueue(NewMessage(1, "Hello Go"))
	dequeuedMessage, err := queue.dequeue()

	assert.Nil(err)
	assert.NotNil(dequeuedMessage)

	assert.Equal(1, dequeuedMessage.ID, "they should be equal")
	assert.Equal("Hello Go", dequeuedMessage.Content, "they should be equal")
}

func TestGetQueueLength_ShouldReturnLength(t *testing.T) {
	assert := assert.New(t)

	//queue -> []*Message
	queue := &SliceQueue[Message]{}

	assert.Equal(0, queue.len())

	queue.enqueue(NewMessage(1, "Message 1"))
	queue.enqueue(NewMessage(2, "Message 2"))
	queue.enqueue(NewMessage(3, "Message 3"))

	assert.Equal(3, queue.len())

	queue.dequeue()

	assert.Equal(2, queue.len())
}

func TestGetQueueLength_ShouldReturnLength2(t *testing.T) {
	type args struct {
		queue       *SliceQueue[Message]
		enqueueMsgs []Message
		dequeueMsgs int
	}

	tests := []struct {
		name      string
		arguments args
		want      int
	}{
		{"Test len 0", args{queue: &SliceQueue[Message]{}, enqueueMsgs: []Message{}, dequeueMsgs: 0}, 0},
		{"Test len 0 - Only Dequeue", args{queue: &SliceQueue[Message]{}, enqueueMsgs: []Message{}, dequeueMsgs: 1}, 0},
		{"Test len 3 - Only Enqueue",
			args{
				queue: &SliceQueue[Message]{},
				enqueueMsgs: []Message{
					{ID: 1, Content: "Mesg 1"},
					{ID: 2, Content: "Mesg 2"},
					{ID: 3, Content: "Mesg 3"},
				},
				dequeueMsgs: 0,
			}, 3},
		{"Te len 2 - Enqueue and Dequeue", args{queue: &SliceQueue[Message]{},
			enqueueMsgs: []Message{
				{ID: 1, Content: "Mesg 1"},
				{ID: 2, Content: "Mesg 2"},
				{ID: 3, Content: "Mesg 3"}},
			dequeueMsgs: 1,
		}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := tt.arguments.queue

			for _, enqMsgs := range tt.arguments.enqueueMsgs {
				queue.enqueue(&enqMsgs)
			}

			for i := 0; i < tt.arguments.dequeueMsgs; i++ {
				queue.dequeue()
			}

			got := queue.len()
			if got != tt.want {
				t.Errorf("Len() got = %v, want %v", got, tt.want)
			}
		})
	}
}
