package twopointer

// 42. 接雨水
// 42. Trapping Rain Water
// 思路：Dynamic Programming
// min(left_max[i],right_max[i])−height[i]
// time O(N) space O(N)
func trapRain(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}

	leftMax, rightMax := make([]int, n), make([]int, n)

	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(height[i], leftMax[i-1])
	}

	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(height[i], rightMax[i+1])
	}

	answer := 0
	for i := 1; i < n-1; i++ {
		answer += min(leftMax[i], rightMax[i]) - height[i]
	}
	return answer
}

// 42. 接雨水
// 42. Trapping Rain Water
// 思路：双指针
// time O(N) space O(1)
func trapRainWater(height []int) int {
	n := len(height)
	left, right := 0, n-1
	answer, leftMax, rightMax := 0, 0, 0
	for left < right {
		// leftMax 是 height[0..left] 中最高柱子的高度
		// rightMax 是 height[right..end] 的最高柱子的高度。
		leftMax = max(height[left], leftMax)
		rightMax = max(height[right], rightMax)
		if leftMax < rightMax {
			// 左边是短板（右边给左边保证）
			answer += leftMax - height[left]
			left++
		} else {
			// 右边是短板（左边给右边保证）
			answer += rightMax - height[right]
			right--
		}
	}
	return answer
}
