package stack

import "fmt"

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	x := obj.Pop()
	fmt.Println(x)
}

type MyStack struct {
	queue []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (m *MyStack) Push(x int) {
	m.queue = append(m.queue, x)
}

func (m *MyStack) Pop() int {
	v := m.queue[len(m.queue)-1]
	m.queue = m.queue[:len(m.queue)-1]
	return v
}

func (m *MyStack) Top() int {
	return m.queue[len(m.queue)-1]
}

func (m *MyStack) Empty() bool {
	return len(m.queue) == 0
}
