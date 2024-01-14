package linkedlist

import (
	"testing"
)

var testCases = []struct {
	index int
	value string
}{
	{index: 0, value: "G"},
	{index: 1, value: "O"},
	{index: 2, value: "L"},
	{index: 3, value: "A"},
	{index: 4, value: "N"},
	{index: 5, value: "G"},
}

func TestDoublyLinkedList_InsertElements(t *testing.T) {

	expectedOrder := []string{"G", "O", "L", "A", "N", "G"}

	dl := &DoublyLinkedList[string]{}

	err := dl.InsertElements(testCases)
	if err != nil {
		t.Fatalf("InsertElements failed: %v", err)
	}

	// Check if the elements are in the expected order
	currentNode := dl.head
	for i, expectedValue := range expectedOrder {
		if currentNode == nil || currentNode.value != expectedValue {
			t.Errorf("Expected node %d to have value %v, got %v", i, expectedValue, currentNode.value)
		}
		currentNode = currentNode.next
	}

	if dl.size != len(expectedOrder) {
		t.Errorf("Expected list size %d, got %d", len(expectedOrder), dl.size)
	}

}

func TestDoublyLinkedList_PeekFirst(t *testing.T) {
	dlEmpty := &DoublyLinkedList[string]{}
	_, err := dlEmpty.PeekFirst()
	if err == nil {
		t.Error("PeekFirst should fail on an empty list, but it did not")
	}

	// Test with non-empty list
	expectedValue := "G"
	dl := &DoublyLinkedList[string]{}
	err = dl.InsertElements(testCases)
	if err != nil {
		t.Fatalf("Failed to insert elements: %v", err)
	}

	val, err := dl.PeekFirst()
	if err != nil {
		t.Fatalf("Failed to peek first element: %v", err)
	}

	if val != expectedValue {
		t.Errorf("Expected value %s at first position, got %s", expectedValue, val)
	}

}

func TestDoublyLinkedList_PeekLast(t *testing.T) {
	dlEmpty := &DoublyLinkedList[string]{}
	_, err := dlEmpty.PeekFirst()
	if err == nil {
		t.Error("PeekLast should fail on an empty list, but it did not")
	}

	// Test with non-empty list
	expectedValue := "G"
	dl := &DoublyLinkedList[string]{}
	err = dl.InsertElements(testCases)
	if err != nil {
		t.Fatalf("Failed to insert elements: %v", err)
	}

	val, err := dl.PeekLast()
	if err != nil {
		t.Fatalf("Failed to peek last element: %v", err)
	}

	if val != expectedValue {
		t.Errorf("Expected value %s at last position, got %s", expectedValue, val)
	}

}
