/*
* Implement an autocomplete system. That is, given a query string s and a set of all possible query strings,
* return all strings in the set that have s as a prefix.
*
* For example, given the query string de and the set of strings [dog, deer, deal], return [deer, deal].
*
* Hint: Try preprocessing the dictionary into a more efficient data structure to speed up queries.
* */

package dailycodingproblems

import "testing"

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{children: make(map[rune]*TrieNode)}
}

func NewTrie() *Trie {
	return &Trie{root: NewTrieNode()}
}

func (t *Trie) Insert(word string) {
	node := t.root

	for _, char := range word {
		if _, exists := node.children[char]; !exists {
			node.children[char] = NewTrieNode()
		}

		node = node.children[char]
	}

	node.isEnd = true
}

func (t *Trie) SearchPrefix(prefix string) *TrieNode {
	node := t.root

	for _, char := range prefix {
		if _, exists := node.children[char]; !exists {
			return nil
		}

		node = node.children[char]
	}

	return node
}

func (t *Trie) dfs(node *TrieNode, prefix string, results *[]string) {
	if node.isEnd {
		*results = append(*results, prefix)
	}

	for char, child := range node.children {
		t.dfs(child, prefix+string(char), results)
	}
}

func (t *Trie) Autocomplete(prefix string) []string {
	node := t.SearchPrefix(prefix)
	if node == nil {
		return []string{}
	}

	results := []string{}
	t.dfs(node, prefix, &results)

	return results
}

func TestAutocomplete(t *testing.T) {
	trie := NewTrie()
	words := []string{"dog", "deer", "deal"}
	for _, word := range words {
		trie.Insert(word)
	}

	results := trie.Autocomplete("de")
	t.Log(results)
}
