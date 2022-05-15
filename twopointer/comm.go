package twopointer

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 验证回文性
func checkPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// 创建链表
// 尾插法
// time O(N) space O(1)
func createListFromTail(nums []int) *ListNode {
	var head, tail, curr *ListNode
	curr = &ListNode{}
	head = curr
	for _, value := range nums {
		// 新建节点，并指向尾节点
		node := &ListNode{Val: value}
		node.Next = tail
		curr.Next = node
		curr = node
	}
	return head.Next
}

// 创建链表
// 头插法
// time O(N), space O(1)
func createListFromHead(nums []int) *ListNode {
	var prev, curr *ListNode
	for _, value := range nums {
		node := &ListNode{}
		node.Val = value
		node.Next = prev
		curr = node
		prev = curr
	}
	return curr
}

func printLinkedList(head *ListNode) {
	for head != nil {
		if head.Next != nil {
			fmt.Printf("%d->", head.Val)
		} else {
			fmt.Printf("%d", head.Val)
		}
		head = head.Next
	}
	fmt.Println()
}
