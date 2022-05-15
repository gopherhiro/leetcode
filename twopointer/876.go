package twopointer

// 876. Middle of the Linked List
// 876. 链表的中间结点
// 思路：快慢指针
// time O(N), space O(1)
func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	// 当fast到达链表末尾时，slow即是中点。
	return slow
}
