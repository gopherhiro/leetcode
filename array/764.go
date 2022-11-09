package array

// 764. Largest Plus Sign
// 764. 最大加号标志
// 思路：前缀数组
// time O(n^2) space O(n^2)
func orderOfLargestPlusSign(n int, mines [][]int) int {
	// create five n x n matrix
	grid := make([][]int, n)
	top, left, bottom, right := make([][]int, n), make([][]int, n), make([][]int, n), make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, n)
		top[i], left[i], bottom[i], right[i] = make([]int, n), make([]int, n), make([]int, n), make([]int, n)
	}
	// grid's matrix having all values 1
	// assign the mines values as 0
	// so do top,left,bottom,right
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			grid[i][j], top[i][j], left[i][j], bottom[i][j], right[i][j] = 1, 1, 1, 1, 1
		}
	}
	for _, v := range mines {
		i, j := v[0], v[1]
		grid[i][j], top[i][j], left[i][j], bottom[i][j], right[i][j] = 0, 0, 0, 0, 0
	}

	// Fill the above four matrix
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			// from left to right and from top to bottom
			if i > 0 && top[i][j] == 1 {
				top[i][j] += top[i-1][j]
			}
			if j > 0 && left[i][j] == 1 {
				left[i][j] += left[i][j-1]
			}
			// from right to left and from bottom to top
			x := n - i - 1
			y := n - j - 1
			if x < n-1 && bottom[x][y] == 1 {
				bottom[x][y] += bottom[x+1][y]
			}
			if y < n-1 && right[x][y] == 1 {
				right[x][y] += right[x][y+1]
			}
		}
	}

	answer := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			k := minOfArr([]int{top[i][j], left[i][j], bottom[i][j], right[i][j]})
			answer = max(answer, k)
		}
	}
	return answer
}

func minOfArr(arr []int) int {
	ans := arr[0]
	for _, v := range arr {
		ans = min(ans, v)
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
