package queue

import "testing"

func TestQueue_Enqueue_Dequeue(t *testing.T) {
	queue := &Queue[int]{}

	//Test when the Queue is empty
	if !queue.IsEmpty() {
		t.Errorf("IsEmpty() should be true for an empty stack")
	}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	firstElement, ok := queue.Dequeue()

	if !ok {
		t.Errorf("Dequeue() should remove element from the front of queue")
	} else {
		expected := 1

		if firstElement != expected {
			t.Errorf("Dequeue() = %v, want %v", firstElement, expected)
		}
	}

	nextElement, ok := queue.Dequeue()

	if !ok {
		t.Errorf("Dequeue() should remove element from the front of queue")
	} else {
		expected := 2

		if nextElement != expected {
			t.Errorf("Dequeue() = %v, want %v", nextElement, expected)
		}
	}

}

func TestQueue_Size(t *testing.T) {
	queue := &Queue[int]{}

	//Test when the Queue is empty
	if !queue.IsEmpty() {
		t.Errorf("IsEmpty() should be true for an empty stack")
	}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	expectedSize := 3
	result := queue.size

	if result != expectedSize {
		t.Errorf("Dequeue() = %v, want %v", result, expectedSize)
	}

	queue.Dequeue()
	expectedSize = 2
	result = queue.size

	if result != expectedSize {
		t.Errorf("Dequeue() = %v, want %v", result, expectedSize)
	}
}
