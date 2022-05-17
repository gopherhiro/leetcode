package prefixsum

// 1413. Minimum Value to Get Positive Step by Step Sum
// 1413. 逐步求和得到正数的最小值
// time O(N) space O(1)
// Algorithm:
// 1.Traverse the array nums and calculate every step-by-step total,
// use total to record the current step-by-step total,
// and minVal to record the minimum step-by-step total.
// 2.Return -minVal + 1, that is the minimum valid startValue.
func minStartValue(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	minVal := 0
	total := 0
	for _, num := range nums {
		total += num
		minVal = min(minVal, total)
	}

	return 1 - minVal
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
