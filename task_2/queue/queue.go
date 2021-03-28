package main

import (
	"fmt"
)

type Node struct {
	item interface{}
	next *Node
}

type Queue struct {
	head *Node
	tail *Node
	len int
}

func NewQueue() *Queue {
	q := &Queue{}
	return q
}

func (Q *Queue) Enqueue(item interface{}) {
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

func (Q *Queue) Dequeue() {
	Q.head = Q.head.next
	if Q.head == nil {
		Q.tail = nil
	}
	Q.len--
}

func (Q *Queue) Peek() interface{} {
	n := Q.head
	return n.item
}

func main() {
	queue := NewQueue()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)
	fmt.Println(queue.Peek())
	queue.Dequeue()
	fmt.Println(queue.Peek())
}
