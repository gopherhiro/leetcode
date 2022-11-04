package array

// 59. Spiral Matrix II
// 59. 螺旋矩阵 II
// 思路：模拟
// time O(n^2) space O(1)
func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	top, bottom := 0, n-1
	left, right := 0, n-1

	total := n * n
	count := 1
	for count <= total {
		// traverse top boundary, top++
		for j := left; j <= right && count <= total; j++ {
			matrix[top][j] = count
			count++
		}
		top++

		// traverse right boundary, right--
		for i := top; i <= bottom && count <= total; i++ {
			matrix[i][right] = count
			count++
		}
		right--

		// traverse bottom boundary, bottom--
		for j := right; j >= left && count <= total; j-- {
			matrix[bottom][j] = count
			count++
		}
		bottom--

		// traverse left boundary, left++
		for i := bottom; i >= top && count <= total; i-- {
			matrix[i][left] = count
			count++
		}
		left++
	}
	return matrix
}
