package dp

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
		sum := prev + curr
		prev = curr
		curr = sum
	}
	return curr
}

// 509. Fibonacci Number
// 509. 斐波那契数
// 思路：动态迭代（自下而上）
// time O(N) space O(N)
func fibDP(n int) int {
	if n <= 1 {
		return n
	}
	// 自下而上，逐步迭代
	dp := make(map[int]int, n+1)
	dp[0], dp[1] = 0, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 509. Fibonacci Number
// 509. 斐波那契数
// 思路：带备忘录的递归（自顶而下）
// time O(N) space O(N)
func fibM(n int) int {
	memo := make(map[int]int, n+1)

	var helper func(n int) int
	helper = func(n int) int {
		if n <= 1 {
			return n
		}
		// 先查询备忘录，存在直接返回。
		if _, ok := memo[n]; ok {
			return memo[n]
		}
		// 计算并存储到备忘录中，并返回值
		memo[n] = helper(n-1) + helper(n-2)
		return memo[n]
	}

	return helper(n)
}

// 509. Fibonacci Number
// 509. 斐波那契数
// 思路：暴力递归
// time O(2^N) space O(N)
func fibR(n int) int {
	if n <= 1 {
		return n
	}
	return fibR(n-1) + fibR(n-2)
}
