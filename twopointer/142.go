package twopointer

// 142. Linked List Cycle II
// 142. 环形链表 II
// 思路：快慢指针
// time O(N) space O(1)
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	// 如果不存在环，则返回 null
	if fast == nil || fast.Next == nil {
		return nil
	}

	// 如果存在环
	// 则让其中一个指针从头节点重新出发
	slow = head
	for slow != fast {
		// 两指针以相同的速度前进
		slow = slow.Next
		fast = fast.Next
	}
	// 两指针相遇的节点，即是环的起点
	return slow
}
