package prefixsum

// 724. Find Pivot Index
// 724. 寻找数组的中心下标
// time O(N) space O(N)
func pivotIndexI(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	total := 0
	for _, num := range nums {
		total += num
	}

	// leftSum + num[i] + rightSum == total
	// if leftSum == rightSum
	// so that 2 * sum + num[i] == total
	// so that i is pivot index
	sum := 0 // sum of left
	for i := 0; i < len(nums); i++ {
		if 2*sum+nums[i] == total {
			return i
		}
		sum += nums[i]
	}
	return -1
}

// 724. Find Pivot Index
// 724. 寻找数组的中心下标
// 思路：数组中心下标：左侧所有元素相加的和等于右侧所有元素相加的和。
// 遍历数组，利用前缀和求出下标左右侧元素之和，比较左右两侧元素和是否相等。
// 若相等，则找到中心下标。
// time O(N) space O(N)
// 注意：
// preSum[i] 对应 num[0..i-1]的和。
// num[i..j] 的和为 preSum[j+1] - preSum[i]
func pivotIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	// 构建前缀和
	preSum := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}

	// 遍历数组，找寻中心下标记
	n := len(nums)
	for i := 0; i < n; i++ {
		leftSum := preSum[i] - preSum[0]    // nums[0..i-1] 的和
		rightSum := preSum[n] - preSum[i+1] // nums[i+1..n-1] 的和
		if leftSum == rightSum {
			return i
		}
	}
	return -1
}
