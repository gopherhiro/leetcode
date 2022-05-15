package twopointer

// 2095. Delete the Middle Node of a Linked List
// 2095. 删除链表的中间节点
// 思路：快慢指针
// time O(N) space O(1)
func deleteMiddle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	// 使用快慢指针，找出中间节点和其前驱节点
	slow, fast := head, head

	// 记录中间节点的前驱节点
	prev := &ListNode{}
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		prev = slow
		slow = slow.Next
	}
	// slow 是中间节点
	// prev 是中间节点的前一个节点
	// 删除中间节点
	prev.Next = slow.Next
	return head
}
