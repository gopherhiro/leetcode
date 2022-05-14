package sort

// 912. Sort an Array
// 912. 排序数组
// 策略：堆排序
// time O(N*logN) space O(1)
func heapSort(nums []int) []int {
	end := len(nums) - 1
	initHeap(nums, end)
	for i := end; i >= 0; i-- {
		swap2(nums, 0, i)
		end--
		heapify(nums, 0, end)
	}
	return nums
}

func initHeap(nums []int, end int) {
	for i := end / 2; i >= 0; i-- {
		heapify(nums, i, end)
	}
}

// 自上往下堆化
func heapify(nums []int, i, end int) {
	for {
		left, right := i*2+1, i*2+2
		if left > end {
			break
		}
		maxPos := left
		if left < end && nums[left] <= nums[right] {
			maxPos = right
		}
		// 停止堆化
		if nums[i] > nums[maxPos] {
			break
		}
		swap2(nums, i, maxPos)
		i = maxPos
	}
}

func swap2(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
