package sort

// 912. Sort an Array
// 912. 排序数组
// 策略：归并排序
// time O(N*logN) space O(N)
func mergeSort(nums []int) []int {
	n := len(nums)
	if n == 0 || n == 1 {
		return nums
	}
	var sort func(num []int, lo, hi int)
	sort = func(nums []int, lo, hi int) {
		// base case
		if lo >= hi {
			return
		}
		mid := lo + (hi-lo)>>1

		// 左边已排序
		sort(nums, lo, mid)

		// 右边已排序
		sort(nums, mid+1, hi)

		// 合并左边和右边
		merge(nums, lo, mid, hi)
	}
	sort(nums, 0, n-1)
	return nums
}

func merge(nums []int, lo, mid, hi int) {
	merged := make([]int, 0)
	i, j := lo, mid+1
	for i <= mid && j <= hi {
		if nums[i] < nums[j] {
			merged = append(merged, nums[i])
			i++
		} else {
			merged = append(merged, nums[j])
			j++
		}
	}
	// 如果左边部分剩余，则合并剩余左边部分
	if i <= mid {
		merged = append(merged, nums[i:mid+1]...)
	}
	// 如果右边部分剩余，合并剩余右边部分
	if j <= hi {
		merged = append(merged, nums[j:hi+1]...)
	}
	// 将 merged 中的数组拷贝回 nums
	copy(nums[lo:hi+1], merged)
}
