package dp

import "leetcode/pkg"

// 337. House Robber III
// 337. 打家劫舍 III
// 思路：分层操作取舍(dynamic programming)
// time O(N) space O(N)
func rob(root *pkg.TreeNode) int {
	if root == nil {
		return 0
	}
	/* 返回一个大小为 2 的数组 arr
	arr[0] 表示不抢 root，得到的最大钱数
	arr[1] 表示抢 root，得到的最大钱数 */
	var dp func(root *pkg.TreeNode) []int
	dp = func(root *pkg.TreeNode) []int {
		if root == nil {
			return []int{0, 0}
		}
		res := make([]int, 2)
		left := dp(root.Left)
		right := dp(root.Right)

		// 不抢，下家可抢可不抢，取决于收益大小。
		res[0] = max(left[0], left[1]) + max(right[0], right[1])

		// 抢，下家就不能抢了。
		res[1] = root.Val + left[0] + right[0]

		return res
	}
	ans := dp(root)
	return max(ans[0], ans[1])
}

// 213. House Robber II
// 213. 打家劫舍 II
// 思路：首尾房间不能同时取钱的情况：00,01,10 (0不取，1取)
// 都是正数最小，00 case 可忽略，不参与选择。即：
// We can divide this problem to two sub problems:
// Let's take following example:
// Sub problem 1: rob house 1 ~ 8
// Sub problem 2: rob house 2 ~ 9
// And find the bigger one of these two sub problems.
// time O(N) space O(1)
func robII(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	aimHead := dp(nums, 0, n-2)
	aimTail := dp(nums, 1, n-1)
	return max(aimHead, aimTail)
}

func dp(nums []int, start, end int) int {
	dpi, dpi1, dpi2 := 0, 0, 0 // base case
	for i := end; i >= start; i-- {
		dpi = max(dpi1, nums[i]+dpi2)
		dpi2 = dpi1
		dpi1 = dpi
	}
	return dpi
}
