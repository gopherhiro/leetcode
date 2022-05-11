package dp

// 651. 4 Keys Keyboard
// 651. 4键键盘
// 思路：dynamic programming
// dp数组定义：dp[i] 为第 i 次操作后，能显示A的个数。
// time O(n^2) space O(n)
func maxA(n int) int {
	if n <= 1 {
		return n
	}
	dp := make([]int, n+1)

	// base case
	dp[0] = 0
	for i := 1; i <= n; i++ {
		// 按 A 键
		dp[i] = dp[i-1] + 1

		// 剪切板中的 A 的个数，取决于上次按 CA CC的时机。
		// 即当前位置 j 准备 CV，则剪切板中的 A 的个数为：dp[j-2]。
		// 按 CV 键
		for j := 2; j < i; j++ {
			// 全选(CA)并复制(CC) dp[j-2]
			// 连续粘贴 i - j次，再加上 dp[j-2] 本身
			pressCV := dp[j-2] * (i - j + 1)
			// 选取最大值作为 dp[i] 的值
			dp[i] = max(dp[i], pressCV)
		}
	}

	return dp[n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 651. 4 Keys Keyboard
// 651. 4键键盘
// 思路：recursion + memo
// dp函数定义：状态 (n, number, copy) 能显示A的个数为：dp(n, number, copy)
// number:屏幕上显示 A 的个数
// copy:剪切板 A 的个数。
// time O(n^3) space O(n^3)
// 三个状态的组合，复杂度确实很高。
func maxAM(n int) int {
	if n <= 1 {
		return n
	}

	memo := make(map[string]int, 0)

	var dp func(n, number, copy int) int
	dp = func(n, number, copy int) int {
		// base case
		if n <= 0 {
			return number
		}

		key := fmt.Sprintf("%d,%d,%d", n, number, copy)
		if val, ok := memo[key]; ok {
			return val
		}

		// 每一次操作，有 4 种选择
		// 按下 A 键，屏幕上增加一个字符，同时消耗一个操作数。
		A := dp(n-1, number+1, copy)

		// 按下 CV 粘贴，剪切板中的字符加入屏幕，同时消耗一个操作数。
		CV := dp(n-1, number+copy, copy)

		// 全选和复制（CA、CC）键联合使用
		// 屏幕上的字符数量变为剪切板上的字符数，同时消耗两个操作数。
		CACC := dp(n-2, number, number)

		// 选择最大的返回
		memo[key] = max(A, max(CV, CACC))
		return memo[key]
	}
	// 可以按下 N 次，屏幕和剪切板都还没有字符A。
	return dp(n, 0, 0)
}

// 651. 4 Keys Keyboard
// 651. 4键键盘
// 思路：recursion
// dp函数定义：状态 (n, number, copy) 能显示A的个数为：dp(n, number, copy)
// number:屏幕上显示 A 的个数
// copy:剪切板 A 的个数。
// time O(n^3) space O(n^3)
// 三个状态的组合，复杂度确实很高。
func maxAR(n int) int {
	if n <= 1 {
		return n
	}

	var dp func(n, number, copy int) int
	dp = func(n, number, copy int) int {
		// base case
		if n <= 0 {
			return number
		}
		// 每一次操作，有 4 种选择
		// 选择其中的最大即是我们想要的answer。
		// 按下 A 键，屏幕上增加一个字符，同时消耗一个操作数。
		A := dp(n-1, number+1, copy)

		// 按下 CV 粘贴，剪切板中的字符加入屏幕，同时消耗一个操作数。
		CV := dp(n-1, number+copy, copy)

		// 全选和复制（CA、CC）键联合使用
		// 屏幕上的字符数量变为剪切板上的字符数，同时消耗两个操作数。
		CACC := dp(n-2, number, number)

		// 选择最大的返回
		return max(A, max(CV, CACC))
	}
	return dp(n, 0, 0)
}
