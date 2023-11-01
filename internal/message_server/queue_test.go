package message_server

import "testing"

func TestCreateQueue_ShouldCreateQueue(t *testing.T) {
	queue := NewSliceQueue()

	if queue == nil {
		t.Error("Queue should not be nil")
	}
}

func TestQueue_ShouldEnqueueAndDequeue(t *testing.T) {
	queue := NewSliceQueue()

	msg := NewMessage(1, "Hello Go")

	queue.Enqueue(*msg)

	if queue.Len() != 1 {
		t.Errorf("Expected queue length = 1")
	}

	popMsg, err := queue.Dequeue()

	if err != nil {
		t.Fatalf("Error dequeuing message: %v", err)
	}

	if popMsg.ID != 1 {
		t.Errorf("Expected message ID 1, got %d", popMsg.ID)
	}

	if popMsg.Content != "Hello Go" {
		t.Errorf("Expected message content 'Hello Go', got '%s'", popMsg.Content)
	}

	if queue.Len() != 0 {
		t.Errorf("Expected queue length = 0")
	}

}
