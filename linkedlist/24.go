package linkedlist

import "leetcode/pkg"

// 24. Swap Nodes in Pairs
// 24. 两两交换链表中的节点
// 策略：递归
// time O(N) space O(N)
// 拿三个节点画图，链接清楚即可。（不要进入递归）
func swapPairsR(head *pkg.ListNode) *pkg.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = swapPairsR(next.Next) // 链接 next 后面的交换好的节点
	next.Next = head                  // next 链接 head，完成交换
	return next                       // next 即是交换好链表的头节点
}

// 24. Swap Nodes in Pairs
// 24. 两两交换链表中的节点
// 策略：迭代
// time O(N) space O(1)
func swapPairs(head *pkg.ListNode) *pkg.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 虚拟节点
	dummyHead := &pkg.ListNode{}
	dummyHead.Next = head

	tmp := dummyHead
	for tmp.Next != nil && tmp.Next.Next != nil {
		n1 := tmp.Next
		n2 := tmp.Next.Next

		tmp.Next = n2
		n1.Next = n2.Next
		n2.Next = n1

		tmp = n1
	}
	return dummyHead.Next
}

// 24. Swap Nodes in Pairs
// 24. 两两交换链表中的节点
// 策略：迭代
// time O(N) space O(1)
// 这个可读性似乎更好
func swapPairsK(head *pkg.ListNode) *pkg.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 虚拟节点
	dummyHead := &pkg.ListNode{}
	dummyHead.Next = head

	prev := dummyHead
	for prev.Next != nil && prev.Next.Next != nil {
		curr := prev.Next
		next := curr.Next

		// 交换
		prev.Next = curr.Next
		curr.Next = next.Next
		next.Next = curr

		// 继续下一个迭代
		prev = curr
	}
	return dummyHead.Next
}
