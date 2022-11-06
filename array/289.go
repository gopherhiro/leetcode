package array

// 289. Game of Life
// 289. 生命游戏
// 思路：模拟
// time O(mn) space O(1)
func gameOfLife(board [][]int) {
	if len(board) == 0 {
		return
	}
	// we use 2 bits to store 2 states:
	// 0 <- 00  dead (next) <- dead (current)
	// 1 <- 01  dead (next) <- live (current)
	// 2 <- 10  live (next) <- dead (current)
	// 3 <- 11  live (next) <- live (current)
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			lives := liveNeighbors(board, i, j)
			// In the beginning, every 2nd bit is 0;
			// So we only need to care about when will the 2nd bit become 1.
			if isLive(board[i][j]) && (lives == 2 || lives == 3) {
				board[i][j] = 3 // Make the 2nd bit 1: 01 ---> 11
			}
			if isDead(board[i][j]) && lives == 3 {
				board[i][j] = 2 // Make the 2nd bit 0: 00 ---> 10
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			board[i][j] >>= 1 // Get the 2nd state
		}
	}
}

func liveNeighbors(board [][]int, i, j int) int {
	count := 0
	m, n := len(board), len(board[0])

	// eight directions
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
	for _, dir := range dirs {
		x := i + dir[0]
		y := j + dir[1]
		if x < 0 || y < 0 || x >= m || y >= n {
			continue
		}
		count += board[x][y] & 1
	}
	return count
}

func isLive(c int) bool {
	return c == 1
}

func isDead(c int) bool {
	return c == 0
}
