package number

// 118. Pascal's Triangle
// 118. 杨辉三角
// 思路：数学DP
// time O(n^2) space O(n^2)
func generate(rows int) [][]int {
	ans := make([][]int, rows)
	for i := range ans {
		ans[i] = make([]int, i+1)
		ans[i][0] = 1 // 每行第一个元素
		ans[i][i] = 1 // 每行最后一个元素
		// 中间元素：迭代计算出来
		for j := 1; j < i; j++ {
			ans[i][j] = ans[i-1][j] + ans[i-1][j-1]
		}
	}
	return ans
}
