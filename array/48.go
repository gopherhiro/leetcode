package array

// 48. Rotate Image
// 48. 旋转图像
// 思路：transpose & reflect
// 只是交换元素的位置，就不需要额外的空间
// time O(M) space O(1)
func rotate(matrix [][]int) {
	transpose(matrix)
	reflect(matrix)
}

func transpose(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = tmp
		}
	}
}

func reflect(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[i][n-j-1]
			matrix[i][n-j-1] = tmp
		}
	}
}

// 48. Rotate Image
// 48. 旋转图像
// 思路：模拟
// time O(n*n) space O(n*n)
func rotate2(matrix [][]int) {
	n := len(matrix)
	tmp := make([][]int, len(matrix))
	for i := 0; i < len(tmp); i++ {
		tmp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			tmp[j][n-i-1] = matrix[i][j]
		}
	}
	copy(matrix, tmp)
}
