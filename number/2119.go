package number

// 2119. A Number After a Double Reversal
// 2119. 反转两次的数字
// 思路：反转两次，对比即可。
// time O(N) space O(1)
func isSameAfterReversalsN(num int) bool {
	if num == 0 {
		return true
	}
	r1 := reverseN(num)
	r2 := reverseN(r1)
	return r2 == num
}
