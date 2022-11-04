package array

// 54. Spiral Matrix
// 54. 螺旋矩阵
// 思路：模拟
// time O(mn) space O(1)
func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	up, down := 0, m-1
	left, right := 0, n-1

	// 在结果集长度在[0, m*n] 范围内进行遍历
	ans := make([]int, 0)
	total := m * n
	for len(ans) < total {
		// traverse up boundary, up++
		for j := left; j <= right && len(ans) < total; j++ {
			ans = append(ans, matrix[up][j])
		}
		up++

		// traverse right boundary, right--
		for i := up; i <= down && len(ans) < total; i++ {
			ans = append(ans, matrix[i][right])
		}
		right--

		// traverse down boundary, down--
		for j := right; j >= left && len(ans) < total; j-- {
			ans = append(ans, matrix[down][j])
		}
		down--

		// traverse left boundary, left++
		for i := down; i >= up && len(ans) < total; i-- {
			ans = append(ans, matrix[i][left])
		}
		left++
	}
	return ans
}

// 54. Spiral Matrix
// 54. 螺旋矩阵
// 思路：旋转矩阵
// time O(mmn) space O(mmn)
func spiralOrder2(matrix [][]int) []int {
	answer := make([]int, 0)

	var helper func(matrix [][]int)
	helper = func(matrix [][]int) {
		// 先把当前数组第一行 append 到结果集中
		// 把剩下剩下的元素，逆时针旋转 90
		// 重复以上操作，直到遍历完所有元素
		answer = append(answer, matrix[0]...)
		matrix = matrix[1:]
		if len(matrix) == 0 {
			return
		}
		matrix = rotate(matrix)
		helper(matrix)
	}
	helper(matrix)

	return answer
}

// 48. Rotate Image
// 48. 旋转图像（逆时针）
// 思路：模拟
// time O(mn) space O(mn)
func rotate(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	tmp := make([][]int, n)
	for i := 0; i < n; i++ {
		tmp[i] = make([]int, m)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			tmp[n-j-1][i] = matrix[i][j]
		}
	}
	return tmp
}
