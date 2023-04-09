package offer

type CQueue struct {
	A []int
	B []int
}

func Constructor() CQueue {
	return CQueue{}
}

// 加入队尾 appendTail()函数： 将数字 val 加入栈 A 即可。
func (q *CQueue) AppendTail(value int) {
	q.A = append(q.A, value)
}

func (q *CQueue) DeleteHead() int {
	// 队列为空，return -1
	if q.isEmpty() {
		return -1
	}
	// 如果栈B不为空，直接 pop 栈顶元素即可。
	if len(q.B) != 0 {
		val := q.B[len(q.B)-1]
		q.B = q.B[:len(q.B)-1]
		return val
	}
	// 如果栈B为空，则需要从栈A pop元素进入栈B
	// 再从栈B pop 栈顶元素
	for len(q.A) > 0 {
		q.B = append(q.B, q.A[len(q.A)-1])
		q.A = q.A[:len(q.A)-1]
	}
	val := q.B[len(q.B)-1]
	q.B = q.B[:len(q.B)-1]
	return val
}

func (q *CQueue) isEmpty() bool {
	return len(q.A) == 0 && len(q.B) == 0
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
