package linkedlist

// 25. K 个一组翻转链表
// 思路：递归
// time O(n) space O(n)
func reverseKGroupN(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	a, b := head, head
	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.Next
	}
	newHead := reverse(a, b)
	a.Next = reverseKGroupN(b, k)
	return newHead
}

func reverse(a, b *ListNode) *ListNode {
	var prev, curr *ListNode
	for a != b {
		curr = a
		a = a.Next
		curr.Next = prev
		prev = curr
	}
	return curr
}
