package slidingwindow

// 220. Contains Duplicate III
// 220. 存在重复元素 III
// 思路：暴力枚举所有可能
// time O(N^2) space O(1)
func containsNearbyDup(nums []int, k, t int) bool {
	if len(nums) < 2 {
		return false
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if abs(nums[i], nums[j]) <= t && abs(i, j) <= k {
				return true
			}
		}
	}
	return false
}
