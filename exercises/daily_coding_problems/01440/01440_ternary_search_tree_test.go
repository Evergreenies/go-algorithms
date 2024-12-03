/**
A ternary search tree is a trie-like data structure where each node may have up to three children.
Here is an example which represents the words code, cob, be, ax, war, and we.

       c
    /  |  \
   b   o   w
 / |   |   |
a  e   d   a
|    / |   | \
x   b  e   r  e
The tree is structured according to the following rules:

left child nodes link to words lexicographically earlier than the parent prefix
right child nodes link to words lexicographically later than the parent prefix
middle child nodes continue the current word
For instance, since code is the first word inserted in the tree, and cob lexicographically precedes
cod, cob is represented as a left child extending from cod.

Implement insertion and search functions for a ternary search tree.
* */

package dailycodingproblems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Node struct {
	char             byte
	isEndOfWord      bool
	left, right, mid *Node
}

type TST struct {
	root *Node
}

func _search(root *Node, word string, index int) bool {
	if root == nil {
		return false
	}

	char := word[index]
	if char < root.char {
		return _search(root.left, word, index)
	} else if char > root.char {
		return _search(root.right, word, index)
	} else {
		if index+1 == len(word) {
			return root.isEndOfWord
		}

		return _search(root.mid, word, index+1)
	}
}

func (tst *TST) search(word string) bool {
	return _search(tst.root, word, 0)
}

func _insert(root *Node, word string, index int) *Node {
	if index >= len(word) {
		return root
	}

	char := word[index]
	if root == nil {
		root = &Node{char: char}
	}

	if char < root.char {
		root.left = _insert(root.left, word, index)
	} else if char > root.char {
		root.right = _insert(root.right, word, index)
	} else {
		if index+1 == len(word) {
			root.isEndOfWord = true
		} else {
			root.mid = _insert(root.mid, word, index+1)
		}
	}

	return root
}

func (tst *TST) insert(word string) {
	tst.root = _insert(tst.root, word, 0)
}

func TestTST(t *testing.T) {
	tst := &TST{}
	words := []string{"code", "cob", "be", "ax", "war", "we"}
	for _, word := range words {
		tst.insert(word)
	}

	assert := assert.New(t)

	assert.True(tst.search("code"))
	assert.True(tst.search("cob"))
	assert.True(tst.search("be"))
	assert.True(tst.search("ax"))
	assert.True(tst.search("war"))
	assert.True(tst.search("we"))
	assert.False(tst.search("not"))
}
