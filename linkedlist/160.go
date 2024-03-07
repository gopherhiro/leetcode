package linkedlist

import "leetcode/pkg"

// 推荐前两种解法，容易理解。

// 160. Intersection of Two Linked Lists
// 160. 相交链表
// 策略：哈希表
// time O(M+N) space O(N)
func getIntersectionNodeH(headA, headB *pkg.ListNode) *pkg.ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	visit := make(map[*pkg.ListNode]bool)
	pa, pb := headA, headB
	// 使用哈希表记录下链表A中的所有节点
	for pa != nil {
		visit[pa] = true
		pa = pa.Next
	}
	// 遍历链表B，搜索相交的节点
	// 如果搜索到，则返回节点
	// 如果搜索不到，则返回空
	for pb != nil {
		if visit[pb] {
			return pb
		}
		pb = pb.Next
	}
	return nil
}

// 160. Intersection of Two Linked Lists
// 160. 相交链表
// 策略：链表长度差值法
// time O(M+N) space O(1)
func getIntersectionNode(headA, headB *pkg.ListNode) *pkg.ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	lenA, lenB := 0, 0
	pa, pb := headA, headB
	for pa != nil {
		lenA++
		pa = pa.Next
	}
	for pb != nil {
		lenB++
		pb = pb.Next
	}

	// 其中一个先走 k 步，让 p1,p2 到达尾部的距离相同
	k := pkg.Abs(lenA, lenB)
	p1, p2 := headA, headB
	if lenA > lenB {
		for i := 0; i < k; i++ {
			p1 = p1.Next
		}
	} else {
		for i := 0; i < k; i++ {
			p2 = p2.Next
		}
	}
	// p1 == p2
	// 1、链表不相交，同时走到尾部，都空指针。
	// 2、链表相交，同时走到相交节点。
	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p1
}

// 160. Intersection of Two Linked Lists
// 160. 相交链表
// 策略：双指针法
// time O(M+N) space O(1)
// 两个指针分别遍历两个链表，当其中一个遍历到链表末尾时，将其重新定位到另一个链表的头部，这样两个指针总共遍历了两个链表的长度之和。
// 由于两个指针同时移动，它们的相对速度相同，因此在某个时刻，它们要么在交点相遇，要么同时到达链表末尾。
// 核心思想是在两个链表上同时进行遍历，通过双指针的方式，避免了需要知道链表长度的情况，从而找到链表的交点。
func getIntersectionNodeT(headA, headB *pkg.ListNode) *pkg.ListNode {
	pa, pb := headA, headB
	for pa != pb {
		// 如果某个指针到达链表末尾，将其重新定位到另一个链表的头节点
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}

		if pb == nil {
			// 返回相遇的节点（可能是交点或者两个链表末尾的 nil）
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}
