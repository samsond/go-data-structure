package queue

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestNotificationQueue(t *testing.T) {
	// Create a new NotificationQueue
	nq := NewNotificationQueue()

	// Test Push and Pop
	nq.Push(Notification{UserID: 1, Message: "Message 1"})
	nq.Push(Notification{UserID: 2, Message: "Message 2"})

	notification, ok := nq.Pop()
	if !ok {
		t.Error("Expected Pop to succeed")
	}
	if notification.UserID != 1 || notification.Message != "Message 1" {
		t.Errorf("Expected {1, 'Message 1'}, got %v", notification)
	}

	notification, ok = nq.Pop()
	if !ok {
		t.Error("Expected Pop to succeed")
	}
	if notification.UserID != 2 || notification.Message != "Message 2" {
		t.Errorf("Expected {2, 'Message 2'}, got %v", notification)
	}

	// Test IsEmpty
	if !nq.IsEmpty() {
		t.Error("Expected IsEmpty to return true")
	}

	// Test Size
	if nq.Size() != 0 {
		t.Errorf("Expected Size to be 0, got %d", nq.Size())
	}

}
func producer(nq *NotificationQueue, wg *sync.WaitGroup, producedCount *int32, numMessages int, producerID int) {
	defer wg.Done()
	for m := 0; m < numMessages; m++ {
		nq.Push(Notification{UserID: producerID, Message: "test"})
		atomic.AddInt32(producedCount, 1)
	}
}

func consumer(nq *NotificationQueue, wg *sync.WaitGroup, consumedCount *int32) {
	defer wg.Done()
	for {
		_, ok := nq.Pop()
		if !ok {
			return // Exit the loop if the queue is closed and empty
		}
		atomic.AddInt32(consumedCount, 1)
	}
}

func TestNotificationQueueConcurrent(t *testing.T) {
	nq := NewNotificationQueue()
	var numProducers, numConsumers int = 3, 3
	var numMessages int = 5

	var wgProducers, wgConsumers sync.WaitGroup
	var producedCount, consumedCount int32

	// Start producers
	wgProducers.Add(numProducers)
	for p := 0; p < numProducers; p++ {
		go producer(nq, &wgProducers, &producedCount, numMessages, p)
	}

	// Start consumers
	wgConsumers.Add(numConsumers)
	for c := 0; c < numConsumers; c++ {
		go consumer(nq, &wgConsumers, &consumedCount)
	}

	// Wait for all producers to finish
	wgProducers.Wait()

	// Once all producers have finished, close the queue
	nq.Close()

	// Now wait for all consumers to finish after closing the queue
	wgConsumers.Wait()

	// Check if all produced messages were consumed
	if producedCount != consumedCount {
		t.Errorf("Mismatch in produced and consumed counts: produced %d, consumed %d", producedCount, consumedCount)
	}
}
