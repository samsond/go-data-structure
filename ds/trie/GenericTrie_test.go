package trie

import (
	"slices"
	"sort"
	"testing"
)

var jobTitles = []string{"Software Engineer", "Senior Software Engineer", "Software Developer", "System Analyst", "Sales Manager"}

var testCases = []struct {
	prefix          string
	expectedResults []string
}{
	{"Softw", []string{"Software Engineer", "Software Developer"}},
	{"Senior", []string{"Senior Software Engineer"}},
	{"XYZ", []string{}},
}

func TestAutocomplete(t *testing.T) {
	trie := NewTrie[struct{}]()

	for _, title := range jobTitles {
		trie.Insert(title, struct{}{})
	}

	for _, tc := range testCases {
		results := Autocomplete(trie, tc.prefix)
		sort.Strings(results)
		sort.Strings(tc.expectedResults)
		if !slices.Equal(results, tc.expectedResults) {
			t.Errorf("Autocomplete(%q) = %v; want %v", tc.prefix, results, tc.expectedResults)
		}
	}
}

func TestTrie_Search(t *testing.T) {
	trie := NewTrie[struct{}]()

	for _, title := range jobTitles {
		trie.Insert(title, struct{}{})
	}

	for _, title := range jobTitles {
		result, found := trie.Search(title)

		if !found {
			t.Errorf("Search(%q) = got %v", title, result)
		}
	}
}
