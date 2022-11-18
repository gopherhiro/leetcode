package slidingwindow

// 1004. Max Consecutive Ones III
// 1004. 最大连续1的个数 III
// 思路：sliding window
// time O(n) space O(1)
func longestOnes(nums []int, k int) int {
	left, right := 0, 0
	for right < len(nums) {
		if nums[right] == 0 {
			k--
		}
		if k < 0 {
			if nums[left] == 0 {
				k++
			}
			left++
		}
		right++
		// right forward alone, or left and right forward together.
		// so the window only going to get bigger
	}
	return right - left
}

// 1004. Max Consecutive Ones III
// 1004. 最大连续1的个数 III
// 思路：sliding window
// time O(n) space O(1)
func longestOnesW(nums []int, k int) int {
	res := 0
	left, right := 0, 0
	for right < len(nums) {
		if nums[right] == 0 {
			k-- // flip 0 to 1
		}
		for k < 0 { // sub array only contain k 0's
			if nums[left] == 0 {
				k++
			}
			left++
		}
		// get the result
		res = max(res, right-left+1)
		right++
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
