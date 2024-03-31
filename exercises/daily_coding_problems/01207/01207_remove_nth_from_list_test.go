package dailycodingproblems

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
* Given a singly linked list and an integer k, remove the kth last element
* from the list. k is guaranteed to be smaller than the length of the list.
* The list is very long, so making more than one pass is prohibitively expensive.
*
* Do this in constant space and in one pass.
* **/

type Node struct {
	value interface{}
	next  *Node
}

func (node *Node) init(value interface{}, next *Node) *Node {
	return &Node{
		value: value,
		next:  next,
	}
}

func newNode(value interface{}, next *Node) *Node {
	return new(Node).init(value, next)
}

func (node *Node) removeKthFromEnd(kth int) *Node {
	fast := node
	slow := node

	for kth > 0 {
		if fast == nil {
			return node.next
		}

		fast = fast.next
		kth--
	}

	for fast != nil && fast.next != nil {
		fast = fast.next
		slow = slow.next
	}

	if slow != nil {
		slow.next = slow.next.next
	}

	return node
}

func (node *Node) printList() {
	current := node
	for current != nil {
		fmt.Printf("%v -> ", current.value)
		current = current.next
	}
	fmt.Println("nil")
}

func TestRemoveKthFromLast(t *testing.T) {
	node := newNode(1, newNode(2, newNode(3, newNode(4, newNode(5, nil)))))
	node.removeKthFromEnd(2)
	node.printList()

	assert := assert.New(t)
	assert.Equal(node, newNode(1, newNode(2, newNode(3, newNode(5, nil)))), "now 4 should be removed from the linked list")

	node = newNode(1, newNode(2, newNode(3, newNode(4, newNode(5, nil)))))
	node.removeKthFromEnd(4)
	node.printList()

	assert.Equal(node, newNode(1, newNode(3, newNode(4, newNode(5, nil)))), "now 2 should be removed from the linked list")
}
