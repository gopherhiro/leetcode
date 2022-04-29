package pkg

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// CreateListFromTail 创建链表
// 尾插法
// time O(N) space O(1)
func CreateListFromTail(nums []int) *ListNode {
	// 虚拟头节点(构建新链表时的常用技巧)
	dummyHead := &ListNode{}

	var tail *ListNode
	tail = dummyHead

	for _, value := range nums {
		new := &ListNode{Val: value}
		tail.Next = new
		tail = tail.Next
	}
	return dummyHead.Next
}

// CreateListFromTailN 创建链表
// 尾插法
// time O(N) space O(1)
func CreateListFromTailN(nums []int) *ListNode {
	// head 标记头节点
	// tail 标记尾节点
	var head, tail *ListNode
	for _, value := range nums {
		new := &ListNode{Val: value}
		if head == nil {
			head = new // 标记头节点
			tail = new // 标记尾节点
		} else {
			tail.Next = new
			tail = tail.Next
		}
	}
	return head
}

func PrintLinkedList(head *ListNode) {
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
