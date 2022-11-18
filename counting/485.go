package counting

// 485. Max Consecutive Ones
// 485. 最大连续 1 的个数
// 思路：一次遍历
// time O(n) space O(1)
func findMaxConsecutiveOnes(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	ans, count := 0, 0
	for _, num := range nums {
		if num == 1 {
			count += 1
		} else {
			count = 0
		}
		ans = max(ans, count)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
