/*
* You are given a set of synonyms, such as (big, large) and (eat, consume). Using this set,
* determine if two sentences with the same number of words are equivalent.
*
* For example, the following two sentences are equivalent:
*
* "He wants to eat food."
* "He wants to consume food."
*
* Note that the synonyms (a, b) and (a, c) do not necessarily imply (b, c): consider the case
* of (coach, bus) and (coach, teacher).
*
* Follow-up: what if we can assume that (a, b) and (a, c) do in fact imply (b, c)?
* */

package dailycodingproblems

import (
	"fmt"
	"strings"
	"testing"
)

// Union-Find structure
type UnionFind struct {
	parent map[string]string
	rank   map[string]int
}

func NewUnionFind() *UnionFind {
	return &UnionFind{
		parent: make(map[string]string),
		rank:   make(map[string]int),
	}
}

func (uf *UnionFind) find(x string) string {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x]) // Path compression
	}
	return uf.parent[x]
}

func (uf *UnionFind) union(x, y string) {
	rootX := uf.find(x)
	rootY := uf.find(y)

	if rootX != rootY {
		if uf.rank[rootX] > uf.rank[rootY] {
			uf.parent[rootY] = rootX
		} else if uf.rank[rootX] < uf.rank[rootY] {
			uf.parent[rootX] = rootY
		} else {
			uf.parent[rootY] = rootX
			uf.rank[rootX]++
		}
	}
}

func (uf *UnionFind) add(x string) {
	if _, exists := uf.parent[x]; !exists {
		uf.parent[x] = x
		uf.rank[x] = 0
	}
}

func areSentencesEquivalent(synonyms [][]string, sentence1, sentence2 string) bool {
	uf := NewUnionFind()

	// Initialize Union-Find
	for _, pair := range synonyms {
		uf.add(pair[0])
		uf.add(pair[1])
		uf.union(pair[0], pair[1])
	}

	words1 := strings.Split(sentence1, " ")
	words2 := strings.Split(sentence2, " ")

	if len(words1) != len(words2) {
		return false
	}

	// Check if corresponding words are equivalent
	for i := 0; i < len(words1); i++ {
		root1, root2 := uf.find(words1[i]), uf.find(words2[i])
		if root1 != root2 {
			return false
		}
	}
	return true
}

func TestAreSentencesEquivalent(t *testing.T) {
	synonyms := [][]string{{"big", "large"}, {"eat", "consume"}}
	sentence1 := "He wants to eat food"
	sentence2 := "He wants to consume food"

	if areSentencesEquivalent(synonyms, sentence1, sentence2) {
		fmt.Println("The sentences are equivalent.")
	} else {
		fmt.Println("The sentences are not equivalent.")
	}
}
