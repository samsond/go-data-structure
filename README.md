## Stack Implementation in Go

This repository contains a generic stack implementation in Go, located in the ds/stack directory. The stack is implemented using Go generics, allowing it to be used with any data type.

### Directory structure

```
ds/stack/
├── stack.go      # Stack implementation
└── stack_test.go # Test file for the stack
```

### Implementation details

The stack is a simple LIFO (Last In, First Out) data structure. It supports basic stack operations:

```
1. Push: Add an element to the top of the stack.
2. Pop: Remove and return the top element of the stack. Returns an error if the stack is empty.
3. Peek: Return the top element of the stack without removing it. Returns an error if the stack is empty.
4 IsEmpty: Check if the stack is empty.
5 Size: Return the size of the stack.
```

### Usage

To use the stack, first import the package in your Go file:

`import "path/to/your/project/ds/stack"`

Then,
```
func main() {
	s := stack.Stack[int]{}
	s.Push(42)
	value, err := s.Pop()
	if err != nil {
		// Handle error
	}
	fmt.Println("value =", value)
}
```

## Running Tests

`go test ./ds/stack`

## Running with code coverage
`go test -coverprofile=coverage.out ./...`

## view the coverage in html format

``go tool cover -html=coverage.out
``

## Contributions

Contributions to this stack implementation are welcome. Please ensure that any enhancements are accompanied by corresponding tests.
