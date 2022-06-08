package stack

import "math"

// 155. Min Stack
// 155. 最小栈
// 策略：辅助栈
// time O(1) space O(N)
type MinStack struct {
	stack    []int
	minStack []int
}

func New() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{math.MaxInt32},
	}
}

func (m *MinStack) Push(val int) {
	m.stack = append(m.stack, val)
	mtop := m.minStack[len(m.minStack)-1]
	m.minStack = append(m.minStack, min(mtop, val))
}

func (m *MinStack) Pop() {
	m.stack = m.stack[:len(m.stack)-1]
	m.minStack = m.minStack[:len(m.minStack)-1]
}

func (m *MinStack) Top() int {
	return m.stack[len(m.stack)-1]
}

func (m *MinStack) GetMin() int {
	return m.minStack[len(m.minStack)-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
