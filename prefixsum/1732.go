package prefixsum

// 1732. Find the Highest Altitude
// 1732. 找到最高海拔
// 策略：前缀和
// time O(N) space O(1)
func largestAltitude(gain []int) int {
	largest := 0
	preSum := make([]int, len(gain)+1)
	for i := 1; i <= len(gain); i++ {
		preSum[i] = preSum[i-1] + gain[i-1]
		largest = max(largest, preSum[i])
	}

	return largest
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
