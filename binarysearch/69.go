package binarysearch

// 69. x 的平方根
// 69. Sqrt(x)
// only the integer part of the result is returned.
// 思路：在 0 ~ x 范围内进行二分搜索即可。
func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	left, right := 2, x
	for left <= right {
		mid := left + (right-left)>>1
		if mid == x/mid {
			return mid
		} else if mid < x/mid {
			left = mid + 1
		} else if mid > x/mid {
			right = mid - 1
		}
	}
	return right
}
