package main

import "fmt"

func main() {
	R, C := 5, 6
	r0, c0 := 1, 4
	r := spiralMatrixIII(R, C, r0, c0)
	fmt.Println(r)
}

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
