package backtracking

// 79. Word Search
// 79. 单词搜索
// 思路：DFS + 回溯
// time O(mnp) space O(p)
func exist(board [][]byte, word string) bool {
	found := false
	m, n := len(board), len(board[0])
	var dfs func(i, j, k int)
	dfs = func(i, j, k int) {
		// 超出索引范围
		if i < 0 || j < 0 || i >= m || j >= n {
			return
		}
		// 走过，不能再走
		if board[i][j] == '*' {
			return
		}
		// 元素不相等
		if board[i][j] != word[k] {
			return
		}
		// 元素相等 && 长度相等，标记找到
		if k == len(word)-1 {
			found = true
			return
		}
		// 标记走过
		tmp := board[i][j]
		board[i][j] = '*'

		// 继续往后走
		dfs(i-1, j, k+1)
		dfs(i+1, j, k+1)
		dfs(i, j-1, k+1)
		dfs(i, j+1, k+1)

		// 从 [i,j] 为中心，向四个方向发散走，进行验证
		// 走完之后需回溯状态，以便于下一个点的验证
		board[i][j] = tmp
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			k := 0 // index of the word
			dfs(i, j, k)
		}
	}
	return found
}
