package number

// 338. Counting Bits
// 338. 比特位计数
// 思路：二进制位运算
// time O(nlogn) space O(n)
func countBits(n int) []int {
	ans := make([]int, 0, n+1)
	for i := 0; i <= n; i++ {
		res := countBit(i)
		ans = append(ans, res)
	}
	return ans
}

// time O(logn) space O(1)
func countBit(num int) int {
	count := 0
	for num > 0 {
		num = num & (num - 1)
		count++
	}
	return count
}

// 338. Counting Bits
// 338. 比特位计数
// 思路：二进制位规律统计
// time O(n) space O(n)
func countBits2(n int) []int {
	ans := make([]int, n+1)
	for i := 1; i <= n; i++ {
		// 二进制表示中
		// 奇数中 1 的个数，比前面「那个偶数」多一个 1
		// 偶数中 1 的个数，和「除以 2 之后的那个数」一样多
		if i%2 == 1 {
			ans[i] = ans[i-1] + 1
		} else {
			ans[i] = ans[i/2]
		}
	}
	return ans
}
