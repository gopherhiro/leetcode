package backtracking

// 52. N-Queens II
// 52. N皇后 II
func totalSolveNQueens(n int) (ans int) {
	if n == 0 {
		return
	}
	board := generateBoard(n)
	var backtrack func(row int)
	backtrack = func(row int) {
		// 结束条件
		if row == n {
			ans++
			return
		}

		// 选择列表
		for col := 0; col < n; col++ {
			// 剪枝：不合法结果
			if !qIsValid(board, row, col, n) {
				continue
			}

			// 做出选择
			board[row][col] = 'Q'
			// 进入下一层决策树
			backtrack(row + 1)
			// 撤销选择
			board[row][col] = '.'
		}
	}
	backtrack(0)
	return
}

// 检测皇后的位置，是否合法。
// 因为Q是从上往下，一行一行放置的。
// 所以，只用检查：左上方、正上方、右上方 三个方向。
func qIsValid(board [][]byte, row, col, n int) bool {
	// 检查所在列是否合法
	for i := 0; i < row; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}
	// 检查右上方是否合法
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	// 检查左上方适合合法
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	return true
}
