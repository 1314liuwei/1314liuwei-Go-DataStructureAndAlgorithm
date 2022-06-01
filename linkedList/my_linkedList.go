package linkedList

import (
	"errors"
	"fmt"
	"sync"
)

type node struct {
	data int
	next *node
}

var _ LinkedList = new(MyLinkedList)

type MyLinkedList struct {
	lock   sync.RWMutex
	size   int
	length int
	head   *node
}

func New(sizes ...int) LinkedList {
	size := -1
	if len(sizes) > 0 {
		size = sizes[0]
	}
	return &MyLinkedList{
		size:   size,
		length: 0,
		head:   new(node),
	}
}

func (m *MyLinkedList) Add(elem int) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.size != -1 && m.length+1 > m.size {
		return errors.New("the current data volume has reached the maximum")
	}

	n := m.head
	for n.next != nil {
		n = n.next
	}

	n.next = &node{
		data: elem,
	}
	m.length++
	return nil
}

func (m *MyLinkedList) Insert(index, elem int) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.size != -1 && m.length+1 > m.size {
		return errors.New("the current data volume has reached the maximum")
	}

	if index < 0 || index > m.length {
		text := fmt.Sprintf("index must be between [0:%d(length))", m.length)
		return errors.New(text)
	}

	n := m.head
	for i := 0; i < index; i++ {
		n = n.next
	}

	pre := n.next
	n.next = &node{
		data: elem,
		next: pre,
	}
	m.length++

	return nil
}

func (m *MyLinkedList) Remove(index int) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.length-1 < 0 {
		return errors.New("you cannot remove an element from an empty linked list")
	}

	if index < 0 || index > m.length {
		text := fmt.Sprintf("index must be between [0:%d(length))", m.length)
		return errors.New(text)
	}

	n := m.head
	for i := 0; i < index-1; i++ {
		n = n.next
	}

	n.next = n.next.next
	m.length--
	return nil
}

func (m *MyLinkedList) Delete(elem int) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.length-1 < 0 {
		return errors.New("you cannot remove an element from an empty linked list")
	}

	n := m.head
	for n.next.data != elem {
		n = n.next
	}

	n.next = n.next.next
	m.length--
	return nil
}

func (m *MyLinkedList) Pop() (int, error) {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.length-1 < 0 {
		return 0, errors.New("you cannot remove an element from an empty linked list")
	}

	n := m.head
	for n.next != nil {
		n = n.next
	}

	m.length--
	return n.data, nil
}

func (m *MyLinkedList) Clear() {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.head.next = nil
	m.length = 0
}

func (m *MyLinkedList) Set(index, elem int) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	if index < 0 || index > m.length {
		text := fmt.Sprintf("index must be between [0:%d(length))", m.length)
		return errors.New(text)
	}

	n := m.head
	for i := 0; i < index; i++ {
		n = n.next
	}
	n.next.data = elem
	return nil
}

func (m *MyLinkedList) Resize(size int) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	if size < 0 {
		return errors.New("size cannot be negative")
	}

	m.size = size
	return nil
}

func (m *MyLinkedList) Find(elem int) (int, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	n := m.head
	i := 0
	flag := false
	for n != nil {
		if n.data == elem {
			flag = true
			break
		}
		n = n.next
		i++
	}

	if !flag {
		return -1, errors.New("the element was not found")
	}

	return i, nil
}

func (m *MyLinkedList) Get(index int) (int, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	if index < 0 || index > m.length {
		text := fmt.Sprintf("index must be between [0:%d(length))", m.length)
		return 0, errors.New(text)
	}

	n := m.head
	for i := 0; i < index; i++ {
		n = n.next
	}

	return n.data, nil
}

func (m *MyLinkedList) Tail() (int, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	if m.length <= 0 {
		return 0, errors.New("the linked list is empty")
	}

	n := m.head
	for n.next != nil {
		n = n.next
	}

	return n.data, nil
}

func (m *MyLinkedList) Head() (int, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	if m.length <= 0 {
		return 0, errors.New("the linked list is empty")
	}

	return m.head.next.data, nil
}

func (m *MyLinkedList) Data() []int {
	m.lock.RLock()
	defer m.lock.RUnlock()

	var data []int
	n := m.head
	for n.next != nil {
		data = append(data, n.next.data)
		n = n.next
	}
	return data
}

func (m *MyLinkedList) Len() int {
	return m.length
}

func (m *MyLinkedList) Cap() int {
	return m.size
}
