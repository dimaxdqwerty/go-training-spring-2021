package queue

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
	Len  int   // the Len of the queue
	Cap  int   // the max capacity of the queue
}

// More compares two items.
// Returns true if i1 > i2. Returns false if i1 < i2.
func More(i1, i2 interface{}) bool {
	switch i1.(type) {
	case int:
		return i1.(int) > i2.(int)
	case float64:
		return i1.(float64) > i2.(float64)
	case string:
		return i1.(string) > i2.(string)
	default:
		fmt.Println("More: unknown type")
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
	Q.Len++
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
	Q.Len--
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
	if id < 0 || id >=  Q.Len {
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
	for i := 1; i < Q.Len; i++ {
		x := Q.GetAtPos(i).item
		j := i
		for ; j >= 1 && More(Q.GetAtPos(j-1).item, x); j-- {
			Q.GetAtPos(j).item = Q.GetAtPos(j-1).item
		}
		Q.GetAtPos(j).item = x
	}
}

// IsEmpty checks if the queue is empty
func (Q *Queue) IsEmpty() bool {
	return Q.Len == 0
}

// IsFull checks if the queue is full
func (Q *Queue) IsFull() bool {
	return Q.Cap <= Q.Len
}

// NewQueue creates a queue with entered Cap
func NewQueue(size int) *Queue {
	return &Queue{
		Cap: size,
	}
}
