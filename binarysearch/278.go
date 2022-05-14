package binarysearch

// 278. 第一个错误的版本
// 思路：二分左边界搜索
// 注意性质：当一个版本为正确版本，则该版本之前的所有版本均为正确版本；
// 当一个版本为错误版本，则该版本之后的所有版本均为错误版本。
func firstBadVersion(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	left, right := 1, n
	for left <= right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func isBadVersion(version int) bool {
	return true
}
