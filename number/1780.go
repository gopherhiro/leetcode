package number

// 1780. Check if Number is a Sum of Powers of Three
// 1780. 判断一个数字是否可以表示成三的幂的和
// 思路：看看这个数是否可以转换成「三进制」表示
// 三进制，即三的幂的和
// 与数的「二进制」表示一样，数的「三进制」表示，也只有 0 or 1
// time O(logn) space O(1)
func checkPowersOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	for n > 0 {
		remain := n % 3
		// 出现非 0 or 1 的余数，则不能表示为「三进制」
		if remain == 2 {
			return false
		}
		n = n / 3
	}
	return true
}
