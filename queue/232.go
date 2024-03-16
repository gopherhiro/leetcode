package queue

// 232. 用栈实现队列
// 我们使用两个栈 s1, s2 就能实现一个队列的功能。
// 当调用 push 让元素入队时，只要把元素压入 s1 即可。
// 使用 peek 或 pop 操作队头的元素时，若 s2 为空，可以把 s1 的所有元素取出再添加进 s2，这时候 s2 中元素就是先进先出顺序了。

type MyQueue struct {
	s1 []int
	s2 []int
}

// Constructor Initialize your data structure here.
func Constructor() MyQueue {
	return MyQueue{
		s1: make([]int, 0),
		s2: make([]int, 0),
	}
}

func (m *MyQueue) Push(x int) {
	m.s1 = append(m.s1, x)
}

func (m *MyQueue) Pop() int {
	if m.Empty() {
		return -1
	}
	m.Peek()
	x := m.s2[len(m.s2)-1]
	m.s2 = m.s2[:len(m.s2)-1]
	return x
}

func (m *MyQueue) Peek() int {
	if m.Empty() {
		return -1
	}
	if len(m.s2) == 0 {
		for len(m.s1) > 0 {
			x := m.s1[len(m.s1)-1]
			m.s1 = m.s1[:len(m.s1)-1]
			m.s2 = append(m.s2, x)
		}
	}
	return m.s2[len(m.s2)-1]
}

func (m *MyQueue) Empty() bool {
	return len(m.s1) == 0 && len(m.s2) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
