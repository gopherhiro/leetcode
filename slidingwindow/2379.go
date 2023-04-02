package slidingwindow

import "math"

// 2379. Minimum Recolors to Get K Consecutive Black Blocks
// 2379. 得到 K 个黑块的最少涂色次数
// solution: sliding window
// time O(n) space O(1)
func minimumRecolors(blocks string, k int) int {
	res, wCount := math.MaxInt, 0
	left, right := 0, 0
	for right < len(blocks) {
		if blocks[right] == 'W' {
			wCount++
		}
		// Get K Consecutive Black Blocks
		// the W's count is the number of recolors
		// the minimum of the W's count is the answer
		if right-left+1 == k {
			res = min(res, wCount)
			// shrink the window
			if blocks[left] == 'W' {
				wCount--
			}
			left++
		}
		right++
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
