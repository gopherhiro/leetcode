package offer

// 剑指 Offer 22. 链表中倒数第k个节点
// 思路：two pointer（快慢指针）
// time O() space O(1)
func getKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return nil
	}

	slow, fast := head, head

	// fast go k step first
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	// slow and fast go together
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	// now slow is in kth node from end
	return slow
}
