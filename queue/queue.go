package queue

type Interface interface {
	Get() interface{}
	Remove()
	Add(interface{})
	Empty() bool
}

type IntSlice []int

func (a *IntSlice) Get() int { return (*a)[0] }
func (a *IntSlice) Remove() { (*a) = (*a)[1:] }
func (a *IntSlice) Add(i int) { (*a) = append((*a), i) }
func (a *IntSlice) Empty() bool { return len(*a) == 0 }
