/*
* Implement an LFU (Least Frequently Used) cache. It should be able to be initialized with a
* cache size n, and contain the following methods:
*
* set(key, value): sets key to value. If there are already n items in the cache and we are
* adding a new item, then it should also remove the least frequently used item. If there is
* a tie, then the least recently used key should be removed.
* get(key): gets the value at key. If no such key exists, return null.
*
* Each operation should run in O(1) time.
* */
package dailycodingproblems

import (
	"fmt"
	"testing"
)

type Node struct {
	key, value, count int
	prev, next        *Node
}

type DoublyLinkedList struct {
	head, tail *Node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	dll := &DoublyLinkedList{
		head: &Node{},
		tail: &Node{},
	}

	dll.head.next = dll.tail
	dll.tail.prev = dll.head

	return dll
}

func (*DoublyLinkedList) remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (dll *DoublyLinkedList) add(node *Node) {
	node.next = dll.head.next
	node.prev = dll.head

	dll.head.next.prev = node
	dll.head.next = node
}

func (dll *DoublyLinkedList) isEmpty() bool {
	return dll.head.next == dll.tail
}

type LFUCache struct {
	capacity, size, minFreq int

	nodeMap map[int]*Node
	freqMap map[int]*DoublyLinkedList
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity: capacity,
		nodeMap:  make(map[int]*Node),
		freqMap:  make(map[int]*DoublyLinkedList),
	}
}

func (cache *LFUCache) updateNode(node *Node) {
	count := node.count
	cache.freqMap[count].remove(node)

	if cache.freqMap[count].isEmpty() {
		if count == cache.minFreq {
			cache.minFreq++
		}

		delete(cache.freqMap, count)
	}

	node.count++
	if cache.freqMap[node.count] == nil {
		cache.freqMap[node.count] = NewDoublyLinkedList()
	}

	cache.freqMap[node.count].add(node)
}

func (cache *LFUCache) Get(key int) int {
	if node, exists := cache.nodeMap[key]; exists {
		cache.updateNode(node)

		return node.value
	}

	return -1
}

func (cache *LFUCache) Put(key, value int) {
	if cache.capacity == 0 {
		return
	}

	if node, exists := cache.nodeMap[key]; exists {
		node.value = value
		cache.updateNode(node)

		return
	}

	if cache.size == cache.capacity {
		dll := cache.freqMap[cache.minFreq]
		delete(cache.nodeMap, dll.tail.prev.key)
		dll.remove(dll.tail.prev)
		cache.size--
	}

	node := &Node{key: key, value: value, count: 1}
	cache.nodeMap[key] = node
	if cache.freqMap[1] == nil {
		cache.freqMap[1] = NewDoublyLinkedList()
	}

	cache.freqMap[1].add(node)
	cache.minFreq = 1
	cache.size++
}

func TestDoublyLinkedList(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1)) // Output: 1
	cache.Put(3, 3)
	fmt.Println(cache.Get(2)) // Output: -1
	fmt.Println(cache.Get(3)) // Output: 3
	cache.Put(4, 4)
	fmt.Println(cache.Get(1)) // Output: -1
	fmt.Println(cache.Get(3)) // Output: 3
	fmt.Println(cache.Get(4)) // Output: 4
}
