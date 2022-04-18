package main

import "fmt"

func main() {
	nums := []int{9, 9, 9, 9}
	head1 := createListFromTail(nums)
	printLinkedList(head1)
	nums2 := []int{9, 9, 9, 9, 9, 9}
	head2 := createListFromTail(nums2)
	printLinkedList(head2)

	head := addTwoNumbers(head1, head2)
	printLinkedList(head)
}

type ListNode struct {
	Val  int
	Next *ListNode
}




// 2. Add Two Numbers
// 2. 两数相加
// 思路：模拟加法
// Complexity: time O(max(m,n)) space O(max(m,n))
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	// 虚拟头节点
	dummyHead := &ListNode{}

	var tail *ListNode
	tail = dummyHead

	carry := 0 // 进位
	for l1 != nil || l2 != nil {
		// 如果两个链表的长度不同，默认短链的后面都是0
		// 这样的话，即在理论上认为两个链表总是长度相同
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		sum := v1 + v2 + carry // 求和
		carry = sum / 10       // 进位
		curr := sum % 10       // 当前值

		// 尾插法建立链表
		new := &ListNode{Val: curr}
		tail.Next = new
		tail = tail.Next
	}
	// 如果最后有进位，则需要生成一个新的节点
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}

	return dummyHead.Next
}

// 创建链表
// 尾插法
// time O(N) space O(1)
func createListFromTail(nums []int) *ListNode {
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

// 创建链表
// 尾插法
// time O(N) space O(1)
func createListFromTailN(nums []int) *ListNode {
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
