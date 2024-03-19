package main

import "fmt"

/**
* Doubly Linked List
**/
type DLLNode struct {
	data     interface{}
	previous *DLLNode
	next     *DLLNode
}

type DLL struct {
	head *DLLNode
	tail *DLLNode
	size int
}

func (ll *DLL) display() {
	currentNode := ll.head
	for currentNode != nil {
		fmt.Printf("%v --> ", currentNode.data)
		currentNode = currentNode.next
	}
	fmt.Println("nil")
}

func (ll *DLL) length() int {
	size := 0
	currentNode := ll.head

	for currentNode != nil {
		size++
		currentNode = currentNode.next
	}

	return size
}

func (ll *DLL) isEmpty() bool {
	return ll.size == 0
}

func (ll *DLL) insertAtBeginning(data interface{}) {
	newNode := &DLLNode{
		data:     data,
		previous: nil,
		next:     nil,
	}

	if ll.isEmpty() {
		ll.head = newNode
		ll.tail = newNode
	} else {
		head := ll.head
		newNode.next = head
		newNode.previous = nil

		head.previous = newNode
		ll.head = newNode
	}
	ll.size++
}

func (ll *DLL) insertAtEnd(data interface{}) {
	newNode := &DLLNode{
		data:     data,
		previous: nil,
		next:     nil,
	}

	if ll.isEmpty() {
		ll.head = newNode
		ll.tail = newNode
	} else {
		currentNode := ll.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		newNode.previous = currentNode
		newNode.next = nil

		currentNode.next = newNode
		ll.tail = newNode
		ll.size++
	}
}

func (ll *DLL) insertAt(data interface{}, position int) {
	if position > ll.length()+1 {
		fmt.Errorf("InsertAt: List index out of range")
	}
	newNode := &DLLNode{
		data:     data,
		previous: nil,
		next:     nil,
	}

	if ll.isEmpty() {
		ll.head = newNode
		ll.tail = newNode
	} else if position == 1 {
		ll.insertAtBeginning(data)
	} else if position == ll.length() {
		ll.insertAtEnd(data)
	} else {
		currentNode := ll.head
		for index := 1; index <= ll.length(); index++ {
			if index == position {
				newNode.previous = currentNode.previous
				newNode.next = currentNode

				if currentNode.previous != nil {
					currentNode.previous.next = newNode
				}
				currentNode.previous = newNode

				ll.size++
				break
			}
			currentNode = currentNode.next
		}
	}
}

func (ll *DLL) deleteFirst() (interface{}, error) {
	if ll.isEmpty() {
		return nil, fmt.Errorf("DeleteFirst: Index out or bound.")
	}

	currentNode := ll.head
	var deletedNode interface{}
	if currentNode.previous == nil {
		deletedNode = currentNode.data
		ll.head = currentNode.next
		ll.head.previous = nil

		ll.size--
	}

	return deletedNode, nil
}

func (ll *DLL) deleteLast() (interface{}, error) {
	if ll.isEmpty() {
		return nil, fmt.Errorf("DeleteLast: List is empty.")
	}

	currentNode := ll.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}

	ll.tail = currentNode.previous
	ll.tail.next = nil

	ll.size--

	return currentNode.data, nil
}

func (ll *DLL) deleteAt(position int) (interface{}, error) {
	if ll.isEmpty() {
		return nil, fmt.Errorf("DeleteAt: List is empty.")
	}

	if position > ll.length() {
		return nil, fmt.Errorf("DeleteAt: Index out of bound.")
	}

	currentNode := ll.head
	var currentNodeData interface{}
	for index := 1; index <= position; index++ {
		if position == index {
			currentNodeData = currentNode.data
			currentNode.previous.next = currentNode.next
			currentNode.next.previous = currentNode.previous

			ll.size--
			break
		}

		currentNode = currentNode.next
	}

	return currentNodeData, nil
}

func main() {
	var dll DLL
	fmt.Printf("Length of list is: %v\n", dll.length())
	dll.display()

	dll.insertAtBeginning(10)
	dll.insertAtBeginning(5)
	fmt.Printf("Length of list is: %v\n", dll.length())
	dll.display()

	dll.insertAtEnd(15)
	dll.insertAtEnd(20)
	fmt.Printf("Length of list is: %v\n", dll.length())
	dll.display()

	dll.insertAt(25, 3)
	dll.insertAt(0, 1)
	dll.insertAt(30, 6)
	fmt.Printf("Length of list is: %v\n", dll.length())
	dll.display()

	data, err := dll.deleteFirst()
	if err != nil {
		fmt.Errorf("Error while deleting first element: %v\n", err)
	} else {
		fmt.Printf("DeleteFirst: deleted %v\n", data)
	}
	fmt.Printf("Length of list is: %v\n", dll.length())
	dll.display()

	data, err = dll.deleteLast()
	if err != nil {
		fmt.Errorf("Error while deleting last element: %v\n", err)
	} else {
		fmt.Printf("DeleteLast: deleted %v\n", data)
	}
	fmt.Printf("Length of list is: %v\n", dll.length())
	dll.display()

	data, err = dll.deleteAt(2)
	if err != nil {
		fmt.Errorf("Error while deleting first element: %v\n", err)
	} else {
		fmt.Printf("DeleteAt(%v): deleted %v\n", 2, data)
	}
	fmt.Printf("Length of list is: %v\n", dll.length())
	dll.display()
}
