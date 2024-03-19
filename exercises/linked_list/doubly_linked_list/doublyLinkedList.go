package main

import "fmt"

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
}
