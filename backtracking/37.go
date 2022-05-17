package backtracking

// 37. Sudoku Solver
// 37. 解数独
// 思路：回溯
// time O(9^N), spaceO(N)
func solveSudoku(board [][]byte) {
	if len(board) == 0 {
		return
	}

	m, n := 9, 9 // 9 x 9 boxes
	var backtrack func(i, j int) bool
	backtrack = func(i, j int) bool {
		// 结束条件
		// 穷举到最后一列的话就换到下一行重新开始
		if j == n {
			return backtrack(i+1, 0)
		}

		// 完成了所有的穷举，则结束
		// 即找到一个可行解，触发 base case
		if i == m {
			return true
		}

		// 该位置是预设的数字，则不用管
		if board[i][j] != '.' {
			return backtrack(i, j+1)
		}

		// 遍历选择列表，做出选择
		for ch := byte('1'); ch <= byte('9'); ch++ {
			// 剪枝：过滤不合法的数字。
			// 即在同一行或同一列或同一个 3×3 的区域中存在相同的数字
			if !isValid(board, i, j, ch) {
				continue
			}

			// 做出选择
			board[i][j] = ch

			// 进入下一层决策树
			// 如果找到一个可行解，立即结束
			if backtrack(i, j+1) {
				return true
			}

			// 撤销选择
			board[i][j] = '.'

		}
		return false
	}
	backtrack(0, 0)
}

// 判断在同一行或同一列或同一个 3×3 的区域中存在相同的数字
func isValid(board [][]byte, row, col int, number byte) bool {
	for i := 0; i < 9; i++ {
		// whether the row is repeated
		if board[row][i] == number {
			return false
		}
		// whether the col is repeated
		if board[i][col] == number {
			return false
		}
		// Check for duplicates in 3 x 3 boxes
		r := row/3*3 + i/3
		c := col/3*3 + i%3
		if board[r][c] == number {
			return false
		}
	}
	return true
}
