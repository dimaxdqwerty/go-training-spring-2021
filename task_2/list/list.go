package main

import (
	"fmt"
)

type Node struct {
	item interface{}
	next *Node
}

type LinkedList struct {
	head *Node
	len int
}

func more(i1 interface{}, i2 interface{}) bool {
	switch i1.(type) {
	case int:
		return i1.(int) > i2.(int)
	case float64:
		return i1.(float64) > i2.(float64)
	case string:
		return i1.(string) > i2.(string)
	default:
		panic("Unknown type")
	}
}

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

func (L *LinkedList) Display() {
	if L.len == 0 {
		panic("The list is empty")
	}
	ptr := L.head
	for i := 0; i < L.len; i++ {
		fmt.Printf("%v ", ptr.item)
		ptr = ptr.next
	}
	fmt.Println()
}

func (L *LinkedList) Search(id int) interface{} {
	if id >= L.len {
		panic("id >= len")
	}
	return L.GetAtPos(id).item
}

func (L *LinkedList) Deletion() {
	L.head = L.head.next
	L.len--
}

func (L *LinkedList) Delete(id int) {
	if id >= L.len {
		panic("id >= len")
	}
	if id == 0 {
		L.Deletion()
	} else {
		prevNode := L.GetAtPos(id-1)
		prevNode.next = L.GetAtPos(id).next
		L.len--
	}
}

func (L *LinkedList) GetAtPos(id int) *Node {
	ptr := L.head
	for i := 0; i < id; i++ {
		ptr = ptr.next
	}
	return ptr
}

func (L *LinkedList) Sort() {
	for i := 1; i < L.len; i++ {
		x := L.GetAtPos(i).item
		j := i
		for ; j >= 1 && more(L.GetAtPos(j-1).item, x); j-- {
			L.GetAtPos(j).item = L.GetAtPos(j-1).item
		}
		L.GetAtPos(j).item = x
	}
}

func main() {
	list := LinkedList{}
	list.Insert(1)
	list.Insert(4)
	list.Insert(2)
	list.Insert(5)
	list.Insert(3)
	list.Display()
	fmt.Println(list.Search(3))
	list.Deletion()
	list.Display()
	list.Delete(2)
	list.Display()
	list.Sort()
	list.Display()
}
