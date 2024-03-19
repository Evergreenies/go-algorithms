package main

import "fmt"

type CSLLNode struct {
	data interface{}
	next *CSLLNode
}

type CSLL struct {
	head *CSLLNode
	size int
}

func (ll *CSLL) length() int {
	if ll.head == nil {
		return 0
	}

	size := 1
	currentNode := ll.head
	currentNode = currentNode.next
	for currentNode != ll.head {
		currentNode = currentNode.next
		size++
	}

	return size
}

func (ll *CSLL) display() {
	currentNode := ll.head
	for index := 0; index < ll.size; index++ {
		fmt.Printf("%v --> ", currentNode.data)
		currentNode = currentNode.next
	}
	fmt.Println("...")
}

func (ll *CSLL) isEmpty() bool {
	return ll.size == 0
}

func (ll *CSLL) insertAtBeginning(data interface{}) {
	newNode := &CSLLNode{
		data: data,
		next: nil,
	}

	if ll.isEmpty() {
		ll.head = newNode
		ll.head.next = newNode
	} else {
		currentNode := ll.head
		newNode.next = currentNode

		for currentNode.next != ll.head {
			currentNode = currentNode.next
		}

		currentNode.next = newNode
		ll.head = newNode
	}

	ll.size++
}

func (ll *CSLL) insertAtEnd(data interface{}) {
	newNode := &CSLLNode{
		data: data,
		next: nil,
	}

	if ll.isEmpty() {
		ll.head = newNode
		ll.head.next = newNode
	} else {
		currentNode := ll.head
		for currentNode.next != ll.head {
			currentNode = currentNode.next
		}

		currentNode.next = newNode
		newNode.next = ll.head
	}

	ll.size++
}

func (ll *CSLL) insertAt(data interface{}, position int) {
	newNode := &CSLLNode{
		data: data,
		next: nil,
	}

	if ll.isEmpty() {
		ll.head = newNode
		ll.head.next = newNode
		ll.size++
	} else {
		currentNode := ll.head

		if position == 1 {
			ll.insertAtBeginning(data)
			return
		}

		count := 1
		for {
			if currentNode.next == ll.head && count < position-1 {
				break
			}

			if count == position-1 {
				newNode.next = currentNode.next
				currentNode.next = newNode
				ll.size++
				break
			}

			currentNode = currentNode.next
			count++
		}
	}
}

func (ll *CSLL) deleteFromBeginning() (interface{}, error) {
	if ll.isEmpty() {
		return nil, fmt.Errorf("DeleteFromBeginning: List is empty.")
	}

	currentNode := ll.head
	elementToDelete := currentNode.data

	if ll.size == 1 {
		ll.head = nil
		ll.size--
		return elementToDelete, nil
	}

	head := ll.head
	ll.head = currentNode.next

	for currentNode.next != head {
		currentNode = currentNode.next
	}

	currentNode.next = ll.head
	ll.size--

	return elementToDelete, nil
}

func (ll *CSLL) deleteFromEnd() (interface{}, error) {
	if ll.isEmpty() {
		return nil, fmt.Errorf("DeleteFromEnd: List is empty.")
	}

	currentNode := ll.head
	if ll.size == 1 {
		return ll.deleteFromBeginning()
	}

	for currentNode.next.next != ll.head {
		currentNode = currentNode.next
	}
	elementToDelete := currentNode.next.data
	currentNode.next = ll.head
	ll.size--

	return elementToDelete, nil
}

func (ll *CSLL) deleteAt(position int) (interface{}, error) {
	if ll.isEmpty() {
		return nil, fmt.Errorf("DeleteAt: List is empty.")
	}

	if position > ll.size {
		return nil, fmt.Errorf("DeleteAt: List index out of range.")
	}

	if ll.size == 1 && position == 1 {
		return ll.deleteFromBeginning()
	} else if ll.size == position {
		return ll.deleteFromEnd()
	}

	currentNode := ll.head
	count := 1

	for count != position-1 {
		currentNode = currentNode.next
		count++
	}

	elementToDelete := currentNode.next.data
	currentNode.next = currentNode.next.next
	ll.size--

	return elementToDelete, nil
}

func main() {
	var cll CSLL

	fmt.Printf("Size of linked list: %v\n", cll.length())
	cll.display()

	cll.insertAtBeginning(5)
	cll.insertAtBeginning(0)
	fmt.Printf("Size of linked list: %v\n", cll.length())
	cll.display()

	cll.insertAtEnd(20)
	cll.insertAtEnd(25)
	fmt.Printf("Size of linked list: %v\n", cll.length())
	cll.display()

	cll.insertAt(10, 3)
	cll.insertAt(15, 4)
	fmt.Printf("Size of linked list: %v\n", cll.length())
	cll.display()

	_, err := cll.deleteFromBeginning()
	if err != nil {
		fmt.Errorf("Unable to delete element: %v\n", err)
	}
	fmt.Printf("Size of linked list: %v\n", cll.length())
	cll.display()

	_, err = cll.deleteFromEnd()
	if err != nil {
		fmt.Errorf("Unable to delete element: %v\n", err)
	}
	fmt.Printf("Size of linked list: %v\n", cll.length())
	cll.display()

	_, err = cll.deleteAt(3)
	if err != nil {
		fmt.Errorf("Unable to delete element: %v\n", err)
	}
	fmt.Printf("Size of linked list: %v\n", cll.length())
	cll.display()
}
