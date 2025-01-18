/*
* Given a linked list of numbers and a pivot k, partition the linked list
* so that all nodes less than k come before nodes greater than or equal to k.
*
* For example, given the linked list 5 -> 1 -> 8 -> 0 -> 3 and k = 3, the
* solution could be 1 -> 0 -> 5 -> 8 -> 3.
* */
package DCP01486

import (
	"fmt"
	"testing"
)

type Node struct {
	Val  int
	Next *Node
}

func partition(head *Node, k int) *Node {
	var lessHead, lessTail, greaterHead, greaterTrail *Node

	for head != nil {
		next := head.Next
		head.Next = nil
		if head.Val < k {
			if lessHead == nil {
				lessHead = head
				lessTail = head
			} else {
				lessTail.Next = head
				lessTail = head
			}
		} else {
			if greaterHead == nil {
				greaterHead = head
				greaterTrail = head
			} else {
				greaterTrail.Next = head
				greaterTrail = head
			}
		}

		head = next
	}

	if lessTail != nil {
		lessTail.Next = greaterHead

		return lessHead
	}

	return greaterHead
}

func printList(head *Node) {
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Println("nil")
}

func TestPartition(*testing.T) {
	head := &Node{Val: 5}
	head.Next = &Node{Val: 1}
	head.Next.Next = &Node{Val: 8}
	head.Next.Next.Next = &Node{Val: 0}
	head.Next.Next.Next.Next = &Node{Val: 3}

	k := 3
	newHead := partition(head, k)
	printList(newHead) // Output: 1 -> 0 -> 5 -> 8 -> 3 -> nil
}
