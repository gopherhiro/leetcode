package linkedlist

// 92. Reverse Linked List
// 92. 反转链表 II
// 策略：迭代
// time O(N) space O(1)
func reverseBetweenIter(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if left == right {
		return head
	}

	dummy := &ListNode{}
	dummy.Next = head

	// Move to the node before the reverse start position
	pre := dummy
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	curr := pre.Next // 记录开始反转的位置
	tail := curr     // 反转后这里变成了尾部，提前记录尾部位置。

	var prev *ListNode
	for i := 0; i < right-left+1; i++ {
		tmp := curr
		curr = curr.Next
		tmp.Next = prev
		prev = tmp
	}
	// Connect the reversed part with the original list
	pre.Next = prev  // prev 是反转后的头结点
	tail.Next = curr // 反转后，curr 变成了反转部分的下一个节点了。

	return dummy.Next
}
