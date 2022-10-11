package number

// 231. Power of Two
// 231. 2 的幂
// 思路：bitwise operations
// 如果一个数是 2 的指数，那么它的二进制表示一定只含有一个 1。
// time O(1) space O(1)
func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	return n&(n-1) == 0
}
