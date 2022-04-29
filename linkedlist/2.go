package main

import "leetcode/pkg"

// 2. Add Two Numbers
// 2. 两数相加
// 思路：模拟加法
// Complexity: time O(max(m,n)) space O(max(m,n))
func addTwoNumbers(l1 *pkg.ListNode, l2 *pkg.ListNode) *pkg.ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	// 虚拟头节点
	dummyHead := &pkg.ListNode{}

	var tail *pkg.ListNode
	tail = dummyHead

	carry := 0 // 进位
	for l1 != nil || l2 != nil {
		// 如果两个链表的长度不同，默认短链的后面都是0
		// 这样的话，即在理论上认为两个链表总是长度相同
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		sum := v1 + v2 + carry // 求和
		carry = sum / 10       // 进位
		curr := sum % 10       // 当前值

		// 尾插法建立链表
		new := &pkg.ListNode{Val: curr}
		tail.Next = new
		tail = tail.Next
	}
	// 如果最后有进位，则需要生成一个新的节点
	if carry > 0 {
		tail.Next = &pkg.ListNode{Val: carry}
	}

	return dummyHead.Next
}
