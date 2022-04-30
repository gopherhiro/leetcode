package linkedlist

import (
	"leetcode/pkg"
)

// 203. Remove Linked List Elements
// 203. 移除链表元素
// 策略：迭代
// time O(N) space O(1)
func removeElementsI(head *pkg.ListNode, val int) *pkg.ListNode {
	if head == nil {
		return nil
	}
	// 虚拟节点，操作方便
	dummyHead := &pkg.ListNode{}
	dummyHead.Next = head

	// 只维持一个前驱，即可操作删除。妙啊！
	prev := dummyHead
	for prev.Next != nil {
		if prev.Next.Val == val {
			prev.Next = prev.Next.Next
		} else {
			prev = prev.Next
		}
	}
	return dummyHead.Next
}

// 203. Remove Linked List Elements
// 203. 移除链表元素
// 策略：迭代
// time O(N) space O(1)
func removeElements(head *pkg.ListNode, val int) *pkg.ListNode {
	if head == nil {
		return nil
	}
	// 虚拟节点，操作方便
	dummyHead := &pkg.ListNode{}
	dummyHead.Next = head

	prev, curr := dummyHead, head
	for curr != nil {
		if curr.Val == val {
			// 值相等，prev 走两步（即是删除curr操作）
			prev.Next = curr.Next
		} else {
			// 值不等，prev 走一步
			prev = prev.Next
		}
		// curr 每次都走一步
		curr = curr.Next
	}
	return dummyHead.Next
}
