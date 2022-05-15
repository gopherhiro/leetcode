package twopointer

// 141. Linked List Cycle
// 141. 环形链表
// 思路：快慢指针
// time O(N), space O(1)
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	// 快慢指针指向头节点
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		// 快指针每次前进两步
		fast = fast.Next.Next

		// 慢指针每次前进一步
		slow = slow.Next

		// 快指针、慢指针相遇，链表存在环
		if fast == slow {
			return true
		}
	}
	return false
}
