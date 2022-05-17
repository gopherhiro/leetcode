package linkedlist

import "math/rand"

/********* 如何在无限序列中随机抽取元素 ********/

// 382. Linked List Random Node
// 382. 链表随机节点
// 思路：水塘抽样算法
// time O(N) space O(1)
type List struct {
	head *ListNode
}

func ConstructorB(head *ListNode) List {
	return List{
		head: head,
	}
}

func (l *List) GetRandom() int {
	var ans int
	i := 0
	p := l.head
	for p != nil {
		i++
		// 生成 [0,i) 之间整数
		// 该整数等于0的概率是 1/i
		if rand.Intn(i) == 0 {
			ans = p.Val
		}
		p = p.Next
	}
	return ans
}
