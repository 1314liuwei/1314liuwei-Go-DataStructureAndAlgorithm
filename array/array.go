package array

type Array interface {
	Add(elem int) error
	Insert(index, elem int) error
	Remove(index int) error
	Delete(elem int) error
	Pop() (int, error)
	Set(index, elem int) error
	Resize(size int) error
	Find(elem int) (int, error)
	Get(index int) (int, error)
	Tail() (int, error)
	Head() (int, error)
	Sort() (Array, error)
	Data() []int
}
