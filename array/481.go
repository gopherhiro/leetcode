package array

// 885. Spiral Matrix III
// 885. 螺旋矩阵 III
// 思路：模拟
func spiralMatrixIII(R int, C int, r0 int, c0 int) [][]int {
	// Examining the lengths of our walk in each direction,
	// we find the following pattern: 1, 1, 2, 2, 3, 3, 4, 4, ...
	// That is, we walk 1 unit right, then 1 unit bottom,
	// then 2 units left, then 2 units top, etc.

	count := 0
	total := R * C
	answer := make([][]int, 0, total)

	answer = append(answer, []int{r0, c0})
	count++

	// number of step should go in each direction
	distances := 1
	for count < total {
		// move right
		for i := 1; i <= distances; i++ {
			c0++
			takeIt(&answer, &count, r0, c0, R, C)
		}
		// move bottom
		for i := 1; i <= distances; i++ {
			r0++
			takeIt(&answer, &count, r0, c0, R, C)
		}
		distances++ // turn twice direction, distances + 1

		// move left
		for i := 1; i <= distances; i++ {
			c0--
			takeIt(&answer, &count, r0, c0, R, C)
		}
		// move top
		for i := 1; i <= distances; i++ {
			r0--
			takeIt(&answer, &count, r0, c0, R, C)
		}
		distances++
	}
	return answer
}

func takeIt(answer *[][]int, count *int, r0, c0, R, C int) {
	if !isOk(r0, c0, R, C) {
		return
	}
	*answer = append(*answer, []int{r0, c0})
	*count++
}

func isOk(r0, c0, R, C int) bool {
	return 0 <= r0 && r0 < R && 0 <= c0 && c0 < C
}

// 885. Spiral Matrix III
// 885. 螺旋矩阵 III
// 思路：模拟，Walk in a Spiral
// time O(max(R,C)^2) space O(R*C)
func spiralMatrixIII2(R int, C int, r0 int, c0 int) [][]int {
	// direction row : move on row, index need to change value
	// direction col : move on rol, index need to change value
	// east, south, west, north
	dr := []int{0, 1, 0, -1}
	dc := []int{1, 0, -1, 0}

	count := 0
	total := R * C
	answer := make([][]int, 0, total)

	answer = append(answer, []int{r0, c0})
	count++

	if total == 1 {
		return answer
	}

	// k step by 2, is for construct dk: 1,1,2,2 | 3,3,4,4 | 5,5,6,6
	for k := 1; k < 2*(R+C); k += 2 {

		// i: direction index
		for i := 0; i < 4; i++ {

			// number of steps in this direction
			dk := k + i/2

			// for each step in this direction...
			for j := 0; j < dk; j++ {

				// step in the i-th direction
				r0 += dr[i]
				c0 += dc[i]

				if r0 < 0 || c0 < 0 || r0 >= R || c0 >= C {
					continue
				}

				answer = append(answer, []int{r0, c0})
				count++
				if count == total {
					return answer
				}
			}
		}
	}
	return [][]int{}
}
