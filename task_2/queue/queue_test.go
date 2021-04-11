package queue

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue_Enqueue(t *testing.T) {
	assert := assert.New(t)
	actual := NewQueue(5)
	actual.Enqueue(0)
	actual.Enqueue(1)
	actual.Enqueue(2)
	actual.Enqueue(3)
	actual.Enqueue(4)
	expected := make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		expected = append(expected, i)
		n, _ := actual.GetAtPos(i)
		actualItem, _ := n.item.(int)
		assert.Equal(actualItem, expected[i], fmt.Sprintf("%v and %v not equal, but expected", actualItem, expected[i]))
	}
}

func TestQueue_Dequeue(t *testing.T) {
	assert := assert.New(t)
	actual := NewQueue(5)
	actual.Enqueue(0)
	actual.Enqueue(1)
	actual.Enqueue(2)
	actual.Enqueue(3)
	actual.Enqueue(4)
	actual.Dequeue()
	expected := make([]int, 0, 5)
	for i := 1; i < 5; i++ {
		expected = append(expected, i)
	}
	for i, v := range expected {
		n, _ := actual.GetAtPos(i)
		actualItem, _ := n.item.(int)
		assert.Equal(v, actualItem, fmt.Sprintf("%v and %v not equal, but expected", actualItem, v))
	}
}

func TestQueue_Peek(t *testing.T) {
	assert := assert.New(t)
	queue := NewQueue(5)
	queue.Enqueue(0)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	actual, _ := queue.Peek()
	expected := 0
	assert.Equal(actual, expected, fmt.Sprintf("%v and %v not equal, but expected", actual, expected))
}

func TestQueue_Sort(t *testing.T) {
	assert := assert.New(t)
	actual := NewQueue(5)
	actual.Enqueue(0)
	actual.Enqueue(2)
	actual.Enqueue(1)
	actual.Enqueue(3)
	actual.Enqueue(4)
	actual.Sort()
	expected := make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		expected = append(expected, i)
		n, _ := actual.GetAtPos(i)
		actualItem, _ := n.item.(int)
		assert.Equal(actualItem, expected[i], fmt.Sprintf("%v and %v not equal, but expected", actualItem, expected[i]))
	}
}

func TestQueue_IsEmpty(t *testing.T) {
	assert := assert.New(t)
	queue := NewQueue(5)
	actual := queue.IsEmpty()
	expected := true
	assert.Equal(actual, expected, fmt.Sprintf("%v and %v not equal, but expected", actual, expected))
}

func TestQueue_IsFull(t *testing.T) {
	assert := assert.New(t)
	queue := NewQueue(5)
	queue.Enqueue(0)
	queue.Enqueue(2)
	queue.Enqueue(1)
	queue.Enqueue(3)
	queue.Enqueue(4)
	actual := queue.IsFull()
	expected := true
	assert.Equal(actual, expected, fmt.Sprintf("%v and %v not equal, but expected", actual, expected))
}

func TestQueue_Len(t *testing.T) {
	assert := assert.New(t)
	queue := NewQueue(5)
	queue.Enqueue(0)
	queue.Enqueue(2)
	queue.Enqueue(1)
	queue.Enqueue(3)
	queue.Enqueue(4)
	actual := queue.Len
	slice := make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		slice = append(slice, i)
	}
	expected := len(slice)
	assert.Equal(actual, expected, fmt.Sprintf("%v and %v not equal, but expected", actual, expected))
}

func TestQueue_Cap(t *testing.T) {
	assert := assert.New(t)
	queue := NewQueue(5)
	actual := queue.Cap
	slice := make([]int, 0, 5)
	expected := cap(slice)
	assert.Equal(actual, expected, fmt.Sprintf("%v and %v not equal, but expected", actual, expected))
}

func TestQueue_Enqueue_Error(t *testing.T) {
	assert := assert.New(t)
	queue := NewQueue(5)
	queue.Enqueue(0)
	queue.Enqueue(2)
	queue.Enqueue(1)
	queue.Enqueue(3)
	queue.Enqueue(4)
	actual := queue.Enqueue(5)
	expected := fmt.Errorf("the queue is full in Enqueue")
	assert.Equal(actual, expected, fmt.Sprintf("%v and %v not equal, but expected", actual, expected))
}

func TestQueue_Dequeue_Error(t *testing.T) {
	assert := assert.New(t)
	queue := NewQueue(5)
	actual := queue.Dequeue()
	expected := fmt.Errorf("the queue is empty in Dequeue")
	assert.Equal(actual, expected, fmt.Sprintf("%v and %v not equal, but expected", actual, expected))
}

func TestQueue_Peek_Error(t *testing.T) {
	assert := assert.New(t)
	queue := NewQueue(5)
	_, actual := queue.Peek()
	expected := fmt.Errorf("the queue is empty in Peek")
	assert.Equal(actual, expected, fmt.Sprintf("%v and %v not equal, but expected", actual, expected))
}

func TestQueue_GetAtPos_Error(t *testing.T) {
	assert := assert.New(t)
	queue := NewQueue(5)
	queue.Enqueue(0)
	queue.Enqueue(2)
	queue.Enqueue(1)
	queue.Enqueue(3)
	queue.Enqueue(4)
	_, actual := queue.GetAtPos(10)
	expected := fmt.Errorf("wrong entered id in GetAtPos")
	assert.Equal(actual, expected, fmt.Sprintf("%v and %v not equal, but expected", actual, expected))
}