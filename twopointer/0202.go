package twopointer

// 面试题 02.02. 返回倒数第 k 个节点
// 思路：快慢指针
// time O(N) space O(1)
func kthToLast(head *ListNode, k int) int {
	if head == nil || k == 0 {
		return -1
	}
	// 让快指针先都 k 步
	// 然后快慢指针一起走
	// 快指针到达链表末尾时，慢指针就是倒数第k个位置。
	slow, fast := head, head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	return slow.Val
}
