package offer

// 509. Fibonacci Number
// 509. 斐波那契数
// 思路：迭代（自下而上，动态压缩）
// time O(N) space O(1)
func fib(n int) int {
	if n <= 1 {
		return n
	}
	// f(i) 只与 f(i-1)、f(i-2)两种状态相关
	// 故而进行状态压缩，迭代两个状态即可。
	// 从而大大降低空间复杂度。
	prev, curr := 0, 1
	for i := 2; i <= n; i++ {
		sum := (prev + curr) % 1000000007
		prev = curr
		curr = sum
	}
	return curr
}

func fiba(n int) int {
	if n <= 1 {
		return n
	}

	dp := make(map[int]int, n+1)

	dp[0], dp[1] = 0, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
		dp[i] %= 1000000007
	}
	return dp[n]
}
