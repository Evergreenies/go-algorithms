package main

import "fmt"

type ListNode struct {
	data interface{}
	next *ListNode
}

type LinkedList struct {
	head *ListNode
	size int
}

func (ll *LinkedList) Display() error {
	if ll.head == nil {
		return fmt.Errorf("Display: LinkedList is empty")
	}

	current := ll.head
	for current != nil {
		fmt.Printf("%v --> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
	return nil
}

func (ll *LinkedList) Length() int {
	size := 0
	current := ll.head
	for current != nil {
		size++
		current = current.next
	}
	return size
}

func (ll *LinkedList) InsertAtBeginning(data interface{}) {
	newNode := &ListNode{
		data: data,
	}

	if ll.head == nil {
		ll.head = newNode
	} else {
		newNode.next = ll.head
		ll.head = newNode
	}

	ll.size++
}

func (ll *LinkedList) InsertAtEnd(data interface{}) {
	newNode := &ListNode{
		data: data,
	}

	if ll.head == nil {
		ll.head = newNode
	} else {
		currentNode := ll.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = newNode
	}

	ll.size++
}

func (ll *LinkedList) InsertAt(position int, data interface{}) error {
	if position < 1 || position > ll.size+1 {
		return fmt.Errorf("InsertAt: Insert out of bounds")
	}

	newNode := &ListNode{
		data: data,
		next: nil,
	}

	var previousNode, currentNode *ListNode
	previousNode = nil
	currentNode = ll.head

	for position > 1 {
		previousNode = currentNode
		currentNode = currentNode.next
		position = position - 1
	}

	if previousNode != nil {
		previousNode.next = newNode
		newNode.next = currentNode
	} else {
		newNode.next = currentNode
		ll.head = newNode
	}

	ll.size++
	return nil
}

func (ll *LinkedList) DeleteFirst() (interface{}, error) {
	if ll.head == nil {
		return nil, fmt.Errorf("DeleteFirst: LinkedList is empty.")
	}

	data := ll.head.data
	ll.head = ll.head.next

	ll.size--
	return data, nil
}

func (ll *LinkedList) DeleteLast() (interface{}, error) {
	if ll.head == nil {
		return nil, fmt.Errorf("DeleteLast: LinkedList is empty")
	}

	var previousNode *ListNode
	currentNode := ll.head

	for currentNode.next != nil {
		previousNode = currentNode
		currentNode = currentNode.next
	}

	if previousNode != nil {
		previousNode.next = nil
	} else {
		ll.head = nil
	}

	ll.size--

	return currentNode.data, nil
}

func (ll *LinkedList) DeleteAt(position int) (interface{}, error) {
	if position < 1 || position > ll.size+1 {
		return nil, fmt.Errorf("DeleteAt: Index out of bounds")
	}

	data := ll.head.data
	if position == 0 {
		ll.head = ll.head.next
		return data, nil
	}

	var previousNode, currentNode *ListNode
	previousNode = nil
	currentNode = ll.head
	index := 1

	for currentNode != nil && index != position {
		previousNode = currentNode
		currentNode = currentNode.next
		index++
	}

	if currentNode == nil {
		return nil, fmt.Errorf("DeleteAt: Index out of bounds")
	}

	data = currentNode.data
	previousNode.next = currentNode.next

	ll.size--
	return data, nil
}

func main() {
	var linkedList LinkedList
	linkedList.InsertAtBeginning(5)
	linkedList.InsertAtBeginning(3)
	linkedList.InsertAtBeginning(13)
	linkedList.InsertAtBeginning(15)

	linkedList.InsertAtEnd(45)

	linkedList.InsertAt(4, 14)
	linkedList.InsertAt(6, 16)

	err := linkedList.Display()
	if err != nil {
		fmt.Errorf("Some error while deplaying LinkedList: %v", err)
	}

	data, err := linkedList.DeleteFirst()
	if err != nil {
		fmt.Errorf(err.Error())
	} else {
		fmt.Printf("DeleteFirst: %v\n", data)
	}
	err = linkedList.Display()
	if err != nil {
		fmt.Errorf("Some error while deplaying LinkedList: %v", err)
	}

	data, err = linkedList.DeleteAt(2)
	if err != nil {
		fmt.Errorf(err.Error())
	} else {
		fmt.Printf("DeleteAt[%v]: %v\n", 2, data)
	}

	err = linkedList.Display()
	if err != nil {
		fmt.Errorf("Some error while deplaying LinkedList: %v", err)
	}

	data, err = linkedList.DeleteAt(4)
	if err != nil {
		fmt.Errorf(err.Error())
	} else {
		fmt.Printf("DeleteAt[%v]: %v\n", 4, data)
	}

	err = linkedList.Display()
	if err != nil {
		fmt.Errorf("Some error while deplaying LinkedList: %v", err)
	}

	data, err = linkedList.DeleteLast()
	if err != nil {
		fmt.Errorf(err.Error())
	} else {
		fmt.Printf("DeleteLast: %v\n", data)
	}

	fmt.Printf("Length of LinkedList is: %v\n", linkedList.Length())
	err = linkedList.Display()
	if err != nil {
		fmt.Errorf("Some error while deplaying LinkedList: %v", err)
	}
}
