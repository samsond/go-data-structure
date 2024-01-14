package linkedlist

import (
	"errors"
	"fmt"
	"sort"
)

type Node[T any] struct {
	value      T
	next, prev *Node[T]
}

type DoublyLinkedList[T any] struct {
	size int
	head *Node[T]
	tail *Node[T]
}

func (l *DoublyLinkedList[T]) InsertAt(index int, v T) error {
	if index < 0 || index > l.size {
		return errors.New("index out of bounds")
	}

	newNode := &Node[T]{value: v}

	if l.size == 0 { // The list is empty.
		l.head = newNode
		l.tail = newNode
	} else if index == 0 { // Inserting at the beginning.
		newNode.next = l.head
		l.head.prev = newNode
		l.head = newNode
	} else if index == l.size { // Inserting at the end.
		newNode.prev = l.tail
		l.tail.next = newNode
		l.tail = newNode
	} else { // Inserting in the middle.
		prevNode := l.head
		for i := 0; i < index-1; i++ {
			prevNode = prevNode.next
		}

		newNode.next = prevNode.next

		// If inserting before the last node, update the next node's previous pointer.
		if prevNode.next != nil {
			prevNode.next.prev = newNode
		}

		// Update the new node's previous pointer and the previous node's next pointer.
		newNode.prev = prevNode
		prevNode.next = newNode

		// If inserting at the position just before the last, update the tail.
		if newNode.next == nil {
			l.tail = newNode
		}
	}

	l.size++
	return nil
}

func (l *DoublyLinkedList[T]) PeekFirst() (T, error) {
	if l.size == 0 {
		var zeroValue T
		return zeroValue, errors.New("list is empty")
	}
	return l.head.value, nil
}

func (l *DoublyLinkedList[T]) PeekLast() (T, error) {
	if l.size == 0 {
		var zeroValue T
		return zeroValue, errors.New("list is empty")
	}
	return l.tail.value, nil
}

// InsertSortedElements adds multiple elements to the doubly linked list in a sorted manner.
func (l *DoublyLinkedList[T]) InsertSortedElements(elements []struct {
	index int
	value T
}) error {

	if len(elements) == 0 {
		return nil // No elements to add
	}

	// Sort elements by index
	sort.Slice(elements, func(i, j int) bool {
		return elements[i].index < elements[j].index
	})

	for _, element := range elements {
		if err := l.InsertAt(element.index, element.value); err != nil {
			return fmt.Errorf("error inserting element at index %d: %w", element.index, err)
		}
	}

	return nil
}

// InsertElements adds multiple elements to the doubly linked list. This method does not assume any
// specific order of the elements
func (l *DoublyLinkedList[T]) InsertElements(elements []struct {
	index int
	value T
}) error {

	for _, e := range elements {

		if e.index < 0 || e.index > l.size {
			return fmt.Errorf("index out of bounds: %d", e.index)
		}

		if err := l.InsertAt(e.index, e.value); err != nil {
			return err
		}
	}

	return nil
}
