package linkedlist

import (
	"leetcode/pkg"
)

// 206. Reverse Linked List
// 206. 反转链表
// 策略：迭代
// time O(N) space O(1)
func reverseList(head *pkg.ListNode) *pkg.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev, curr *pkg.ListNode
	for head != nil {
		curr = head      // 当前节点
		head = head.Next // 去到下一个节点
		curr.Next = prev // 反转指向
		prev = curr      // 更新前驱
	}
	return curr
}

// 206. Reverse Linked List
// 206. 反转链表
// 策略：递归
// time O(N) space O(N)
func reverseListR(head *pkg.ListNode) *pkg.ListNode {
	// base case
	if head == nil || head.Next == nil {
		return head
	}
	// 递归已经帮你反转了 head 之后的所有节点，并返回反转后的头节点 last
	// 你需要做的就是把当前节点 head 和反转后的链表链接起来即可。
	last := reverseListR(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}