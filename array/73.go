package array

// 73. Set Matrix Zeroes
// 73. 矩阵置零
// 思路：r0,c0 标记置零
// time O(mn) space O(1)
func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	// we can use the first cell of every row and column as a flag.
	r0HasZero := false
	c0HasZero := false
	// check the first row
	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			r0HasZero = true
			break
		}
	}

	// check the first column
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			c0HasZero = true
			break
		}
	}

	// we can use the first cell of every row and column as a flag.
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// checking the respective first row cell or first column cell.
	// If any of them was marked, we set the value in the cell to 0.
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// handle the first row
	if r0HasZero {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}

	// handle the first column
	if c0HasZero {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}

// 73. Set Matrix Zeroes
// 73. 矩阵置零
// 思路：标记零-两个集合
// time O(mn) space O(m+n)
func setZeroes2(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])

	// 收集元素值为 0 的坐标
	rows := make(map[int]bool, 0)
	cols := make(map[int]bool, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}
	// 元素值为 0 的行和列，赋值为 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if rows[i] || cols[j] {
				matrix[i][j] = 0
			}
		}
	}
}

// 73. Set Matrix Zeroes
// 73. 矩阵置零
// 思路：标记零-数对
// time O(mn) space O(mn)
func setZeroes1(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])

	// 收集元素值为 0 的坐标
	zeros := make([][]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				zeros = append(zeros, []int{i, j})
			}
		}
	}
	// 元素值为 0 的行和列，赋值为 0
	for _, zero := range zeros {
		i := zero[0]
		j := zero[1]

		// process rows
		for p := 0; p < m; p++ {
			matrix[p][j] = 0
		}
		// process cols
		for q := 0; q < n; q++ {
			matrix[i][q] = 0
		}
	}
}
