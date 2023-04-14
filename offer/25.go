package offer

// 剑指 Offer 25. 合并两个排序的链表
// 21. Merge Two Sorted Lists
// 21. 合并两个有序链表
// 策略：迭代
// 思路：双指针
// time O(M+N) space O(1)
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	// 虚拟头结点
	dummy := &ListNode{}
	// p,dummy 指向同一个虚拟头结点
	// p 用于遍历，dummy 用于返回
	p, p1, p2 := dummy, l1, l2

	for p1 != nil && p2 != nil {
		// 链接较小的节点
		if p1.Val < p2.Val {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
		}
		// 继续链接
		p = p.Next
	}

	// p1 剩余链接
	if p1 != nil {
		p.Next = p1
	}
	// p2 剩余链接
	if p2 != nil {
		p.Next = p2
	}

	// 返回合并后的链表头结点
	return dummy.Next
}
