package array

import (
	"errors"
	"fmt"
	"sort"
	"sync"
)

var _ Array = new(MyArray)

type MyArray struct {
	lock   sync.RWMutex
	data   []int
	length int
	size   int
}

func New(sizes ...int) Array {
	size := -1
	if len(sizes) > 0 {
		size = sizes[0]
	}

	return &MyArray{
		size:   size,
		data:   []int{},
		length: 0,
	}
}

func (m *MyArray) Add(elem int) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.size != -1 && m.length+1 > m.size {
		return errors.New("the current data volume has reached the maximum")
	}

	m.data = append(m.data, elem)
	m.length++
	return nil
}

func (m *MyArray) Insert(index, elem int) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.size >= 0 && m.length+1 > m.size {
		return errors.New("the current data volume has reached the maximum")
	}

	if index < 0 || index > m.length {
		text := fmt.Sprintf("index must be between [0:%d(length))", m.length)
		return errors.New(text)
	}

	var data []int
	data = append(data, m.data[:index]...)
	data = append(data, elem)
	data = append(data, m.data[index:]...)
	m.data = data
	m.length++
	return nil
}

func (m *MyArray) Remove(index int) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.length <= 0 {
		return errors.New("cannot delete an element from an empty array")
	}

	if index < 0 || index > m.length-1 {
		text := fmt.Sprintf("index must be between [0:%d(length))", m.length)
		return errors.New(text)
	}

	var data []int
	data = append(data, m.data[:index]...)
	data = append(data, m.data[index+1:]...)
	m.data = data
	m.length--
	return nil
}

func (m *MyArray) Delete(elem int) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.length <= 0 {
		return errors.New("cannot delete an element from an empty array")
	}

	var data []int
	for i := 0; i < m.length; i++ {
		if m.data[i] != elem {
			data = append(data, m.data[i])
			m.length--
		}
	}
	m.data = data
	return nil
}

func (m *MyArray) Pop() (int, error) {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.length <= 0 {
		return 0, errors.New("cannot delete an element from an empty array")
	}

	elem := m.data[m.length-1]
	m.data = m.data[:m.length-1]
	m.length--
	return elem, nil
}

func (m *MyArray) Clear() {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.data = []int{}
	m.length = 0
}

func (m *MyArray) Set(index, elem int) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.length <= 0 {
		return errors.New("cannot set an element from an empty array")
	}

	if index < 0 || index > m.length-1 {
		text := fmt.Sprintf("index must be between [0:%d(length))", m.length)
		return errors.New(text)
	}

	m.data[index] = elem
	return nil
}

func (m *MyArray) Resize(size int) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	if size < 0 {
		return errors.New("size must be greater than 0")
	}
	m.size = size
	return nil
}

func (m *MyArray) Get(index int) (int, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	if m.length <= 0 {
		return 0, errors.New("cannot set an element from an empty array")
	}

	if index < 0 || index > m.length-1 {
		text := fmt.Sprintf("index must be between [0:%d(length))", m.length)
		return 0, errors.New(text)
	}

	return m.data[index], nil
}

func (m *MyArray) Find(elem int) (int, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	if m.length <= 0 {
		return 0, errors.New("cannot find an element from an empty array")
	}

	for i := 0; i < m.length; i++ {
		if m.data[i] == elem {
			return i, nil
		}
	}

	return -1, nil
}

func (m *MyArray) Tail() (int, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	if m.length <= 0 {
		return 0, errors.New("cannot get an element from an empty array")
	}

	return m.data[m.length-1], nil
}

func (m *MyArray) Head() (int, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	if m.length <= 0 {
		return 0, errors.New("cannot get an element from an empty array")
	}

	return m.data[0], nil
}

func (m *MyArray) Sort() (Array, error) {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.length <= 0 {
		return &MyArray{
			size:   m.size,
			data:   []int{},
			length: 0,
		}, nil
	}

	data := make([]int, m.length)
	copy(data, m.data)

	sort.Ints(data)
	return &MyArray{
		size:   m.size,
		data:   data,
		length: m.length,
	}, nil
}

func (m *MyArray) Data() []int {
	m.lock.Lock()
	defer m.lock.Unlock()

	data := make([]int, m.length)
	copy(data, m.data)
	return data
}

func (m *MyArray) Len() int {
	return m.length
}

func (m *MyArray) Cap() int {
	return m.size
}
