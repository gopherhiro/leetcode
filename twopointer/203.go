package twopointer

// 203. 移除链表元素
// 思路：遍历链表，查找删除即可。
// 创建虚拟头节点，方便删除。
// time O(N) space O(1)
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	// 使用虚拟节点，方便删除
	virHead := &ListNode{}
	virHead.Next = head

	prev := virHead

	for prev.Next != nil {
		if prev.Next.Val == val {
			prev.Next = prev.Next.Next
		} else {
			prev = prev.Next
		}
	}

	return virHead.Next
}
