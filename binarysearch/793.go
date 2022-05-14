package binarysearch

import "math"

// 793. 阶乘函数后 K 个零
// 思路：搜索有多少个 n 满足 trailingZeroes(n) == K，
// 其实就是在问，满足条件的 n 最小是多少，最大是多少，最大值和最小值一减，就可以算出来有多少个 n 满足条件了。
// 那就是 二分查找 中「搜索左侧边界」和「搜索右侧边界」呀
// 观察题目给出的数据取值范围，n 可以在区间 [0, LONG_MAX] 中取值，
// 寻找满足 trailingZeroes(n) == K 的左侧边界和右侧边界， 左侧边界和右侧边界之差 + 1即是答案。
func preimageSizeFZF(k int) int {
	min := leftBorder(k)
	max := rightBorder(k)
	return max - min + 1
}

func leftBorder(target int) int {
	left, right := 0, math.MaxInt64
	for left <= right {
		mid := left + (right-left)/2
		if trailingZeroes(mid) < target {
			left = mid + 1
		} else if trailingZeroes(mid) > target {
			right = mid - 1
		} else if trailingZeroes(mid) == target {
			right = mid - 1
		}
	}
	return left
}

func rightBorder(target int) int {
	left, right := 0, math.MaxInt64
	for left <= right {
		mid := left + (right-left)/2
		if trailingZeroes(mid) < target {
			left = mid + 1
		} else if trailingZeroes(mid) > target {
			right = mid - 1
		} else if trailingZeroes(mid) == target {
			left = mid + 1
		}
	}
	return right
}
