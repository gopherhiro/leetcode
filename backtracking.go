package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

/********************
回溯复杂度：
时间复杂度：time O(n * 2^n)
空间复杂度：space O(n)
********************/

func main() {
	res := letterCombinations("23")
	fmt.Println(res)
}

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

// 52. N-Queens II
// 52. N皇后 II
func totalSolveNQueens(n int) (ans int) {
	if n == 0 {
		return
	}
	board := generateBoard(n)
	var backtrack func(row int)
	backtrack = func(row int) {
		// 结束条件
		if row == n {
			ans++
			return
		}

		// 选择列表
		for col := 0; col < n; col++ {
			// 剪枝：不合法结果
			if !QisValid(board, row, col, n) {
				continue
			}

			// 做出选择
			board[row][col] = 'Q'
			// 进入下一层决策树
			backtrack(row + 1)
			// 撤销选择
			board[row][col] = '.'
		}
	}
	backtrack(0)
	return
}

// 51. N-Queens
// 51. N 皇后
// time O(N^(N+1))
func solveNQueens(n int) (res [][]string) {
	if n == 0 {
		return
	}
	board := generateBoard(n)
	var backtrack func(row int)
	backtrack = func(row int) {
		// 结束条件
		if row == n {
			tmp := formatConversion(board)
			res = append(res, tmp)
			return
		}

		// 选择列表
		for col := 0; col < n; col++ {
			// 剪枝：不合法结果
			if !QisValid(board, row, col, n) {
				continue
			}

			// 做出选择
			board[row][col] = 'Q'
			// 进入下一层决策树
			backtrack(row + 1)
			// 撤销选择
			board[row][col] = '.'
		}
	}
	backtrack(0)
	return
}

func generateBoard(n int) [][]byte {
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}
	return board
}

func formatConversion(board [][]byte) []string {
	ans := make([]string, 0)
	for _, row := range board {
		str := ""
		for _, elem := range row {
			str += string(elem)
		}
		ans = append(ans, str)
	}
	return ans
}

// 检测皇后的位置，是否合法。
// 因为Q是从上往下，一行一行放置的。
// 所以，只用检查：左上方、正上方、右上方 三个方向。
func QisValid(board [][]byte, row, col, n int) bool {
	// 检查所在列是否合法
	for i := 0; i < row; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}
	// 检查右上方是否合法
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	// 检查左上方适合合法
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	return true
}

// 22. Generate Parentheses
// 22. 括号生成
func generateParenthesis(n int) (res []string) {
	if n == 0 {
		return
	}

	track := make([]string, 0)

	// left 记录还可以使用多少个左括号，
	// right 记录还可以使用多少个右括号。
	var backtrack func(left, right int)
	backtrack = func(left, right int) {
		// 剪枝：不合法的结果
		// 合法的括号组合中：剩余的左括号一定是 <= 右括号的
		if left > right {
			return
		}
		if left < 0 || right < 0 {
			return
		}
		// 结束条件
		// 当左右括号刚好用完，得到一个合法的括号组合。
		if left == 0 && right == 0 {
			tmp := make([]string, len(track))
			copy(tmp, track)
			res = append(res, strings.Join(tmp, ""))
		}

		// 选择列表，进行路径选择
		// 尝试选择左括号
		track = append(track, "(")
		backtrack(left-1, right)
		track = track[:len(track)-1]

		// 尝试选择右括号
		track = append(track, ")")
		backtrack(left, right-1)
		track = track[:len(track)-1]
	}
	// 刚开始时，左括号和右括号相等
	backtrack(n, n)
	return
}

// 37. Sudoku Solver
// 37. 解数独
// 思路：回溯
// time O(9^N), spaceO(N)
func solveSudoku(board [][]byte) {
	if len(board) == 0 {
		return
	}

	m, n := 9, 9 // 9 x 9 boxes
	var backtrack func(i, j int) bool
	backtrack = func(i, j int) bool {
		// 结束条件
		// 穷举到最后一列的话就换到下一行重新开始
		if j == n {
			return backtrack(i+1, 0)
		}

		// 完成了所有的穷举，则结束
		// 即找到一个可行解，触发 base case
		if i == m {
			return true
		}

		// 该位置是预设的数字，则不用管
		if board[i][j] != '.' {
			return backtrack(i, j+1)
		}

		// 遍历选择列表，做出选择
		for ch := byte('1'); ch <= byte('9'); ch++ {
			// 剪枝：过滤不合法的数字。
			// 即在同一行或同一列或同一个 3×3 的区域中存在相同的数字
			if !isValid(board, i, j, ch) {
				continue
			}

			// 做出选择
			board[i][j] = ch

			// 进入下一层决策树
			// 如果找到一个可行解，立即结束
			if backtrack(i, j+1) {
				return true
			}

			// 撤销选择
			board[i][j] = '.'

		}
		return false
	}
	backtrack(0, 0)
}

// 判断在同一行或同一列或同一个 3×3 的区域中存在相同的数字
func isValid(board [][]byte, row, col int, number byte) bool {
	for i := 0; i < 9; i++ {
		// whether the row is repeated
		if board[row][i] == number {
			return false
		}
		// whether the col is repeated
		if board[i][col] == number {
			return false
		}
		// Check for duplicates in 3 x 3 boxes
		r := row/3*3 + i/3
		c := col/3*3 + i%3
		if board[r][c] == number {
			return false
		}
	}
	return true
}

// 401. Binary Watch
// 401. 二进制手表
// 思路：此题是一个不错的回溯算法思想运用题。
func readBinaryWatch(turnedOn int) (res []string) {
	if turnedOn < 0 {
		return
	}

	// 可选列表：LED灯，前4个为小时，后6个为分钟。
	times := []int{8, 4, 2, 1, 32, 16, 8, 4, 2, 1}

	var backtrack func(turnedOn, start, hours, minutes int)
	backtrack = func(turnedOn, start, hours, minutes int) {
		// 结束条件
		if turnedOn == 0 {
			// 剪枝：丢弃不合法时间。
			if hours > 11 || minutes > 59 {
				return
			}
			var tmpTime string

			// 处理小时
			tmpTime += fmt.Sprintf("%d", hours)
			tmpTime += ":"

			// 处理分钟
			if minutes < 10 {
				tmpTime += fmt.Sprintf("0%d", minutes)
			} else {
				tmpTime += fmt.Sprintf("%d", minutes)
			}
			// 收集时间结果
			res = append(res, tmpTime)
			return
		}

		// hours 需要加起来
		// minutes 需要加起来

		for i := start; i < len(times); i++ {
			// 做出选择
			isHour := i < 4
			if isHour {
				hours += times[i]
			} else {
				minutes += times[i]
			}

			// 进入下一层决策树
			// 进入下一层决策树，LED灯打开的数量，需要减1。
			backtrack(turnedOn-1, i+1, hours, minutes)

			// 撤销选择
			if isHour {
				hours -= times[i]
			} else {
				minutes -= times[i]
			}
		}

	}

	backtrack(turnedOn, 0, 0, 0)

	return
}

// 216. Combination Sum III
// 216. 组合总和 III
// 思路：回溯
func combinationSum3(k int, n int) (res [][]int) {
	if k <= 0 || n <= 0 {
		return
	}
	// 记录已选择路径
	track := make([]int, 0)

	var backtrack func(start int)
	backtrack = func(start int) {
		// 结束条件
		if len(track) == k && n == 0 {
			tmp := make([]int, len(track))
			copy(tmp, track)
			res = append(res, tmp)
			return
		}

		if n < 0 {
			return
		}

		// 遍历选择列表，进行路径选择
		for i := start; i <= 9; i++ {
			// 剪枝：
			// track的长度加上区间 [start, n] 的长度小于 k
			// 不可能构造出长度为 k 的 track
			// 比如：start = n 时，
			if len(track)+(n-start)+1 < k {
				continue
			}

			// 做出选择
			n -= i
			track = append(track, i)

			// 进入下一层决策树
			backtrack(i + 1)

			// 撤销选择
			track = track[:len(track)-1]
			n += i
		}

	}
	backtrack(1)
	return
}

// 40. Combination Sum II
// 40. 组合总和II
// 思路：回溯+剪枝。
func combinationSum2(candidates []int, target int) (res [][]int) {
	if len(candidates) == 0 {
		return
	}

	sort.Ints(candidates)

	// 记录已选择的路径
	track := make([]int, 0)
	var backtrack func(start int)
	backtrack = func(start int) {
		// 结束条件
		if target == 0 {
			tmp := make([]int, len(track))
			copy(tmp, track)
			res = append(res, tmp)
			return
		}

		// 不符合条件的都丢弃。
		if target < 0 {
			return
		}

		for i := start; i < len(candidates); i++ {
			// 剪枝：值相同的树枝，只遍历第一条。
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}

			// 做出选择
			target -= candidates[i]
			track = append(track, candidates[i])

			// 进入下一层决策树
			// 往下一层传[i + 1]，可以限定每个数字只能使用一次。
			backtrack(i + 1)

			// 撤销选择
			track = track[:len(track)-1]
			target += candidates[i]
		}

	}
	backtrack(0)
	return
}

// 39. Combination Sum
// 39. 组合总和
// 思路：回溯+剪枝。
func combinationSum(candidates []int, target int) (res [][]int) {
	if len(candidates) == 0 {
		return
	}
	// 记录已选择的路径
	track := make([]int, 0)
	var backtrack func(start int)
	backtrack = func(start int) {
		// 结束条件
		if target == 0 {
			tmp := make([]int, len(track))
			copy(tmp, track)
			res = append(res, tmp)
			return
		}

		// 不符合条件的都丢弃。
		if target < 0 {
			return
		}

		for i := start; i < len(candidates); i++ {
			// 做出选择
			target -= candidates[i]
			track = append(track, candidates[i])

			// 进入下一层决策树
			// 往下一层传[i]，每一个数字都可以重复使用。
			backtrack(i)

			// 撤销选择
			track = track[:len(track)-1]
			target += candidates[i]
		}

	}
	backtrack(0)
	return
}

// 77. Combinations
// 77. 组合
// 思路：backtracking
// time O(N * 2^N) , space O(N)
func combine(n int, k int) (res [][]int) {
	if n <= 0 || k <= 0 {
		return
	}
	// 记录已选择的路径
	track := make([]int, 0)

	var backtrack func(start int)
	backtrack = func(start int) {
		if len(track) == k {
			tmp := make([]int, len(track))
			copy(tmp, track)
			res = append(res, tmp)
			return
		}

		// 使用 start 参数控制递归，进行剪枝。
		for i := start; i <= n; i++ {
			// 剪枝：
			// track的长度加上区间 [start, n] 的长度小于 k
			// 不可能构造出长度为 k 的 track
			// 比如：start = n 时，
			if len(track)+(n-start+1) < k {
				return
			}

			// 做选择
			track = append(track, i)

			// 进入下一层决策树
			backtrack(i + 1)

			// 撤销选择
			track = track[:len(track)-1]
		}

	}
	// n >= 1
	backtrack(1)
	return
}

// 1863. Sum of All Subset XOR Totals
// 1863. 找出所有子集的异或总和再求和
// 思路：
// 1、回溯求出所有的子集。
// 2、遍历所有子集，求异或并求和。
func subsetXORSumN(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sum := 0
	track := make([]int, 0)

	var backtrack func(nums []int, start int)
	backtrack = func(nums []int, start int) {

		// 结束条件
		subset := make([]int, len(track))
		copy(subset, track)

		// 计算子集异或并求和
		if len(subset) != 0 {
			xor := 0
			for _, elem := range subset {
				xor ^= elem
			}
			sum += xor
		}

		// 遍历选择列表，进行路径选择
		for i := start; i < len(nums); i++ {
			// 做出选择
			track = append(track, nums[i])

			// 进入下一层决策树
			backtrack(nums, i+1)

			// 撤销选择
			track = track[:len(track)-1]
		}
	}
	backtrack(nums, 0)

	return sum
}

// 1863. Sum of All Subset XOR Totals
// 1863. 找出所有子集的异或总和再求和
// 思路：
// 1、回溯求出所有的子集。
// 2、遍历所有子集，求异或并求和。
func subsetXORSum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	subsets := make([][]int, 0)
	track := make([]int, 0)

	var backtrack func(nums []int, start int)
	backtrack = func(nums []int, start int) {
		// 结束条件
		tmp := make([]int, len(track))
		copy(tmp, track)
		subsets = append(subsets, tmp)

		// 遍历选择列表，进行路径选择
		for i := start; i < len(nums); i++ {
			// 做出选择
			track = append(track, nums[i])

			// 进入下一层决策树
			backtrack(nums, i+1)

			// 撤销选择
			track = track[:len(track)-1]
		}
	}
	backtrack(nums, 0)

	// 计算子集异或并求和
	if len(subsets) == 0 {
		return 0
	}
	sum := 0
	for _, subset := range subsets {
		if len(subset) == 0 {
			continue
		}
		xor := 0
		for _, elem := range subset {
			xor ^= elem
		}
		sum += xor
	}
	return sum
}

// 90. Subsets II
// 90. 子集 II
// 思路：backtracking
// time O(N*2^N), space O(N)
func subsetsWithDup(nums []int) (res [][]int) {
	if len(nums) == 0 {
		return
	}

	// 先排序，让相同的元素在一起。
	sort.Ints(nums)

	// 记录已选择路径
	track := make([]int, 0)

	var backtrack func(nums []int, start int)
	backtrack = func(nums []int, start int) {
		// 前序位置：收集结果。
		tmp := make([]int, len(track))
		copy(tmp, track)
		res = append(res, tmp)

		// 使用 start 参数控制递归，进行剪枝。
		for i := start; i < len(nums); i++ {

			// 剪枝。值相同的相邻树枝，只遍历第一条。
			if i > start && nums[i] == nums[i-1] {
				continue
			}

			// 做选择
			track = append(track, nums[i])

			// 进入下一层抉择树
			backtrack(nums, i+1)

			// 撤销选择
			track = track[:len(track)-1]
		}

	}
	backtrack(nums, 0)
	return
}

// 78. Subsets
// 78. 子集
// 思路：回溯
// time O(N*2^N), space O(N)
func subsets(nums []int) (res [][]int) {
	if len(nums) == 0 {
		return
	}
	// 记录已选择路径
	track := make([]int, 0)

	var backtrack func(nums []int, start int)
	backtrack = func(nums []int, start int) {
		// 前序位置：收集结果。
		tmp := make([]int, len(track))
		copy(tmp, track)
		res = append(res, tmp)

		// 使用 start 参数控制递归，进行剪枝。
		for i := start; i < len(nums); i++ {
			// 做选择
			track = append(track, nums[i])

			// 进入下一层抉择树
			backtrack(nums, i+1)

			// 撤销选择
			track = track[:len(track)-1]
		}

	}
	backtrack(nums, 0)
	return
}

// 列举所有全排列，同一个位置的数可以使用多次。
// time O(n*n!), space O(n)
func permuteBase(nums []int) (ans [][]int) {
	track := make([]int, 0)

	var backtrack func(nums []int)
	backtrack = func(nums []int) {
		// 结束条件：已选择路径与原始可选列表长度相等时
		// 即表示得到了一个全排列
		if len(track) == len(nums) {
			tmp := make([]int, len(track))
			copy(tmp, track)
			ans = append(ans, tmp)
			return
		}

		for _, num := range nums {
			// 为了维护每个节点的选择列表和路径，需要：
			// 递归之前做出选择，递归之后撤销选择，即回溯。
			// 做出选择
			track = append(track, num)

			// 进入下一层决策树
			backtrack(nums)

			// 撤销选择
			track = track[:len(track)-1]
		}

	}

	backtrack(nums)

	return
}

// 46. Permutations
// 46. 全排列
// 思路：回溯
// time O(n*n!), space O(n)
func permute(nums []int) (ans [][]int) {
	track := make([]int, 0)

	var backtrack func(nums []int)
	backtrack = func(nums []int) {
		// 结束条件：已选择路径与原始可选列表长度相等时
		// 即表示得到了一个全排列
		if len(track) == len(nums) {
			tmp := make([]int, len(track))
			copy(tmp, track)
			ans = append(ans, tmp)
			return
		}

		for _, num := range nums {
			// 剪枝，排除不合理的情况。
			if _, ok := find(track, num); ok {
				continue
			}
			// 为了维护每个节点的选择列表和路径，需要：
			// 递归之前做出选择，
			// 递归之后撤销选择。

			// 做出选择
			track = append(track, num)

			// 进入下一层决策树
			backtrack(nums)

			// 撤销选择
			track = track[:len(track)-1]
		}

	}

	backtrack(nums)

	return
}

// Find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
// time O(N)
func find(s []int, val int) (int, bool) {
	for i, item := range s {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// 46. Permutations
// 46. 全排列
// 思路：回溯，优化了剪枝时的时间复杂度：O(N) -> O(1)
// time O(n*n!), space O(n)
func permuteMap(nums []int) (ans [][]int) {
	track := make([]int, 0)
	// 空间换时间：使用map辅助，剪枝的时间复杂度可降为O(1)
	used := make(map[int]bool)

	var backtrack func(nums []int, track []int)
	backtrack = func(nums []int, track []int) {
		// 结束条件：已选择路径与原始可选列表长度相等时
		// 即表示得到了一个全排列
		if len(track) == len(nums) {
			tmp := make([]int, len(track))
			copy(tmp, track)
			ans = append(ans, tmp)
			return
		}

		for _, num := range nums {
			// 剪枝，排除不合理的情况。
			if used[num] {
				continue
			}
			// 为了维护每个节点的选择列表和路径，需要：
			// 递归之前做出选择，
			// 递归之后撤销选择。

			// 做出选择
			track = append(track, num)
			used[num] = true

			// 进入下一层决策树
			backtrack(nums, track)

			// 撤销选择
			track = track[:len(track)-1]
			used[num] = false
		}

	}

	backtrack(nums, track)

	return
}

// 47. Permutations
// 47. 全排列
// 思路：回溯
// time O(n*n!), space O(n)
func permuteUnique(nums []int) (ans [][]int) {
	// 先排序，让相同的元素在一起
	sort.Ints(nums)

	// 记录路径
	track := make([]int, 0)

	// 空间换时间：使用map辅助，剪枝的时间复杂度可降为O(1)
	used := make(map[int]bool)

	var backtrack func(nums []int)
	backtrack = func(nums []int) {
		// 结束条件：已选择路径与原始可选列表长度相等时
		// 即表示得到了一个全排列
		if len(track) == len(nums) {
			tmp := make([]int, len(track))
			copy(tmp, track)
			ans = append(ans, tmp)
			return
		}

		for i, num := range nums {
			// 剪枝，排除不合理的情况。
			// 该下标元素是否使用过。
			if used[i] {
				continue
			}
			// 新增额外剪枝：把重复元素的排列去掉。
			// 固定相同的元素在排列中的相对位置：
			// 即如果前面的相邻 相等元素没有用过，则剪枝。
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}

			// 为了维护每个节点的选择列表和路径，需要：
			// 递归之前做出选择，
			// 递归之后撤销选择。

			// 做出选择
			track = append(track, num)
			used[i] = true

			// 进入下一层决策树
			backtrack(nums)

			// 撤销选择
			track = track[:len(track)-1]
			used[i] = false
		}

	}

	backtrack(nums)

	return
}

// 47. Permutations
// 47. 全排列
// 思路：回溯
// time O(n*n!), space O(n)
func permuteUniqueII(nums []int) (ans [][]int) {
	// 先排序，让相同的元素在一起
	sort.Ints(nums)

	// 记录路径
	track := make([]int, 0)

	// 空间换时间：使用map辅助，剪枝的时间复杂度可降为O(1)
	used := make(map[int]bool)

	var backtrack func(nums []int)
	backtrack = func(nums []int) {
		// 结束条件：已选择路径与原始可选列表长度相等时
		// 即表示得到了一个全排列
		if len(track) == len(nums) {
			tmp := make([]int, len(track))
			copy(tmp, track)
			ans = append(ans, tmp)
			return
		}

		preNumber := math.MinInt8
		for i, num := range nums {
			// 剪枝，排除不合理的情况。
			if used[i] {
				continue
			}

			// 新增额外剪枝：把重复元素的排列去掉。
			// 要选的元素是否和上一次选择过的相等。
			if num == preNumber {
				continue
			}

			// 为了维护每个节点的选择列表和路径，需要：
			// 递归之前做出选择，
			// 递归之后撤销选择。

			// 做出选择
			track = append(track, num)
			used[i] = true

			// 记录之前树枝上元素的值
			preNumber = num

			// 进入下一层决策树
			backtrack(nums)

			// 撤销选择
			track = track[:len(track)-1]
			used[i] = false
		}

	}

	backtrack(nums)

	return
}
