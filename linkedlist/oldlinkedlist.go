package linkedlist

import "leetcode/pkg"

var breaker *pkg.ListNode

func reverseListN(head *pkg.ListNode, n int) *pkg.ListNode {
	if n == 1 {
		breaker = head.Next
		return head
	}

	last := reverseListN(head.Next, n-1)
	head.Next.Next = head
	head.Next = breaker

	return last
}

func reverseBetween(head *pkg.ListNode, m, n int) *pkg.ListNode {
	if m == 1 {
		return reverseListN(head, n)
	}
	head.Next = reverseBetween(head.Next, m-1, n-1)
	return head
}

func reverseMN(head *pkg.ListNode, left, right int) *pkg.ListNode {
	dummyHead := &pkg.ListNode{}
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

func reverseKGroup(head *pkg.ListNode, k int) *pkg.ListNode {
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
func reverseANull(a *pkg.ListNode) *pkg.ListNode {
	var pre *pkg.ListNode
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
func reverseAB(a, b *pkg.ListNode) *pkg.ListNode {
	var pre *pkg.ListNode
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

func isPalindrome(head *pkg.ListNode) bool {
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

func middleNode(head *pkg.ListNode) *pkg.ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func middleNodeAgain(head *pkg.ListNode) *pkg.ListNode {
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

func kthToLast(head *pkg.ListNode, k int) int {
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

func hasCycle(head *pkg.ListNode) bool {
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

func deleteDuplicates(head *pkg.ListNode) *pkg.ListNode {
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

func deleteDuplicatesRecursion(head *pkg.ListNode) *pkg.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	head.Next = deleteDuplicatesRecursion(head.Next)
	if head.Val == head.Next.Val {
		return head.Next
	}
	return head
}

func deleteDuplicatesII(head *pkg.ListNode) *pkg.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &pkg.ListNode{}
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

func deleteDuplicatesIIRecursion(head *pkg.ListNode) *pkg.ListNode {
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
