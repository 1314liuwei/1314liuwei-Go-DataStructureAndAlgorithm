package linkedList

type LinkedList interface {
	Add(elem int) error
	Insert(index, elem int) error
	Remove(index int) error
	Delete(elem int) error
	Pop() (int, error)
	Clear()
	Set(index, elem int) error
	Resize(size int) error
	Find(elem int) (int, error)
	Get(index int) (int, error)
	Tail() (int, error)
	Head() (int, error)
	Data() []int
	Len() int
	Cap() int
}
