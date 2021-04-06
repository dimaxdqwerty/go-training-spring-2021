package list

import (
	"fmt"
	"os"
)

// Node represents node of data struct linked list.
type Node struct {
	item interface{} // an item that is in the node
	next *Node // pointer to the next node
}

// LinkedList represents the implementation of singly linked list.
type LinkedList struct {
	head *Node // the head of the list
	len int // the len of the list
}

// More compares two items.
// Returns true if i1 > i2. Returns false if i1 < i2
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

// Insert adds an item at the beginning of the list.
func (L *LinkedList) Insert(item interface{}) {
	n := &Node{item: item}
	if L.len == 0 {
		L.head = n
		L.len++
		return
	}
	ptr := L.head
	for i := 0; i < L.len; i++ {
		if ptr.next == nil {
			ptr.next = n
			L.len++
			return
		}
		ptr = ptr.next
	}
}

// Display displays the complete list to the console.
func (L *LinkedList) Display() {
	if L.len == 0 {
		fmt.Println("Display: the list is empty")
		os.Exit(1)
	}
	ptr := L.head
	for i := 0; i < L.len; i++ {
		fmt.Printf("%v ", ptr.item)
		ptr = ptr.next
	}
	fmt.Println()
}

// Search returns an item by the id.
func (L *LinkedList) Search(id int) interface{} {
	if id >= L.len {
		fmt.Println("Search: wrong entered id")
		os.Exit(1)
	}
	return L.GetAtPos(id).item
}

// Deletion deletes an item at the beginning of the list.
func (L *LinkedList) Deletion() {
	L.head = L.head.next
	L.len--
}

// Delete deletes an item by the id.
func (L *LinkedList) Delete(id int) {
	if id >= L.len {
		fmt.Println("Delete: wrong entered id")
		os.Exit(1)
	}
	if id == 0 {
		L.Deletion()
	} else {
		prevNode := L.GetAtPos(id-1)
		prevNode.next = L.GetAtPos(id).next
		L.len--
	}
}

// GetAtPos gets a node by the id.
func (L *LinkedList) GetAtPos(id int) *Node {
	if id < 0 || id >=  L.len {
		fmt.Println("GetAtPos: wrong entered id")
		os.Exit(1)
	}
	ptr := L.head
	for i := 0; i < id; i++ {
		ptr = ptr.next
	}
	return ptr
}

// Sort sorts the list using insertion sort algorithm.
func (L *LinkedList) Sort() {
	for i := 1; i < L.len; i++ {
		x := L.GetAtPos(i).item
		j := i
		for ; j >= 1 && More(L.GetAtPos(j-1).item, x); j-- {
			L.GetAtPos(j).item = L.GetAtPos(j-1).item
		}
		L.GetAtPos(j).item = x
	}
}

// NewLinkedList creates a list.
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}
