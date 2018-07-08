package stack

type Interface interface {
	Push(interface{})
	Pop() interface{}
	Empty() bool
}

type IntSlice []int

func (a *IntSlice) Push(d int) { *a = append(*a, d) }
func (a *IntSlice) Pop() int {
	d := (*a)[len(*a)-1]
	*a = (*a)[:len(*a)-1]
	return d
}
func (a *IntSlice) Empty() bool { return len(*a) == 0 }
