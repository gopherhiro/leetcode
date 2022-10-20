package array

// 36. Valid Sudoku
// 36. 有效的数独
// 思路：哈希表 + 一次遍历
// time O(1) space O(1)
func isValidSudoku(board [][]byte) bool {
	row, col, box := [10][10]bool{}, [10][10]bool{}, [10][10]bool{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			// 转换数据 & 计算 box 的下标
			number := board[i][j] - '0'
			k := i/3*3 + j/3 // 难点：在于确定 number 所在 sub-boxes

			// 判断是否出现重复
			if row[i][number] || col[j][number] || box[k][number] {
				return false
			}
			// 标记 number 出现过
			row[i][number] = true
			col[j][number] = true
			box[k][number] = true
		}
	}
	return true
}
