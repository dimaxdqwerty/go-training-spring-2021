package main

import (
	"fmt"
	"os"
)

// Node represents node of data struct queue.
type Node struct {
	item interface{} // an item that is in the node
	next *Node // pointer to the next node
}

// Queue represents the implementation of queue.
type Queue struct {
	head *Node // the head of the queue
	tail *Node // the tail of the queue
	len int // the len of the queue
	size int // the max size of the queue
}

// more compares two items.
// Returns true if i1 > i2. Returns false if i1 < i2.
func more(i1, i2 interface{}) bool {
	switch i1.(type) {
	case int:
		return i1.(int) > i2.(int)
	case float64:
		return i1.(float64) > i2.(float64)
	case string:
		return i1.(string) > i2.(string)
	default:
		fmt.Println("more: unknown type")
		os.Exit(1)
	}
	return false
}

// Enqueue adds an item at the tail of the queue.
func (Q *Queue) Enqueue(item interface{}) {
	if Q.IsFull() {
		fmt.Println("Enqueue: the queue is full")
		os.Exit(1)
	}
	n := &Node{item: item}
	if Q.tail == nil {
		Q.tail = n
		Q.head = n
	} else {
		Q.tail.next = n
		Q.tail = n
	}
	Q.len++
}

// Dequeue deletes an item at the head of the head.
func (Q *Queue) Dequeue() {
	if Q.IsEmpty() {
		fmt.Println("Dequeue: the queue is empty")
		os.Exit(1)
	}
	Q.head = Q.head.next
	if Q.head == nil {
		Q.tail = nil
	}
	Q.len--
}

// Peek gets the item of the front of the queue without removing it.
func (Q *Queue) Peek() interface{} {
	if Q.IsEmpty() {
		fmt.Println("Peek: the queue is empty")
		os.Exit(1)
	}
	n := Q.head
	return n.item
}

// GetAtPos gets a node by the id.
func (Q *Queue) GetAtPos(id int) *Node {
	if id < 0 || id >=  Q.len {
		fmt.Println("GetAtPos: wrong entered id")
		os.Exit(1)
	}
	ptr := Q.head
	for i := 0; i < id; i++ {
		ptr = ptr.next
	}
	return ptr
}

// Sort sorts the queue using insertion sort algorithm.
func (Q *Queue) Sort() {
	for i := 1; i < Q.len; i++ {
		x := Q.GetAtPos(i).item
		j := i
		for ; j >= 1 && more(Q.GetAtPos(j-1).item, x); j-- {
			Q.GetAtPos(j).item = Q.GetAtPos(j-1).item
		}
		Q.GetAtPos(j).item = x
	}
}

// IsEmpty checks if the queue is empty
func (Q *Queue) IsEmpty() bool {
	return Q.len == 0
}

// IsFull checks if the queue is full
func (Q *Queue) IsFull() bool {
	return Q.size <= Q.len
}

// NewQueue creates a queue with entered size
func NewQueue(size int) *Queue {
	return &Queue{
		size:  size,
	}
}


func main() {
	queue := NewQueue(5)
	fmt.Println(queue.IsEmpty())
	queue.Enqueue(1)
	queue.Enqueue(5)
	queue.Enqueue(3)
	queue.Enqueue(2)
	queue.Enqueue(4)
	fmt.Println(queue.Peek())
	queue.Dequeue()
	fmt.Println(queue.Peek())
	queue.Sort()
	for i := 0; i < queue.len; i++ {
		fmt.Printf("%v ", queue.GetAtPos(i).item)
	}
	fmt.Println()
}
