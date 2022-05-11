package dp

import "sort"

// 354. Russian Doll Envelopes
// 354. 俄罗斯套娃信封问题
// 思路：最长递增子序列
// 1、根据 w 升序排序，if w 相等，则根据 h 降序排列
// 2、把所有的 h作为一个数组，寻找[最长递增子序列]即是[最大信封嵌套个数]
// time O(N*logN) space O(N)
func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) == 0 {
		return 0
	}
	// 根据 w 升序排序，if w 相等，则根据 h 降序排列
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})

	// 获取 height 数组
	height := make([]int, len(envelopes))
	for i := 0; i < len(envelopes); i++ {
		height[i] = envelopes[i][1]
	}
	// 求 height longest incr subsequence
	return lengthOfLISBinary(height)
}
