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
## NotificationQueue

### Key Features

* **Efficient Notification Management:** The NotificationQueue struct is at the core of our system, enabling the concurrent management and distribution of notifications to our users.
* **Exclusive Access with Mutex:** To maintain data integrity and prevent race conditions, we utilize a mutex that ensures exclusive access to the queue's operations, making our system robust even under heavy load.
* **Optimized Waiting with Conditional Variables:** Conditional variables are used to signal waiting consumers when new notifications are available, effectively eliminating the need for busy-waiting and reducing CPU usage.
* **Graceful Shutdown Support:** Implementing an isClosed flag within the NotificationQueue allows us to signal when no more notifications will be pushed to the queue. This mechanism is crucial for a graceful shutdown process, ensuring that all operations are completed without leaving any consumer indefinitely waiting.
* **Reliable Notification Delivery:** We've carefully adjusted the Push and Pop methods to leverage conditional variables for efficient signaling and waiting. This ensures that all notifications are processed and delivered reliably.
* **Deadlock Prevention:** By introducing a Close method, we ensure that all waiting consumers are properly signaled to terminate, thus avoiding potential deadlocks that could arise from indefinite waits.

## Trie Implementation in Go

**Trie:** Your Efficient String Assistant

Named for its purpose, the trie (pronounced "try") is a tree-like data structure rooted in efficient string retrieval. Imagine a tree where each branch represents a letter, guiding you towards words like signposts in a forest. 

**Here's how it works:**

* **Words share branches:** Words with common prefixes follow the same path until they diverge. This saves space and speeds up searches.

* **Fast lookups:** Finding a word is like following a path in a tree—there is no need to compare the entire word at each step.

* **Prefix power:** Not only can you find specific words, but you can also find words that start with a specific prefix – perfect for autocomplete and spell-checking!

Here are some key benefits of using tries:

* **Efficient search:** Finding words is fast, especially for prefixes.

* **Memory-friendly:** Saves space by sharing common prefixes.

* **Flexible:** Can handle any alphabet and varying word lengths.

### Directory structure

```
ds/trie/
├── GenericTrie.go      # Trie implementation
└── GenericTrie_test.go # Test file for the trie
```

### Implementation details

```
1. Insert: adds a word to the trie
2. Search: returns the value associated with a word in the trie
3. Autocomplete generates a list of suggestions based on a given prefix
```


## Contributions

Contributions to this repo are welcome. Please ensure that any enhancements are accompanied by corresponding tests.
