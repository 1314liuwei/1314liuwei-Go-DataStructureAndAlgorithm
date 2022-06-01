package linkedList

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	list := New(10)

	var actual []int
	for i := 0; i < 10; i++ {
		actual = append(actual, i)
		err := list.Add(i)
		assert.Nil(t, err)
	}

	data := list.Data()
	assert.Equal(t, data, actual)
}

func TestInsert(t *testing.T) {
	list := New(10)

	var actual []int
	for i := 0; i < 10; i++ {
		actual = append([]int{i}, actual...)
		err := list.Insert(0, i)
		assert.Nil(t, err)
	}

	data := list.Data()
	assert.Equal(t, data, actual)
}

func TestRemove(t *testing.T) {
	list := New(10)

	var actual []int
	for i := 0; i < 10; i++ {
		actual = append([]int{i}, actual...)
		err := list.Insert(0, i)
		assert.Nil(t, err)
	}

	err := list.Remove(0)
	actual = actual[1:]
	assert.Nil(t, err)

	err = list.Remove(list.Len())
	actual = actual[:len(actual)-1]
	assert.Nil(t, err)

	data := list.Data()
	assert.Equal(t, data, actual)
}

func TestDelete(t *testing.T) {
	list := New(10)

	for i := 0; i < 10; i++ {
		err := list.Insert(0, i)
		assert.Nil(t, err)
	}

	err := list.Delete(0)
	assert.Nil(t, err)

	err = list.Delete(9)
	assert.Nil(t, err)

	err = list.Delete(5)
	assert.Nil(t, err)

	data := list.Data()
	assert.Equal(t, data, []int{8, 7, 6, 4, 3, 2, 1})
}

func TestSet(t *testing.T) {
	list := New(10)

	for i := 0; i < 10; i++ {
		err := list.Add(i)
		assert.Nil(t, err)
	}

	err := list.Set(0, -1)
	assert.Nil(t, err)

	err = list.Set(9, -1)
	assert.Nil(t, err)

	err = list.Set(5, -1)
	assert.Nil(t, err)

	data := list.Data()
	assert.Equal(t, data, []int{-1, 1, 2, 3, 4, -1, 6, 7, 8, -1})
}

func TestFind(t *testing.T) {
	list := New(10)

	for i := 0; i < 10; i++ {
		err := list.Add(i)
		assert.Nil(t, err)
	}

	index, err := list.Find(0)
	assert.Nil(t, err)
	assert.Equal(t, index, 0)

	index, err = list.Find(5)
	assert.Nil(t, err)
	assert.Equal(t, index, 6)

	index, err = list.Find(9)
	assert.Nil(t, err)
	assert.Equal(t, index, 10)

	index, err = list.Find(10)
	assert.Error(t, err)
	assert.Equal(t, index, -1)
}

func TestGet(t *testing.T) {
	list := New(10)

	for i := 0; i < 10; i++ {
		err := list.Add(i)
		assert.Nil(t, err)
	}

	index, err := list.Get(0)
	assert.Nil(t, err)
	assert.Equal(t, index, 0)

	index, err = list.Get(5)
	assert.Nil(t, err)
	assert.Equal(t, index, 4)

	index, err = list.Get(10)
	assert.Nil(t, err)
	assert.Equal(t, index, 9)

	index, err = list.Get(11)
	assert.Error(t, err)
	assert.Equal(t, index, 0)

	index, err = list.Get(-1)
	assert.Error(t, err)
	assert.Equal(t, index, 0)
}

func TestTail(t *testing.T) {
	list := New(10)
	tail, err := list.Tail()
	assert.Error(t, err)
	assert.Equal(t, tail, 0)

	for i := 0; i < 10; i++ {
		err := list.Add(i)
		assert.Nil(t, err)
	}

	tail, err = list.Tail()
	assert.Nil(t, err)
	assert.Equal(t, tail, 9)
}

func TestHead(t *testing.T) {
	list := New(10)
	tail, err := list.Head()
	assert.Error(t, err)
	assert.Equal(t, tail, 0)

	for i := 0; i < 10; i++ {
		err := list.Add(i)
		assert.Nil(t, err)
	}

	tail, err = list.Head()
	assert.Nil(t, err)
	assert.Equal(t, tail, 0)
}
