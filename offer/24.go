package offer

// 剑指 Offer 24. 反转链表
// 206. Reverse Linked List
// 206. 反转链表
// 策略：迭代
// time O(N) space O(1)
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev, curr *ListNode // 迭代节点
	for head != nil {
		curr = head      // 当前节点
		head = head.Next // 去到下一个节点
		curr.Next = prev // 反转指向
		prev = curr      // 更新前驱
	}
	return curr
}
