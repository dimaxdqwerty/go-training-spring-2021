package main

import (
	"fmt"
	"sort"
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

func (Q *Queue) Len() int {
	return Q.len
}

func (Q *Queue) Less(i, j int) bool {
	i1 := Q.GetAtPos(i).item
	i2 := Q.GetAtPos(j).item
	switch i1.(type) {
	case int:
		return i1.(int) < i2.(int)
	case float64:
		return i1.(float64) < i2.(float64)
	case string:
		return i1.(string) < i2.(string)
	default:
		panic("Unknown type")
	}
}

func (Q *Queue) Swap(i, j int) {
	Q.GetAtPos(i).item, Q.GetAtPos(j).item = Q.GetAtPos(j).item, Q.GetAtPos(i).item
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

func (Q *Queue) GetAtPos(id int) *Node {
	ptr := Q.head
	for i := 0; i < id; i++ {
		ptr = ptr.next
	}
	return ptr
}

func (Q *Queue) Sort() {
	sort.Sort(Q)
}

func main() {
	queue := Queue{}
	queue.Enqueue(1)
	queue.Enqueue(5)
	queue.Enqueue(3)
	queue.Enqueue(2)
	queue.Enqueue(4)
	fmt.Println(queue.Peek())
	queue.Dequeue()
	fmt.Println(queue.Peek())
	queue.Sort()
}
