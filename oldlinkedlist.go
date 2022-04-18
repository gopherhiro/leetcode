package main

import (
	"fmt"
)

func main() {
	elems := []int{3, 2, 2, 1, 1}
	head := createList(elems)
	printLinkedList(head)

	h1 := deleteDuplicatesIIRecursion(head)
	printLinkedList(h1)
}


type ListNode struct {
	Val  int
	Next *ListNode
}

func createLinkedList(len int) *ListNode {
	var prev, curr *ListNode
	for i := 1; i <= len; i++ {
		node := &ListNode{}
		node.Val = i
		node.Next = prev
		curr = node

		curr.Next = prev
		prev = curr
	}
	return curr
}

func createList(elms []int) *ListNode {
	var prev, curr *ListNode
	for _, value := range elms {
		node := &ListNode{}
		node.Val = value
		node.Next = prev
		curr = node

		curr.Next = prev
		prev = curr
	}
	return curr
}

func printLinkedList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func reverseList1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var headn, cur *ListNode
	for head != nil {
		cur = head
		head = head.Next
		cur.Next = headn
		headn = cur
	}

	return headn
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var prev *ListNode

	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

var breaker *ListNode

func reverseListN(head *ListNode, n int) *ListNode {
	if n == 1 {
		breaker = head.Next
		return head
	}

	last := reverseListN(head.Next, n-1)
	head.Next.Next = head
	head.Next = breaker

	return last
}

func reverseBetween(head *ListNode, m, n int) *ListNode {
	if m == 1 {
		return reverseListN(head, n)
	}
	head.Next = reverseBetween(head.Next, m-1, n-1)
	return head
}

func reverseMN(head *ListNode, left, right int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head

	pre := dummyHead
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}

	curr := pre.Next
	for i := 0; i < right-left; i++ {
		next := curr.Next
		curr.Next = next.Next

		next.Next = pre.Next
		pre.Next = next
	}

	return dummyHead.Next
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	a, b := head, head
	for i := 0; i < k; i++ {
		// 如果节点不足k个，不需要反转
		if b == nil {
			return head
		}
		b = b.Next
	}

	// 反转前k个元素
	newHead := reverseAB(a, b)

	// 递归反转后续链表（以b为头节点的链表），并链接起来
	a.Next = reverseKGroup(b, k)

	return newHead
}

// 反转以a为头节点的链表
func reverseANull(a *ListNode) *ListNode {
	var pre *ListNode
	cur := a
	// 遍历到链表尾
	for cur != nil {
		// 逐个反转
		nxt := cur.Next
		cur.Next = pre

		// 更新指针位置
		pre = cur
		cur = nxt
	}

	// 返回反转后的头节点
	return pre
}

// 反转区间节点[a,b)的链表元素，注意左闭右开。
func reverseAB(a, b *ListNode) *ListNode {
	var pre *ListNode
	cur := a
	// 遍历到指定的节点即可
	for cur != b {

		nxt := cur.Next
		cur.Next = pre

		pre = cur
		cur = nxt
	}

	return pre
}

func isPalindrome(head *ListNode) bool {
	// 先通过“双指针技巧”中的快、慢指针来找到链表的中点。
	// 找中点要注意区分奇偶情况。如果fast指针没有指向null，说明链表长度为奇数，slow指针还需要再前进一步。
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil {
		slow = slow.Next
	}

	// 从slow开始反转后面的链表。
	left := head
	right := reverseANull(slow)

	// 反转之后，即可进行回文比较。
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func middleNodeAgain(head *ListNode) *ListNode {
	var n int
	cur := head
	for cur != nil {
		n++
		cur = cur.Next
	}
	var k int
	cur = head
	for k < (n / 2) {
		k++
		cur = cur.Next
	}

	return cur
}

// 203. 移除链表元素
// 思路：遍历链表，查找删除即可。
// 创建虚拟头节点，方便删除。
// time O(N) space O(1)
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	virHead := &ListNode{}
	virHead.Next = head

	pre := virHead
	// 循环迭代条件：前驱节点的下一个节点不为空，则进入循环。
	// 为空，则结束循环。
	for pre.Next != nil {
		// 前驱节点的下一个节点值，等于指定值，则删除。
		if pre.Next.Val == val {
			pre.Next = pre.Next.Next
		} else {
			// 否则，移动到下一个元素继续判断。
			pre = pre.Next
		}
	}

	return virHead.Next
}

// 差值法
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var lenA, lenB int

	curA := headA
	for curA != nil {
		lenA++
		curA = curA.Next
	}

	curB := headB
	for curB != nil {
		lenB++
		curB = curB.Next
	}

	gap := abs(lenA, lenB)
	for i := 0; i < gap; i++ {
		if lenA > lenB {
			headA = headA.Next
		} else {
			headB = headB.Next
		}
	}

	for headA != nil && headB != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}

	return nil
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

// 哈希法
func intersectionNode(headA, headB *ListNode) *ListNode {
	hash := make(map[*ListNode]bool)
	for node := headA; node != nil; node = node.Next {
		hash[node] = true
	}
	for node := headB; node != nil; node = node.Next {
		if hash[node] {
			return node
		}
	}
	return nil
}

// 双指针法
func intersectionNodeN(headA, headB *ListNode) *ListNode {
	pa := headA
	pb := headB

	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}

		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}

	return pa
}

func kthToLast(head *ListNode, k int) int {
	slow := head
	fast := head

	for i := 0; i < k; i++ {
		fast = fast.Next
	}

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	return slow.Val
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

func deleteDuplicatesRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	head.Next = deleteDuplicatesRecursion(head.Next)
	if head.Val == head.Next.Val {
		return head.Next
	}
	return head
}

func deleteDuplicatesII(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{}
	dummy.Next = head

	pre, cur := dummy, head
	// 潜在点：pre.next == cur
	for cur != nil {
		// 遍历
		// 如果存在重复元素，使得cur指向当前重复元素的最后一个位置
		// 如果不存在重复元素，则后移
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
		}
		// 核心点：pre.next != cur 表示pre和cur之间存在重复元素，
		// 即cur移动了，使得pre.next不等于cur了。
		if pre.Next != cur {
			pre.Next = cur.Next // 删除pre和cur之间的重复元素
		} else {
			pre = pre.Next
		}
		cur = cur.Next
	}

	return dummy.Next
}

func deleteDuplicatesIIRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// if current node is not unique, return deleteDuplicates with head.next.
	if head.Next != nil && head.Val == head.Next.Val {
		for head.Next != nil && head.Val == head.Next.Val {
			head = head.Next
		}
		return deleteDuplicatesIIRecursion(head.Next)
	}
	// if current node is unique, link it to the result of next list made by recursive call.
	head.Next = deleteDuplicatesIIRecursion(head.Next)
	return head
}