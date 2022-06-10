package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	q := New(10)
	var actual []int
	for i := 0; i < 10; i++ {
		err := q.Push(i)
		actual = append(actual, i)
		assert.Nil(t, err)
	}

	data := q.Data()
	assert.Equal(t, data, actual)

	err := q.Push(10)
	assert.Error(t, err)
}

func TestPop(t *testing.T) {
	q := New(10)
	var actual []int
	for i := 0; i < 10; i++ {
		err := q.Push(i)
		actual = append(actual, i)
		assert.Nil(t, err)
	}

	data, err := q.Pop()
	assert.Equal(t, data, actual[9])
	assert.Nil(t, err)

	assert.Equal(t, q.Data(), actual[:9])
}

func TestPeek(t *testing.T) {
	q := New(10)
	var actual []int
	for i := 0; i < 10; i++ {
		err := q.Push(i)
		actual = append(actual, i)
		assert.Nil(t, err)
	}

	data, err := q.Peek()
	assert.Equal(t, data, actual[9])
	assert.Nil(t, err)

	assert.Equal(t, q.Data(), actual)
}
