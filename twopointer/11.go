package twopointer

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
