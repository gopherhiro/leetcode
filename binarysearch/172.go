package binarysearch

// 172. 阶乘后的零
// 思路：数学原理，计算乘法因子中5的个数。
func trailingZeroes(n int) int {
	ans := 0
	for n > 0 {
		ans += n / 5
		n = n / 5
	}
	return ans
}
