package list

import (
	"fmt"

	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedList_Insert(t *testing.T) {
	assert := assert.New(t)
	actual := NewLinkedList()
	actual.Insert(0)
	actual.Insert(1)
	actual.Insert(2)
	actual.Insert(3)
	actual.Insert(4)
	expected := make([]int, 0, 5)
	for i := 4; i >= 0; i-- {
		expected = append(expected, i)
	}
	for i, v := range expected {
		actualNum := actual.Search(i).(int)
		assert.Equal(v, actualNum, fmt.Sprintf("%d and %d not equal, but expected", actualNum, v))
	}
}

func TestLinkedList_Delete(t *testing.T) {
	assert := assert.New(t)
	actual := NewLinkedList()
	actual.Insert(0)
	actual.Insert(1)
	actual.Insert(2)
	actual.Insert(3)
	actual.Insert(4)
	actual.Delete(3)
	expected := make([]int, 0, 5)
	for i := 4; i >= 0; i-- {
		expected = append(expected, i)
	}
	expected = append(expected[:3], expected[4:]...)
	for i, v := range expected {
		actualNum := actual.Search(i).(int)
		assert.Equal(v, actualNum, fmt.Sprintf("%d and %d not equal, but expected", actualNum, v))
	}
}

func TestLinkedList_Deletion(t *testing.T) {
	assert := assert.New(t)
	actual := NewLinkedList()
	actual.Insert(0)
	actual.Insert(1)
	actual.Insert(2)
	actual.Insert(3)
	actual.Insert(4)
	actual.Deletion()
	expected := make([]int, 0, 5)
	for i := 4; i >= 0; i-- {
		expected = append(expected, i)
	}
	expected = expected[1:]
	for i, v := range expected {
		actualNum := actual.Search(i).(int)
		assert.Equal(v, actualNum, fmt.Sprintf("%d and %d not equal, but expected", actualNum, v))
	}
}

func TestLinkedList_Sort(t *testing.T) {
	assert := assert.New(t)
	actual := NewLinkedList()
	actual.Insert(4)
	actual.Insert(1)
	actual.Insert(0)
	actual.Insert(2)
	actual.Insert(3)
	actual.Sort()
	expected := make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		expected = append(expected, i)
		actualNum := actual.Search(i).(int)
		assert.Equal(expected[i], actualNum, fmt.Sprintf("%d and %d not equal, but expected", actualNum, expected[i]))
	}
}

func TestLinkedList_Len(t *testing.T) {
	assert := assert.New(t)
	list := NewLinkedList()
	list.Insert(0)
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	slice := make([]int, 0, 5)
	for i := 4; i >= 0; i-- {
		slice = append(slice, i)
	}
	actual := list.Len
	expected := len(slice)
	assert.Equal(actual, expected, fmt.Sprintf("%d and %d not equal, but expected", actual, expected))
}