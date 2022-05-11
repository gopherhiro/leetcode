package dp

import "fmt"

// 10. Regular Expression Matching
// 10. 正则表达式匹配
// 思路：dynamic programming
// dp 数组定义：
// 1、p[0..j] 可以匹配 s[0..i]，则 dp[i,j] = true
// 2、p[0..j] 不能匹配 s[0..i]，则 dp[i,j] = false
// time O(MN) , space O(MN)
func isMatch(s string, p string) bool {
	ms, np := len(s), len(p)
	if ms == 0 && np == 0 {
		return true
	}

	dp := generateBoolDp(ms, np)

	// base case
	// 两者都为空，肯定是匹配成功的。
	dp[0][0] = true

	// p 为空串，s 不为空，则不匹配。
	/*	for i := 1; i < ms; i++ {
		dp[i][0] = false
	}*/

	// s 为空串，p不为空。p 有可以匹配空串的情况。
	// 则处理 a* or a*b* or a*b*c* 可以匹配空串情况
	for j := 1; j < np; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	for i := 0; i < ms; i++ {
		for j := 0; j < np; j++ {
			// 匹配
			if s[i] == p[j] || p[j] == '.' {
				dp[i+1][j+1] = dp[i][j]
			}
			if p[j] == '*' {
				// * 跟着他前一个字符走。
				// 前一个能匹配上 s[i]，* 才能有用，可以匹配 0次、一次、多次。
				// 前一个都不能匹配上 s[i]，* 也无能为力，只能让前一个字符消失，也就是匹配 0 次。
				if s[i] == p[j-1] || p[j-1] == '.' {
					zero := dp[i+1][j-1] // 匹配 0 次
					one := dp[i+1][j]    // 匹配 1 次
					many := dp[i][j+1]   // 匹配 多次
					dp[i+1][j+1] = zero || one || many
				} else {
					zero := dp[i+1][j-1] // 匹配 0 次
					dp[i+1][j+1] = zero
				}
			}
		}
	}
	return dp[ms][np]
}

func generateBoolDp(m, n int) [][]bool {
	dp := make([][]bool, m+1)
	for i, _ := range dp {
		dp[i] = make([]bool, n+1)
	}
	return dp
}

// 10. Regular Expression Matching
// 10. 正则表达式匹配
// 思路：Recursion + memo
// dp 函数定义：
// 1、p[0..j] 可以匹配 s[0..i]，则 dp(i,j) = true
// 2、p[0..j] 不能匹配 s[0..i]，则 dp(i,j) = false
// time O(MN) , space O(MN)
func isMatchM(s string, p string) bool {
	ms, np := len(s), len(p)
	if ms == 0 && np == 0 {
		return true
	}
	if ms == 0 || np == 0 {
		return false
	}

	// 备忘录
	memo := make(map[string]bool, 0)

	var dp func(i, j int) bool
	dp = func(i, j int) bool {
		// base case
		// p,s 都恰好走到末尾，则匹配成功
		if j == np {
			return i == ms
		}
		// s 走到末尾，p 还未结束。
		// 则 p[j..] 可以匹配空串，则匹配成功。
		if i == ms {
			// 能匹配空串：字符和"*"一起出现
			// 不是成对出现，则匹配失败。
			if (np-j)%2 == 1 {
				return false
			}
			for ; j+1 < np; j += 2 {
				if p[j+1] != '*' {
					return false
				}
			}
			return true
		}
		key := fmt.Sprintf("%d", i) + "," + fmt.Sprintf("%d", j)

		// 查备忘录
		if val, ok := memo[key]; ok {
			return val
		}

		var ret bool

		// s[i] == p[j]
		if s[i] == p[j] || p[j] == '.' {
			// 匹配 0次 或 多次
			if j < np-1 && p[j+1] == '*' {
				zero := dp(i, j+2) // 匹配 0次
				many := dp(i+1, j) // 匹配 多次
				ret = zero || many // 只要其中一个匹配，则是匹配成功。
			} else {
				// 默认匹配 1次
				ret = dp(i+1, j+1)
			}

		} else {
			// s[i] != p[j]
			// 后面没有通配符，则匹配 0 次。
			if j < np-1 && p[j+1] == '*' {
				// 匹配 0 次
				ret = dp(i, j+2)
			} else {
				// 默认后面没有通配符，则不匹配。
				ret = false
			}
		}
		// 结果存储到备忘录
		memo[key] = ret
		return ret
	}
	return dp(0, 0)
}

// 10. Regular Expression Matching
// 10. 正则表达式匹配
// 思路：Recursion
// dp 函数定义：
// 1、p[0..j] 可以匹配 s[0..i]，则 dp(i,j) = true
// 2、p[0..j] 不能匹配 s[0..i]，则 dp(i,j) = false
// time O(2^MN) , space O(2^MN)
func isMatchR(s string, p string) bool {
	ms, np := len(s), len(p)
	if ms == 0 && np == 0 {
		return true
	}
	if ms == 0 || np == 0 {
		return false
	}

	var dp func(i, j int) bool
	dp = func(i, j int) bool {
		// base case
		// p,s 都恰好走到末尾，则匹配成功
		if j == np {
			return i == ms
		}
		// s 走到末尾，p 还未结束。
		// 则 p[j..] 可以匹配空串，则匹配成功。
		if i == ms {
			// 能匹配空串：字符和"*"一起出现
			// 不是成对出现，则匹配失败。
			if (np-j)%2 == 1 {
				return false
			}
			for ; j+1 < np; j += 2 {
				if p[j+1] != '*' {
					return false
				}
			}
			return true
		}
		// s[i] == p[j]
		if s[i] == p[j] || p[j] == '.' {
			// 匹配 0次 或 多次
			if j < np-1 && p[j+1] == '*' {
				zero := dp(i, j+2)  // 匹配 0次
				many := dp(i+1, j)  // 匹配 多次
				return zero || many // 只要其中一个匹配，则是匹配成功。
			}
			// 匹配 1次
			return dp(i+1, j+1)
		}
		// s[i] != p[j]
		if j < np-1 && p[j+1] == '*' {
			// 匹配 0 次
			return dp(i, j+2)
		}
		// s[i] != p[j] 并且后面也没有通配符，则不匹配。
		return false
	}

	return dp(0, 0)
}
