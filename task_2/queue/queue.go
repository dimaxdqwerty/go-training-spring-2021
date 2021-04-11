package queue

import (
	"fmt"
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
func More(i1, i2 interface{}) (bool, error) {
	switch i1.(type) {
	case int:
		return i1.(int) > i2.(int), nil
	case float64:
		return i1.(float64) > i2.(float64), nil
	case string:
		return i1.(string) > i2.(string), nil
	default:
		return false, fmt.Errorf("get an error when tried to compare")
	}
}

// Enqueue adds an item at the tail of the queue.
func (Q *Queue) Enqueue(item interface{}) error {
	if Q.IsFull() {
		return fmt.Errorf("the queue is full in Enqueue")
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
	return nil
}

// Dequeue deletes an item at the head of the head.
func (Q *Queue) Dequeue() error {
	if Q.IsEmpty() {
		return fmt.Errorf("the queue is empty in Dequeue")
	}
	Q.head = Q.head.next
	if Q.head == nil {
		Q.tail = nil
	}
	Q.Len--
	return nil
}

// Peek gets the item of the front of the queue without removing it.
func (Q *Queue) Peek() (interface{}, error) {
	if Q.IsEmpty() {
		return nil, fmt.Errorf("the queue is empty in Peek")
	}
	n := Q.head
	return n.item, nil
}

// GetAtPos gets a node by the id.
func (Q *Queue) GetAtPos(id int) (*Node, error) {
	if id < 0 || id >=  Q.Len {
		return nil, fmt.Errorf("wrong entered id in GetAtPos")
	}
	ptr := Q.head
	for i := 0; i < id; i++ {
		ptr = ptr.next
	}
	return ptr, nil
}

// Sort sorts the queue using insertion sort algorithm.
func (Q *Queue) Sort() error {
	node := Q.head
	for node != nil{
		nextNode := node.next
		for nextNode != nil {
			compare, err := More(node.item, nextNode.item)
			if err != nil {
				return err
			}
			if compare {
				node.item, nextNode.item = nextNode.item, node.item
			}
			nextNode = nextNode.next
		}
		node = node.next
	}
	return nil
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
