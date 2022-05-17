package slidingwindow

// 219. Contains Duplicate II
// 219. 存在重复元素 II
// 思路：sliding window
// time O(N) space O(N)
func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) < 2 {
		return false
	}
	window := make(map[int]int, 0)
	left, right := 0, 0
	for right < len(nums) {
		c := nums[right]
		window[c]++
		right++

		// 存在重复值
		for window[c] > 1 {
			d := nums[left]
			left++
			window[d]--
			// 让 left 一直移至没有重复元素为止。
			if window[c] <= 1 {
				// 此时的坐标才是正确的
				// 计算坐标绝对值 && 与 k 比较
				if abs(left-1, right-1) <= k {
					return true
				}
			}
		}
	}

	return false
}

// 219. Contains Duplicate II
// 219. 存在重复元素 II
// 思路：hashtable
// time O(N) space O(N)
func containsNearbyDuplicateH(nums []int, k int) bool {
	if len(nums) < 2 {
		return false
	}
	hash := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		elem := nums[i]
		if j, ok := hash[elem]; ok {
			if abs(i, j) <= k {
				return true
			}
		}
		hash[elem] = i
	}
	return false
}

func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}
