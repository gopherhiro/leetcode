package main

import (
	"container/heap"
	"leetcode/pkg"
)

func main() {
}

// 23. Merge k Sorted Lists
// 23. 合并K个升序链表
// 策略：合并两个升序链表
// 思路：连续合并链表
// time O(KN)（N 是链表总节点数，K 是链表条数）
func mergeKListsA(lists []*pkg.ListNode) *pkg.ListNode {
	k := len(lists)
	if k == 0 {
		return nil
	}
	ans := lists[0]
	for i := 1; i < k; i++ {
		ans = mergeTwoLists(ans, lists[i])
	}
	return ans
}

// 23. Merge k Sorted Lists
// 23. 合并K个升序链表
// 策略：最小堆
// 思路：把链表节点放入一个最小堆，就可以每次获得 k 个节点中的最小节点。
// time O(N*logK)（N 是链表节点数，K 是链表条数）
func mergeKLists(lists []*pkg.ListNode) *pkg.ListNode {
	k := len(lists)
	if k == 0 {
		return nil
	}
	// 把 k 个链表头节点，放入最小堆
	h := new(minHeap)
	for _, head := range lists {
		if head != nil {
			heap.Push(h, head)
		}
	}

	dummy := new(pkg.ListNode)
	p := dummy
	for h.Len() > 0 {
		// 获取最小堆中的顶点，接到结果链中
		// 将当前节点的下一个节点，放入最小堆
		node := heap.Pop(h).(*pkg.ListNode)
		p.Next = node
		if node.Next != nil {
			heap.Push(h, node.Next)
		}
		p = p.Next
	}
	return dummy.Next
}

// An minHeap is a min-heap of *ListNode
type minHeap []*pkg.ListNode

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*pkg.ListNode))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 21. Merge Two Sorted Lists
// 21. 合并两个有序链表
// 策略：迭代
// 思路：双指针
// time O(M+N) space O(1)
func mergeTwoLists(l1 *pkg.ListNode, l2 *pkg.ListNode) *pkg.ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	dummy := &pkg.ListNode{} // 虚拟头节点
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

// 2. Add Two Numbers
// 2. 两数相加
// 思路：模拟加法
// Complexity: time O(max(m,n)) space O(max(m,n))
func addTwoNumbers(l1 *pkg.ListNode, l2 *pkg.ListNode) *pkg.ListNode {
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
	dummyHead := &pkg.ListNode{}

	var tail *pkg.ListNode
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
		new := &pkg.ListNode{Val: curr}
		tail.Next = new
		tail = tail.Next
	}
	// 如果最后有进位，则需要生成一个新的节点
	if carry > 0 {
		tail.Next = &pkg.ListNode{Val: carry}
	}

	return dummyHead.Next
}
