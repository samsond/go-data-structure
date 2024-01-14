package trie

// Node represents a node in the trie.
type Node[T any] struct {
	Children map[rune]*Node[T]
	IsEnd    bool
	Value    T
}

// NewTrieNode creates a new Trie node.
func NewTrieNode[T any]() *Node[T] {
	return &Node[T]{
		Children: make(map[rune]*Node[T]),
	}
}

// Trie represents the trie and has a root Trie Node
type Trie[T any] struct {
	Root *Node[T]
}

// NodeAndWord Keeps node and word in sync for accurate suggestions
type NodeAndWord[T any] struct {
	node *Node[T]
	word string
}

// NewTrie creates a new Trie.
func NewTrie[T any]() *Trie[T] {
	return &Trie[T]{Root: NewTrieNode[T]()}
}

// Insert adds a word to the trie.
func (t *Trie[T]) Insert(word string, value T) {
	node := t.Root
	for _, ch := range word {
		if node.Children[ch] == nil {
			node.Children[ch] = NewTrieNode[T]()
		}
		node = node.Children[ch]
	}
	node.IsEnd = true
	node.Value = value
}

// Search returns the value associated with a word in the trie
func (t *Trie[T]) Search(word string) (T, bool) {
	node := t.Root
	for _, ch := range word {
		if node.Children[ch] == nil {
			var zeroValue T
			return zeroValue, false
		}
		node = node.Children[ch]
	}
	if node.IsEnd {
		return node.Value, true
	}
	var zeroValue T
	return zeroValue, false
}

// Autocomplete generates a list of suggestions based on a given prefix
func Autocomplete[T any](trie *Trie[T], prefix string) []string {
	var results []string

	node := trie.Root

	// Traverse the trie to the end of the prefix
	for _, ch := range prefix {
		if node.Children[ch] == nil {
			return results
		}
		node = node.Children[ch]
	}

	// Use a stack for iterative DFS
	stack := []NodeAndWord[T]{{node, prefix}}

	for len(stack) > 0 {
		// Pop from stack
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// If node is end of a word, append to results
		if item.node.IsEnd {
			results = append(results, item.word)
		}

		// Push children to stack
		for ch, child := range item.node.Children {
			stack = append(stack, NodeAndWord[T]{child, item.word + string(ch)})
		}
	}

	return results
}
