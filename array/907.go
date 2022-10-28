package array

// 907. Sum of Subarray Minimums
// 907. 子数组的最小值之和
// 思路：暴力破解
// time O(n^2) space O(1)
// 备注：超时
func sumSubarrayMins(arr []int) int {
	n := len(arr)
	sumMin := 0
	for start := 0; start < n; start++ {
		curMin := arr[start]
		for end := start; end < n; end++ {
			curMin = min(curMin, arr[end])
			sumMin += curMin
		}
	}
	return sumMin
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
