package stack

import "testing"

func TestStack_Push(t *testing.T) {
	expectedSize := 1

	stack := Stack[int]{}

	stack.Push(1)
	if len(stack.elements) != expectedSize {
		t.Errorf("expected stack length of %d, got %d", expectedSize, len(stack.elements))
	}
}

func TestStack_Pop(t *testing.T) {
	stack := Stack[int]{}
	stack.Push(1)
	stack.Push(2)

	element1, err := stack.Pop()

	if err != nil {
		t.Errorf("Pop() should not return an error for a non-empty stack")
	}

	if *element1 != 2 {
		t.Errorf("Expected popped element to be 2, got %d", element1)
	}

	if len(stack.elements) != 1 {
		t.Errorf("Expected stack length to be 1 after one pop, got %d", len(stack.elements))
	}

}

func TestStack_Pop_EmptyStack(t *testing.T) {
	stack := Stack[int]{}

	_, err := stack.Pop()

	if err == nil {
		t.Errorf("Expected an error when popping from an empty stack, but got nil")
	}

	expectedErrMsg := "stack is empty"
	if err.Error() != expectedErrMsg {
		t.Errorf("Expected error message '%s', got '%s'", expectedErrMsg, err.Error())
	}

}

func TestStack_IsEmpty(t *testing.T) {

	emptyStack := Stack[int]{}

	//Test when the stack is empty
	if !emptyStack.IsEmpty() {
		t.Errorf("IsEmpty() should be true for an empty stack")
	}

	//Test when the stack is not empty
	nonEmptyStack := Stack[int]{}
	nonEmptyStack.Push(1)
	if nonEmptyStack.IsEmpty() {
		t.Errorf("IsEmpty() should be false for a non-empty stack")
	}
}

func TestStack_Peek(t *testing.T) {
	stack := Stack[int]{}
	stack.Push(1)
	element, _ := stack.Peek()
	if stack.IsEmpty() {
		t.Errorf("Peek() should not remove from the stack")
	}

	if *element != 1 {
		t.Errorf("Peek() should return the expected element from top of stack")
	}
}

func TestStack_Peek_EmptyStack(t *testing.T) {
	stack := Stack[int]{}

	_, err := stack.Peek()

	if err == nil {
		t.Errorf("Expected an error when popping from an empty stack, but got nil")
	}

	expectedErrMsg := "stack is empty"
	if err.Error() != expectedErrMsg {
		t.Errorf("Expected error message '%s', got '%s'", expectedErrMsg, err.Error())
	}

}

func TestStack_Size(t *testing.T) {
	stack := Stack[int]{}

	size, err := stack.Size()

	if err != nil {
		t.Errorf("Did not expect an error when getting the size of an empty stack: %s", err)
	}

	if size != 0 {
		t.Errorf("Expected size to be 0 for an empty stack, got %d", size)
	}

	stack.Push(1)
	stack.Push(2)

	size, err = stack.Size()

	if err != nil {
		t.Errorf("Unexpected error when getting the size of the stack: %s", err)
	}

	if size != 2 {
		t.Errorf("Expected size to be 2 after pushing two elements, but got %d", size)
	}
}
