package array

// 2133. Check if Every Row and Column Contains All Numbers
// 2133. 检查是否每一行每一列都包含全部整数
// 思路：哈希表 + 一次遍历
// time O(n) space O(n)
func checkValid(matrix [][]int) bool {
	n := len(matrix)
	row, col := make([][]bool, n), make([][]bool, n)
	for i := 0; i < n; i++ {
		row[i] = make([]bool, n+1)
		col[i] = make([]bool, n+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			number := matrix[i][j]
			// 如果出现重复，则不能包含全部整数，故而验证之即可。
			if row[i][number] || col[j][number] {
				return false
			}
			// 标记 number 出现过
			row[i][number] = true
			col[j][number] = true
		}
	}
	return true
}
