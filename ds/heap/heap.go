package heap

type Heap[T any] struct {
	data []T
	less func(T, T) bool
}

func (h *Heap[T]) rootNode() (T, bool) {
	if len(h.data) > 0 {
		return h.data[0], true
	}

	var zero T
	return zero, false
}

func (h *Heap[T]) lastNode() (T, bool) {
	if len(h.data) > 0 {
		return h.data[len(h.data)-1], true
	}
	var zero T
	return zero, false
}

func (h *Heap[T]) leftChildIndex(index int) int {
	return (index * 2) + 1
}

func (h *Heap[T]) rightChildIndex(index int) int {
	return (index * 2) + 2
}

func (h *Heap[T]) parentIndex(index int) int {
	return (index - 1) / 2
}

func (h *Heap[T]) Insert(value T) {
	h.data = append(h.data, value)

	// keep track of the new node index
	newNodeIndex := len(h.data) - 1

	// If the new node is not in the root position
	// it's greater than its parent node:
	for newNodeIndex > 0 && h.less(h.data[h.parentIndex(newNodeIndex)], h.data[newNodeIndex]) {
		h.data[h.parentIndex(newNodeIndex)], h.data[newNodeIndex] = h.data[newNodeIndex], h.data[h.parentIndex(newNodeIndex)]

		newNodeIndex = h.parentIndex(newNodeIndex)
	}
}

func (h *Heap[T]) Delete() (T, bool) {
	if len(h.data) == 0 {
		var zero T
		return zero, false
	}

	root := h.data[0]
	lastIndex := len(h.data) - 1
	h.data[0] = h.data[lastIndex]
	h.data = h.data[:lastIndex]

	// Heapify down from the root to restore the heap property
	h.heapifyDown()

	return root, true
}

// heapifyDown adjusts the heap starting from the root to maintain the heap property
func (h *Heap[T]) heapifyDown() {
	index := 0
	for {
		leftChildIndex := h.leftChildIndex(index)
		rightChildIndex := h.rightChildIndex(index)

		// Find the optimalIndex among root, left child, and right child
		optimalIndex := index
		if leftChildIndex < len(h.data) && h.less(h.data[optimalIndex], h.data[leftChildIndex]) {
			optimalIndex = leftChildIndex
		}
		if rightChildIndex < len(h.data) && h.less(h.data[optimalIndex], h.data[rightChildIndex]) {
			optimalIndex = rightChildIndex
		}

		// Swap if the root is not the optimalIndex
		if optimalIndex != index {
			h.data[index], h.data[optimalIndex] = h.data[optimalIndex], h.data[index]
			// Move down to the optimalIndex child
			index = optimalIndex
		} else {
			break // The heap property is satisfied
		}
	}
}
