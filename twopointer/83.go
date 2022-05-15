package twopointer

// 83. Remove Duplicates from Sorted List
// 83. 删除排序链表中的重复元素
// 思路：双指针
// time O(N), space O(1)
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	slow, fast := head, head
	for fast != nil {
		// 相邻的元素，重复的跳过，不重复的链接起来。
		// 即去掉了链表中的重复元素。
		if fast.Val != slow.Val {
			slow.Next = fast
			slow = slow.Next
		}
		fast = fast.Next
	}
	// 若末尾有重复元素，
	// 则断开与后面重复元素的连接。
	slow.Next = nil
	return head
}
