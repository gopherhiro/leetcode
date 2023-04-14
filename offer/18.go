package offer

// 剑指 Offer 18. 删除链表的节点
// 思路：定位节点 & 删除节点
// 设置虚拟节点，方便操作删除。
// time O(N) space O(1)
func deleteNode(head *ListNode, value int) *ListNode {
	if head == nil {
		return nil
	}

	// 处理头结点等于 value case
	// 可提高执行效率
	if head.Val == value {
		return head.Next
	}

	// 定义虚拟节点
	dummy := &ListNode{}
	dummy.Next = head

	// 定位节点
	prev, curr := dummy, head
	for curr != nil && curr.Val != value {
		prev = curr
		curr = curr.Next
	}
	// 删除节点
	if curr != nil {
		prev.Next = curr.Next
	}

	// 返回链表头节点
	return dummy.Next
}
