package queue

type Queue interface {
	Push(elem int) error
	Pop() (int, error)
	Resize(size int) error
	Peek() (int, error)
	Data() []int
	Empty() bool
	Len() int
	Cap() int
}
