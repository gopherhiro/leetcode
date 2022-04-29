package linkedlist

import (
	"container/heap"
	"leetcode/pkg"
)

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
