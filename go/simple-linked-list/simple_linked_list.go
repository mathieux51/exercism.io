package linkedlist

type Element struct {
	data int
	next *Element
}

type List struct {
	head *Element
	size int
}

func New(arr []int) *List {
	return &List{size: len(arr)}
}
func (l *List) Size() int {
	// return len(l.size)
	return 0
}

func (l *List) Push(int) {}

func (l *List) Pop() (int, error) {
	return 0, nil
}
func (l *List) Array() []int {
	return []int{}
}
func (l *List) Reverse() *List {
	return l
}
