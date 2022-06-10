package queue

import (
	"errors"
	"sync"
)

var _ Queue = new(MyQueue)

type MyQueue struct {
	lock   sync.RWMutex
	data   []int
	length int
	size   int
}

func New(size ...int) Queue {
	s := -1
	if len(size) > 0 {
		s = size[0]
	}

	return &MyQueue{
		size: s,
		data: []int{},
	}
}

func (m *MyQueue) Push(elem int) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.size != -1 && m.length+1 > m.size {
		return errors.New("the current data volume has reached the maximum")
	}

	m.data = append(m.data, elem)
	m.length++
	return nil
}

func (m *MyQueue) Pop() (int, error) {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.length <= 0 {
		return 0, errors.New("stack empty")
	}

	elem := m.data[m.length-1]
	m.data = m.data[:m.length-1]
	m.length--
	return elem, nil
}

func (m *MyQueue) Resize(size int) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	if size < m.length {
		return errors.New("size must be greater than current length")
	}

	m.size = size
	return nil
}

func (m *MyQueue) Peek() (int, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	if m.length <= 0 {
		return 0, errors.New("stack empty")
	}

	elem := m.data[m.length-1]
	return elem, nil
}

func (m *MyQueue) Data() []int {
	m.lock.RLock()
	defer m.lock.RUnlock()

	data := make([]int, m.length)
	copy(data, m.data)
	return data
}

func (m *MyQueue) Empty() bool {
	m.lock.RLock()
	defer m.lock.RUnlock()

	return m.length <= 0
}

func (m *MyQueue) Len() int {
	m.lock.RLock()
	defer m.lock.RUnlock()

	return m.length
}

func (m *MyQueue) Cap() int {
	m.lock.RLock()
	defer m.lock.RUnlock()

	return m.size
}
