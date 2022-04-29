package linkedlist

import "leetcode/pkg"

// 21. Merge Two Sorted Lists
// 21. 合并两个有序链表
// 策略：迭代
// 思路：双指针
// time O(M+N) space O(1)
func mergeTwoLists(l1 *pkg.ListNode, l2 *pkg.ListNode) *pkg.ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	dummy := &pkg.ListNode{} // 虚拟头节点
	p, p1, p2 := dummy, l1, l2
	for p1 != nil && p2 != nil {
		// p 链接到值较小的节点
		if p1.Val < p2.Val {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
		}
		// 继续后续链接
		p = p.Next
	}
	if p1 != nil {
		p.Next = p1
	}
	if p2 != nil {
		p.Next = p2
	}
	return dummy.Next
}
