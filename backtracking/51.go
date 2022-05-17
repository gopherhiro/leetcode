package backtracking

// 51. N-Queens
// 51. N 皇后
// time O(N^(N+1))
func solveNQueens(n int) (res [][]string) {
	if n == 0 {
		return
	}
	board := generateBoard(n)
	var backtrack func(row int)
	backtrack = func(row int) {
		// 结束条件
		if row == n {
			tmp := formatConversion(board)
			res = append(res, tmp)
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

func generateBoard(n int) [][]byte {
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}
	return board
}

func formatConversion(board [][]byte) []string {
	ans := make([]string, 0)
	for _, row := range board {
		str := ""
		for _, elem := range row {
			str += string(elem)
		}
		ans = append(ans, str)
	}
	return ans
}
