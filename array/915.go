package array

// 915. Partition Array into Disjoint Intervals
// 915. 分割数组
// 思路：与前缀和相似
// time O(n) space O(n)
func partitionDisjoint(nums []int) int {
	n := len(nums)
	minRight := make([]int, n)
	minRight[n-1] = nums[n-1]
	// to find min right of i
	for i := n - 2; i >= 0; i-- {
		minRight[i] = min(nums[i], minRight[i+1])
	}

	// to find index of max(left) <= min(right)
	// becaues from left to right query, so the first i of found is smallest
	maxLeft := nums[0]
	for i := 1; i < n; i++ {
		// maxLeft[i-1] is the maximum of subarray nums[:i].
		// minRight[i] is the minimum of subarray nums[i:].
		maxLeft = max(maxLeft, nums[i-1])
		if maxLeft <= minRight[i] {
			return i
		}
	}
	return -1
}

// 915. Partition Array into Disjoint Intervals
// 915. 分割数组
// 思路：与前缀和相似
// time O(n) space O(n)
func partitionDisjoint2(nums []int) int {
	n := len(nums)
	maxLeft, minRight := make([]int, n), make([]int, n)
	maxLeft[0], minRight[n-1] = nums[0], nums[n-1]

	// to find max left of i
	for i := 1; i < n; i++ {
		maxLeft[i] = max(maxLeft[i-1], nums[i])
	}
	// to find min right of i
	for i := n - 2; i >= 0; i-- {
		minRight[i] = min(nums[i], minRight[i+1])
	}

	// to find i of max(left) <= min(right)
	for i := 1; i < n; i++ {
		// maxLeft[i-1] is the maximum of subarray nums[:i].
		// minRight[i] is the minimum of subarray nums[i:].
		if maxLeft[i-1] <= minRight[i] {
			return i
		}
	}
	return -1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
