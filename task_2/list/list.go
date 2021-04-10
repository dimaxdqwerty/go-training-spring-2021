package list

import (
	"fmt"
)

// Node represents node of data struct linked list.
type Node struct {
	item interface{} // an item that is in the node
	next *Node // pointer to the next node
}

// LinkedList represents the implementation of singly linked list.
type LinkedList struct {
	head *Node // the head of the list
	Len  int   // the Len of the list
}

// More compares two items.
// Returns true if i1 > i2. Returns false if i1 < i2
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

// Insert adds an item at the beginning of the list.
func (L *LinkedList) Insert(item interface{}) {
	n := &Node{item: item}
	n.next = L.head
	L.head = n
	L.Len++
}

// Display displays the complete list to the console.
func (L *LinkedList) Display() error {
	if L.Len == 0 {
		return fmt.Errorf("the list is empty")
	}
	ptr := L.head
	for i := 0; i < L.Len; i++ {
		fmt.Printf("%v ", ptr.item)
		ptr = ptr.next
	}
	fmt.Println()
	return nil
}

// Search returns an item by the id.
func (L *LinkedList) Search(id int) (interface{}, error) {
	if id >= L.Len {
		return nil, fmt.Errorf("wrong entered id in Search")
	}
	n, err := L.GetAtPos(id)
	if err != nil {
		return nil, err
	}
	return n.item, nil
}

// Deletion deletes an item at the beginning of the list.
func (L *LinkedList) Deletion() {
	L.head = L.head.next
	L.Len--
}

// Delete deletes an item by the id.
func (L *LinkedList) Delete(id int) error {
	if id >= L.Len {
		return fmt.Errorf("wrong entered id in Delete")
	}
	if id == 0 {
		L.Deletion()
	} else {
		prevNode, err := L.GetAtPos(id-1)
		if err != nil {
			return err
		}
		n , err := L.GetAtPos(id)
		if err != nil {
			return err
		}
		prevNode.next = n.next
		L.Len--
	}
	return nil
}

// GetAtPos gets a node by the id.
func (L *LinkedList) GetAtPos(id int) (*Node, error) {
	if id < 0 || id >=  L.Len {
		return nil, fmt.Errorf("wrong entered id in GetAtPos")
	}
	ptr := L.head
	for i := 0; i < id; i++ {
		ptr = ptr.next
	}
	return ptr, nil
}

// Sort sorts the list using insertion sort algorithm.
func (L *LinkedList) Sort() error {
	node := L.head
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

// NewLinkedList creates a list.
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}
