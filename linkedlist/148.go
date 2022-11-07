package linkedlist

// 148. Sort List
// 148. 排序链表
// 思路：递归-归并排序
// time O(nlogn) space O(logn)
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// step 1. cut the list to two halves
	// step 2. sort each half
	// step 3. merge left and right
	mid := midCut(head)
	left := sortList(head)
	right := sortList(mid)
	return mergeList(left, right)
}

// get the middle node of the list
// and cut the list to two halves
func midCut(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	prev, slow, fast := head, head, head
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = nil
	return slow
}

// 21. Merge Two Sorted Lists
// 21. 合并两个有序链表
// 策略：迭代
// 思路：双指针
// time O(M+N) space O(1)
func mergeList(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	dummy := &ListNode{}
	p, p1, p2 := dummy, l1, l2
	for p1 != nil && p2 != nil {
		// p 链接到值较小的节点
		if p1.Val < p2.Val {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
		}
		// 继续后续链接
		p = p.Next
	}
	if p1 != nil {
		p.Next = p1
	}
	if p2 != nil {
		p.Next = p2
	}
	return dummy.Next
}
