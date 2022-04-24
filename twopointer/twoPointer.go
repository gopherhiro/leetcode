package main

import "fmt"

func main() {
	h := []int{4, 2, 0, 3, 2, 5}
	r := trapRain(h)
	fmt.Println(r)
}

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

// 11. Container With Most Water
// 11. 盛最多水的容器
// 思路：双指针，min(height[left], height[right]) * (right - left)
// time O(N) space O(1)
func maxArea(height []int) int {
	n := len(height)
	left, right := 0, n-1
	ans := 0
	for left < right {
		// 寻找短板
		shortBoard := min(height[left], height[right])
		// 计算面积：储存水量
		S := shortBoard * (right - left)
		// 择优取最大值
		ans = max(ans, S)
		// 移动较低的一边，即朝着矩形面积可能变大的方向走。
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return ans
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
