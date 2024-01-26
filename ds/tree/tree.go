package tree

import (
	"github.com/samsond/go-data-structure/ds/queue"
	"github.com/samsond/go-data-structure/ds/stack"
)

type Comparator[T any] func(a, b T) int
type Node[T any] struct {
	value       T
	Left, Right *Node[T]
}

type BinaryTree[T any] struct {
	root    *Node[T]
	Compare Comparator[T]
}

func (t *BinaryTree[T]) IsEmpty() bool {
	return t.root == nil
}

func (t *BinaryTree[T]) Insert(value T) {
	t.root = t.InsertNode(t.root, value)
}

func (t *BinaryTree[T]) InsertNode(node *Node[T], value T) *Node[T] {
	newNode := &Node[T]{value: value}

	if node == nil {
		return newNode
	}

	if t.Compare(value, node.value) < 0 {
		node.Left = t.InsertNode(node.Left, value)
	} else {
		node.Right = t.InsertNode(node.Right, value)
	}

	return node
}

func (t *BinaryTree[T]) InOrderTraversal() []T {
	var result []T
	return traverseInorder(t.root, result)
}

func traverseInorder[T any](node *Node[T], result []T) []T {

	if node == nil {
		return result
	}

	result = traverseInorder(node.Left, result)
	result = append(result, node.value)
	result = traverseInorder(node.Right, result)

	return result
}

func (t *BinaryTree[T]) BreadthFirst() []T {
	var result []T

	if t.IsEmpty() {
		return result
	}

	itemQueue := queue.NewQueue[*Node[T]]()
	itemQueue.Enqueue(t.root)

	for !itemQueue.IsEmpty() {
		node, ok := itemQueue.Dequeue()
		if !ok {
			break
		}

		result = append(result, node.value)

		if node.Left != nil {
			itemQueue.Enqueue(node.Left)
		}

		if node.Right != nil {
			itemQueue.Enqueue(node.Right)
		}
	}

	return result

}

func (t *BinaryTree[T]) DepthFirstSearch() []T {
	var result []T

	if t.IsEmpty() {
		return result
	}

	stackItem := stack.NewStack[*Node[T]]()
	stackItem.Push(t.root)

	for !stackItem.IsEmpty() {
		nodePtr, err := stackItem.Pop()

		if err != nil {
			break
		}

		node := *nodePtr
		result = append(result, node.value)

		// push right node first
		if node.Right != nil {
			stackItem.Push(node.Right)
		}

		// push left node second
		if node.Left != nil {
			stackItem.Push(node.Left)
		}
	}

	return result
}

func intComparator(a, b int) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	} else {
		return 0
	}
}

func stringComparator(a, b string) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	} else {
		return 0
	}
}
