package stack

import (
	"errors"
	"sync"
)

var _ Stack = new(MyStack)

type MyStack struct {
	lock   sync.RWMutex
	data   []int
	length int
	size   int
}

func New(size ...int) Stack {
	s := -1
	if len(size) > 0 {
		s = size[0]
	}

	return &MyStack{
		size: s,
		data: []int{},
	}
}

func (m *MyStack) Push(elem int) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.size != -1 && m.length+1 > m.size {
		return errors.New("the current data volume has reached the maximum")
	}

	m.data = append([]int{elem}, m.data...)
	m.length++
	return nil
}

func (m *MyStack) Pop() (int, error) {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.length <= 0 {
		return 0, errors.New("stack empty")
	}

	elem := m.data[0]
	m.data = m.data[1:]
	m.length--
	return elem, nil
}

func (m *MyStack) Resize(size int) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	if size < m.length {
		return errors.New("size must be greater than current length")
	}

	m.size = size
	return nil
}

func (m *MyStack) Peek() (int, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	if m.length <= 0 {
		return 0, errors.New("stack empty")
	}

	elem := m.data[0]
	return elem, nil
}

func (m *MyStack) Data() []int {
	m.lock.RLock()
	defer m.lock.RUnlock()

	data := make([]int, m.length)
	copy(data, m.data)
	return data
}

func (m *MyStack) Empty() bool {
	m.lock.RLock()
	defer m.lock.RUnlock()

	return m.length <= 0
}

func (m *MyStack) Len() int {
	m.lock.RLock()
	defer m.lock.RUnlock()

	return m.length
}

func (m *MyStack) Cap() int {
	m.lock.RLock()
	defer m.lock.RUnlock()

	return m.size
}
