package backtracking

// 17. Letter Combinations of a Phone Number
// 17. 电话号码的字母组合
// 思路：回溯算法
// time O(3^m * 4^n) space O(m+n)
// 备注：是有个意思的题，因为它的选择列表居然是变的
func letterCombinations(digits string) (ans []string) {
	if len(digits) == 0 {
		return
	}
	// 数字字符映射关系
	numLetterMap := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	// 已选择路径
	track := make([]byte, 0)

	var backtrack func(index int)
	backtrack = func(index int) {
		// 结束条件：遍历完所有的字符数字
		if index == len(digits) {
			tmp := make([]byte, len(track))
			copy(tmp, track)
			ans = append(ans, string(tmp))
			return
		}
		// 遍历选择列表，做选择
		// 获取选择列表(每一次的选择列表可能不一样，是根据映射关系取得的)
		// 是有个意思的题，因为它的选择列表居然是变的
		number := digits[index]
		letters := numLetterMap[number]
		for i := 0; i < len(letters); i++ {
			// 做出选择
			track = append(track, letters[i])

			// 进入下一层决策树
			backtrack(index + 1)

			// 撤销选择
			track = track[:len(track)-1]
		}
	}

	backtrack(0)

	return
}
