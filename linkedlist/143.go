package linkedlist

// 143. 重排链表
// 思路：模拟, 找寻中点、断开、反转后半段、合并
// time O(N) space O(1)
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}
	/*	count := listCount(head)
		mid := 0
		if count%2 == 0 {
			mid = count / 2
		} else {
			mid = count/2 + 1
		}

		p1 := head
		for i := 0; i < mid-1; i++ {
			p1 = p1.Next
		}

		p2 := p1.Next

		p1.Next = nil*/
	last := midCut(head)

	relast := reverse(last)

	head = merge(head, relast)
}

func midCut(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
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

func merge(p1, p2 *ListNode) *ListNode {
	if p1 == nil {
		return p2
	}
	if p2 == nil {
		return p1
	}

	dummy := &ListNode{}
	p := dummy
	for p1 != nil && p2 != nil {
		p.Next = p1
		p1 = p1.Next
		p = p.Next

		p.Next = p2
		p2 = p2.Next
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

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev, curr *ListNode
	for head != nil {
		curr = head
		head = head.Next
		curr.Next = prev
		prev = curr
	}
	return curr
}

func listCount(head *ListNode) int {
	if head == nil {
		return 0
	}
	count := 0
	for head != nil {
		count++
		head = head.Next
	}
	return count
}
