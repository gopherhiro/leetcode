package topk

// 703. Kth Largest Element in a Stream
// 703. 数据流中的第 K 大元素
// 提示：使用快速选择，超时，推荐使用堆排。
type KthLargest struct {
	nums []int
	k    int
}

func Constructor(k int, nums []int) KthLargest {
	return KthLargest{
		nums: nums,
		k:    k,
	}
}

func (lg *KthLargest) Add(val int) int {
	lg.nums = append(lg.nums, val)
	kth := findKthLargest(lg.nums, lg.k)
	return kth
}
