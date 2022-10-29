package array

// 1886. Determine Whether Matrix Can Be Obtained By Rotation
// 1886. 判断矩阵经轮转后是否一致
// 思路：枚举 0, 90, 180, 270 对应的旋转角度比较即可。
// time O(n^2) space O(1)
func findRotation(mat [][]int, target [][]int) bool {
	// 默认 mat 旋转 0, 90, 180, 270 度都是可以匹配 target 的
	r := [4]bool{true, true, true, true}
	n := len(mat)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			// 旋转 0 度
			if mat[i][j] != target[i][j] {
				r[0] = false
			}
			// 旋转 90 度
			if mat[i][j] != target[j][n-i-1] {
				r[1] = false
			}
			// 旋转 180 度
			if mat[i][j] != target[n-i-1][n-j-1] {
				r[2] = false
			}
			// 旋转 270 度
			if mat[i][j] != target[n-j-1][i] {
				r[3] = false
			}
		}
	}
	// 只要其中一个角度匹配，即表示OK
	return r[0] || r[1] || r[2] || r[3]
}
