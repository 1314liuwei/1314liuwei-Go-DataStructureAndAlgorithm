package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	s := New(10)
	var actual []int
	for i := 0; i < 10; i++ {
		err := s.Push(i)
		actual = append(actual, 9-i)
		assert.Nil(t, err)
	}

	data := s.Data()
	assert.Equal(t, data, actual)

	err := s.Push(10)
	assert.Error(t, err)
}

func TestPop(t *testing.T) {
	s := New(10)
	var actual []int
	for i := 0; i < 10; i++ {
		err := s.Push(i)
		actual = append(actual, 9-i)
		assert.Nil(t, err)
	}

	data, err := s.Pop()
	assert.Equal(t, data, actual[0])
	assert.Nil(t, err)
}
