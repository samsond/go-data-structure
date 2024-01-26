package tree

import (
	"slices"
	"testing"
)

func TestBinaryTree_BreadthFirst(t *testing.T) {
	// Create a BinaryTree for integers with an integer comparison function
	tree := BinaryTree[int]{
		Compare: intComparator,
	}

	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)
	tree.Insert(1)
	tree.Insert(4)
	tree.Insert(6)
	tree.Insert(8)

	result := tree.BreadthFirst()
	expected := []int{5, 3, 7, 1, 4, 6, 8}

	if !slices.Equal(result, expected) {
		t.Errorf("BreadthFirst traversal did not match expected order. Expected: %v, Got: %v", expected, result)
	}

	// Create a BinaryTree for strings with a string comparison function
	stringTree := BinaryTree[string]{
		Compare: stringComparator,
	}

	stringTree.Insert("apple")
	stringTree.Insert("banana")
	stringTree.Insert("tomato")

	resultStringSlice := stringTree.BreadthFirst()
	expectedStringSlice := []string{"apple", "banana", "tomato"}

	if !slices.Equal(resultStringSlice, expectedStringSlice) {
		t.Errorf("BreadthFirst traversal did not match expected order. Expected: %v, Got: %v", expected, result)
	}

}

func TestBinaryTree_InOrderTraversal(t *testing.T) {
	// Create a BinaryTree for integers with an integer comparison function
	tree := BinaryTree[int]{
		Compare: intComparator,
	}

	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)
	tree.Insert(1)
	tree.Insert(4)
	tree.Insert(6)
	tree.Insert(8)

	result := tree.InOrderTraversal()
	expected := []int{1, 3, 4, 5, 6, 7, 8}

	if !slices.Equal(result, expected) {
		t.Errorf("InOrder traversal did not match expected order. Expected: %v, Got: %v", expected, result)
	}

	// Create a BinaryTree for strings with a string comparison function
	stringTree := BinaryTree[string]{
		Compare: stringComparator,
	}

	stringTree.Insert("apple")
	stringTree.Insert("banana")
	stringTree.Insert("tomato")
	stringTree.Insert("avocado")

	resultStringSlice := stringTree.InOrderTraversal()
	expectedStringSlice := []string{"apple", "avocado", "banana", "tomato"}

	if !slices.Equal(resultStringSlice, expectedStringSlice) {
		t.Errorf("InOrder traversal did not match expected order. Expected: %v, Got: %v", resultStringSlice, expectedStringSlice)
	}
}

func TestBinaryTree_DepthFirstSearch(t *testing.T) {
	// Create a BinaryTree for integers with an integer comparison function
	tree := BinaryTree[int]{
		Compare: intComparator,
	}

	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)
	tree.Insert(1)
	tree.Insert(4)
	tree.Insert(6)
	tree.Insert(8)

	result := tree.DepthFirstSearch()
	expected := []int{5, 3, 1, 4, 7, 6, 8}

	if !slices.Equal(result, expected) {
		t.Errorf("InOrder traversal did not match expected order. Expected: %v, Got: %v", expected, result)
	}

	// Create a BinaryTree for strings with a string comparison function
	stringTree := BinaryTree[string]{
		Compare: stringComparator,
	}

	stringTree.Insert("apple")
	stringTree.Insert("banana")
	stringTree.Insert("tomato")
	stringTree.Insert("avocado")

	resultStringSlice := stringTree.DepthFirstSearch()
	expectedStringSlice := []string{"apple", "banana", "avocado", "tomato"}

	if !slices.Equal(resultStringSlice, expectedStringSlice) {
		t.Errorf("InOrder traversal did not match expected order. Expected: %v, Got: %v", resultStringSlice, expectedStringSlice)
	}
}
