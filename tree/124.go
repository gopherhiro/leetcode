package tree

import "math"

// 124. Binary TreeNode Maximum Path Sum
// 定义：给你一个二叉树的根节点 root ，返回其 最大路径和。
// 路径和 是路径中各节点值的总和。
// 递归-分解问题：time O(N), space O(N)
// 参考题解：https://leetcode.cn/problems/binary-tree-maximum-path-sum/solutions/297276/shou-hui-tu-jie-hen-you-ya-de-yi-dao-dfsti-by-hyj8/
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	result := math.MinInt32

	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		// 取得左右子树的最大路径和。
		left := max(dfs(root.Left), 0)
		right := max(dfs(root.Right), 0)

		// 后序遍历位置，更新最大路径和
		sum := left + root.Val + right
		result = max(result, sum)

		// 左右子树的最大路径和加上根节点的值
		// 就是从根节点 root 为起点的最大路径和
		return root.Val + max(left, right)
	}
	dfs(root)
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
