/**
* Given a linked list, uniformly shuffle the nodes. What if we want to prioritize space over time?
* */
package dailycodingproblems

import (
	"fmt"
	"math/rand"
	"testing"
)

type (
	shuffleList struct{}
	lstNode     struct {
		Val  int
		Next *lstNode
	}
)

func (shuffleList) shuffleLinkedList(head *lstNode) *lstNode {
	if head == nil {
		return nil
	}

	var nodes []*lstNode
	current := head
	for current != nil {
		nodes = append(nodes, current)
		current = current.Next
	}

	length := len(nodes)
	for index := 0; index < length; index++ {
		random_idx := rand.Intn(length-index) + index
		nodes[index].Val, nodes[random_idx].Val = nodes[random_idx].Val, nodes[index].Val
	}

	return head
}

func (shuffleList) printList(head *lstNode) {
	current := head
	for current != nil {
		fmt.Print(current.Val, " -> ")
		current = current.Next
	}

	fmt.Println("nil")
}

func TestShuffleList(t *testing.T) {
	head := &lstNode{Val: 1}
	head.Next = &lstNode{Val: 2}
	head.Next.Next = &lstNode{Val: 3}
	head.Next.Next.Next = &lstNode{Val: 4}
	head.Next.Next.Next.Next = &lstNode{Val: 5}

	sh := shuffleList{}
	shuffleHead := sh.shuffleLinkedList(head)

	var result []int
	current := shuffleHead
	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}

	original := []int{1, 2, 3, 4, 5}
	if len(result) != len(original) {
		t.Errorf("Expected length %d; got %d\n", len(original), len(result))
	}

	for _, val := range original {
		found := false
		for _, r := range result {
			if r == val {
				found = true
				break
			}
		}

		if !found {
			t.Errorf("Expected value %d not found in shuffled result", val)
		}
	}
}
