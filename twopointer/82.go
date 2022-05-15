package twopointer

// 82. Remove Duplicates from Sorted List II
// 82. 删除排序链表中的重复元素 II
// 思路：two pointer
// 1、找到重复区域，2、跳过重复区域。
// time O(N), space O(1)
func deleteDuplicatesII(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	virHead := &ListNode{}
	virHead.Next = head

	slow, fast := virHead, virHead.Next
	for fast != nil {

		// fast 指针先走，找到重复的区域。
		for fast.Next != nil && fast.Val == fast.Next.Val {
			fast = fast.Next
		}

		// 如果 slow 没有紧跟着 fast，表示中间有重复元素，需要跳跃。
		if slow.Next != fast {
			// 跳过重复区域元素，即删除重复元素。
			slow.Next = fast.Next
		} else {
			// slow 跟上 fast
			slow = slow.Next
		}

		// 继续遍历
		fast = fast.Next
	}
	return virHead.Next
}
