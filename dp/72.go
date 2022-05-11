package dp

// 72. Edit Distance
// 72. 编辑距离
// 思路：DP(Dynamic Programming)
// time O(mn) space O(mn)
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}

	dp := genDP(m, n)

	// base case
	// dp[i][0] = i
	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}
	// dp[0][j] = j
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				// word1[i] != word2[j]
				// have three options:insert,delete,replace,
				// choose min
				insert := dp[i][j-1] + 1    // insert 之后的操作
				delete := dp[i-1][j] + 1    // delete 之后的操作
				replace := dp[i-1][j-1] + 1 // replace 之后的操作
				dp[i][j] = min(insert, min(delete, replace))
			}
		}
	}
	return dp[m][n]
}

func genDP(m, n int) [][]int {
	dp := make([][]int, m+1)
	for i, _ := range dp {
		dp[i] = make([]int, n+1)
	}
	return dp
}

// 72. Edit Distance
// 72. 编辑距离（遍历 word1,word2 从前往后）
// 思路：Recursion + Memo
// time O(mn) space O(mn)
func minDistanceO(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}
	memo := genMemo(m, n, -1)

	// 遍历 word1,word2 从后往前
	var dp func(i, j int) int
	dp = func(i, j int) int {
		// base case
		if i == m {
			return n - j
		}
		if j == n {
			return m - i
		}

		if memo[i][j] != -1 {
			return memo[i][j]
		}

		// if word1[i] == word2[j]
		if word1[i] == word2[j] {
			memo[i][j] = dp(i+1, j+1)
		} else {
			// if word1[i] != word2[j]
			// have three options： insert,delete,replace,
			// choose min
			insert := dp(i, j+1) + 1    // insert 之后的操作
			delete := dp(i+1, j) + 1    // delete 之后的操作
			replace := dp(i+1, j+1) + 1 // replace 之后的操作
			memo[i][j] = min(insert, min(delete, replace))
		}
		return memo[i][j]
	}
	return dp(0, 0)
}

// 72. Edit Distance
// 72. 编辑距离（遍历 word1,word2 从前往后）
// 思路：Recursion
// time O(2^n) space O(2^n)
func minDistanceR(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}

	// 遍历 word1,word2 从后往前
	var dp func(i, j int) int
	dp = func(i, j int) int {
		// base case
		if i == m {
			return n - j
		}
		if j == n {
			return m - i
		}
		// if word1[i] == word2[j]
		if word1[i] == word2[j] {
			return dp(i+1, j+1)
		}
		// if word1[i] != word2[j]
		// have three options： insert,delete,replace,
		// choose min
		insert := dp(i, j+1) + 1    // insert 之后的操作
		delete := dp(i+1, j) + 1    // delete 之后的操作
		replace := dp(i+1, j+1) + 1 // replace 之后的操作
		return min(insert, min(delete, replace))
	}
	return dp(0, 0)
}

// 72. Edit Distance
// 72. 编辑距离
// 思路：Recursion + Memo （遍历 word1,word2 从后往前）
// time O(m*n) space O(m*n)
func minDistanceM(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}

	// 备忘录
	memo := genMemo(m, n, -1)

	// 遍历 word1,word2 从后往前
	var dp func(i, j int) int
	dp = func(i, j int) int {
		// base case
		if i < 0 {
			return j + 1
		}
		if j < 0 {
			return i + 1
		}

		// 查询备忘录
		if memo[i][j] != -1 {
			return memo[i][j]
		}

		// if word1[i] == word2[j]
		if word1[i] == word2[j] {
			memo[i][j] = dp(i-1, j-1)
		} else {
			// if word1[i] != word2[j] have three options： insert,delete,replace,
			// choose min
			insert := dp(i, j-1) + 1    // insert 之后的操作
			delete := dp(i-1, j) + 1    // delete 之后的操作
			replace := dp(i-1, j-1) + 1 // replace 之后的操作
			memo[i][j] = min(insert, min(delete, replace))
		}
		return memo[i][j]
	}
	return dp(m-1, n-1)
}

func genMemo(m, n, initVal int) [][]int {
	memo := make([][]int, m)
	for i, _ := range memo {
		memo[i] = make([]int, n)
		for j, _ := range memo[i] {
			memo[i][j] = initVal
		}
	}
	return memo
}

// 72. Edit Distance
// 72. 编辑距离
// 思路：Recursion（遍历 word1,word2 从后往前）
// time O(2^n) space O(2^n)
func minDistanceR1(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}

	// 遍历 word1,word2 从后往前
	var dp func(i, j int) int
	dp = func(i, j int) int {
		// base case
		if i < 0 {
			return j + 1
		}
		if j < 0 {
			return i + 1
		}
		// if word1[i] == word2[j]
		if word1[i] == word2[j] {
			return dp(i-1, j-1)
		}
		// if word1[i] != word2[j] have three options： insert,delete,replace,
		// choose min
		insert := dp(i, j-1) + 1    // insert 之后的操作
		delete := dp(i-1, j) + 1    // delete 之后的操作
		replace := dp(i-1, j-1) + 1 // replace 之后的操作
		return min(insert, min(delete, replace))
	}

	return dp(m-1, n-1)
}
