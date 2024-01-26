package queue

type Node[T any] struct {
	value T
	Next  *Node[T]
}

type Queue[T any] struct {
	size  int
	Front *Node[T]
	Back  *Node[T]
}

func (q *Queue[T]) IsEmpty() bool {
	return q.Front == nil
}

func (q *Queue[T]) Size() int {
	return q.size
}

func (q *Queue[T]) Enqueue(value T) {
	newNode := &Node[T]{value: value}
	q.size++
	if q.IsEmpty() {
		q.Front = newNode
		q.Back = newNode
	} else {
		q.Back.Next = newNode
		q.Back = newNode
	}
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if q.IsEmpty() {
		var zero T // default zero value of T
		return zero, false
	}

	q.size--
	value := q.Front.value
	q.Front = q.Front.Next

	if q.Front == nil {
		q.Back = nil
	}

	return value, true
}
