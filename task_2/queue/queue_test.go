package queue

import (
	"fmt"

	"github.com/stretchr/testify/assert"
	"testing"
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
		actualNum := actual.GetAtPos(i).item
		assert.Equal(expected[i], actualNum, fmt.Sprintf("%d and %d not equal, but expected", actualNum, expected[i]))
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
		actualNum := actual.GetAtPos(i).item.(int)
		assert.Equal(v, actualNum, fmt.Sprintf("%d and %d not equal, but expected", actualNum, v))
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
	actual := queue.Peek()
	expected := 0
	assert.Equal(actual, expected, fmt.Sprintf("%d and %d not equal, but expected", actual, expected))
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
		actualNum := actual.GetAtPos(i).item
		assert.Equal(expected[i], actualNum, fmt.Sprintf("%d and %d not equal, but expected", actualNum, expected[i]))
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
	actual := queue.len
	slice := make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		slice = append(slice, i)
	}
	expected := len(slice)
	assert.Equal(actual, expected, fmt.Sprintf("%d and %d not equal, but expected", actual, expected))
}

func TestQueue_Cap(t *testing.T) {
	assert := assert.New(t)
	queue := NewQueue(5)
	actual := queue.size
	slice := make([]int, 0, 5)
	expected := cap(slice)
	assert.Equal(actual, expected, fmt.Sprintf("%d and %d not equal, but expected", actual, expected))
}