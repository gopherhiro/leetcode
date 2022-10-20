package array

// 1572. Matrix Diagonal Sum
// 1572. 矩阵对角线元素的和
// time O(n) space O(1)
func diagonalSum(mat [][]int) int {
	res := 0
	n := len(mat)
	for i := 0; i < n; i++ {
		res += mat[i][i]     // row == col
		res += mat[i][n-1-i] // row + col = n - 1
	}
	// if n is a odd number
	// that mean we have added the center element twice
	if n%2 == 1 {
		res -= mat[n/2][n/2]
	}
	return res
}
