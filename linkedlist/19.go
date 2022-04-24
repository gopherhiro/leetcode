package main

import "leetcode/pkg"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	h := pkg.CreateListFromTail(nums)
	pkg.PrintLinkedList(h)
	head := removeNthFromEnd(h, 2)
	pkg.PrintLinkedList(head)
}

// 19. Remove Nth ListNode From End of List
// 19. 删除链表的倒数第 N 个结点
// 策略：快慢指针
// 思路：
// 1、使用快慢指针确认链表中的倒数第 k 个节点
// 2、要删除倒数第 n 个节点，先找到倒数第 n + 1 个节点
// time O(N) space O(1)
func removeNthFromEnd(head *pkg.ListNode, n int) *pkg.ListNode {
	dummy := &pkg.ListNode{}
	dummy.Next = head

	// 要删除倒数第 n 个节点，先找到倒数第 n + 1 个节点
	// 避免删除第一个节点时，空指针，此处传虚拟头节点
	p := kthFromEnd(dummy, n+1)
	p.Next = p.Next.Next // 删除倒数第 n 个节点
	return dummy.Next
}

// 返回链表的倒数第 k 个节点
func kthFromEnd(head *pkg.ListNode, k int) *pkg.ListNode {
	slow, fast := head, head
	// 快指针先走 k 步
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	// 快慢指针同时走 n - k 步
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	// slow 现在指向第 n - k 个节点，即倒数第 k 个节点
	return slow
}
