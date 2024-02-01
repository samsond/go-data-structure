package queue

import (
	"sync"
)

type Notification struct {
	UserID  int
	Message string
}

type NotificationQueue struct {
	mu               sync.Mutex
	queue            *Queue[Notification]
	hasNotifications *sync.Cond
	isClosed         bool
}

func NewNotificationQueue() *NotificationQueue {
	nq := &NotificationQueue{
		queue: NewQueue[Notification](),
	}
	nq.hasNotifications = sync.NewCond(&nq.mu) // Use nq.mu for synchronization
	return nq
}

func (nq *NotificationQueue) Close() {
	nq.mu.Lock()
	defer nq.mu.Unlock()
	nq.isClosed = true
	// Wake up all goroutines waiting on this condition
	nq.hasNotifications.Broadcast()
}

func (nq *NotificationQueue) Push(notification Notification) {
	nq.mu.Lock()
	defer nq.mu.Unlock()
	nq.queue.Enqueue(notification)
	// Signal that the queue is not empty
	nq.hasNotifications.Signal()

}

func (nq *NotificationQueue) Pop() (Notification, bool) {
	nq.mu.Lock()
	defer nq.mu.Unlock()
	for nq.queue.IsEmpty() {
		if nq.isClosed {
			// Queue is closed, return with false
			return Notification{}, false
		}
		// Wait for a signal that a notification is available
		nq.hasNotifications.Wait()
	}

	notification, _ := nq.queue.Dequeue()

	return notification, true
}

func (nq *NotificationQueue) Size() int {
	nq.mu.Lock()
	defer nq.mu.Unlock()
	return nq.queue.Size()
}

func (nq *NotificationQueue) IsEmpty() bool {
	nq.mu.Lock()
	defer nq.mu.Unlock()
	return nq.queue.IsEmpty()
}
