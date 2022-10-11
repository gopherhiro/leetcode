package linkedlist

import (
	"leetcode/pkg"
)

// 817. Linked List Components
// 817. 链表组件
// 思路：一次遍历 + 哈希表
// time O(n) space O(n)
func numComponents(head *pkg.ListNode, nums []int) int {
	if head == nil {
		return 0
	}
	// 把 nums 的元素存储集合中
	seen := make(map[int]bool, len(nums))
	for _, v := range nums {
		seen[v] = true
	}
	// 遍历链表，统计链表组件个数。
	ans := 0
	for head != nil {
		// 如果当前节点在集合中 && 没有后续节点，
		// 则以当前节点为结尾，确定一个独立组件。
		// 同时达到链表末尾，退出循环。
		if seen[head.Val] && head.Next == nil {
			ans++
			break
		}
		// 如果当前节点在集合中 && 其下一个节点不在集合中，
		// 则以当前节点为结尾，确定一个独立组件。
		if seen[head.Val] && !seen[head.Next.Val] {
			ans++
		}
		// 如果以当前节点为结尾，尚未确定为一个组件，则指针继续往后走
		head = head.Next
	}
	return ans
}

// 817. Linked List Components
// 817. 链表组件
// 思路：（反向思维）一次遍历 + 哈希表
// time O(n) space O(n)
func numComponents2(head *pkg.ListNode, nums []int) int {
	if head == nil {
		return 0
	}
	// 把 nums 的元素存储集合中
	seen := make(map[int]bool, len(nums))
	for _, v := range nums {
		seen[v] = true
	}
	// 遍历链表，统计链表组件个数。
	// 链表中的组件数，最多为 nums 的元素个数
	ans := len(nums)
	for head.Next != nil {
		// 如果当前节点和下一个节点都在 nums 中，
		// 则两个节点链接形成一个新的组件，则组件个数 -1
		if seen[head.Val] && seen[head.Next.Val] {
			ans--
		}
		head = head.Next
	}
	return ans
}
