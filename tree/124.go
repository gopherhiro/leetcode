package tree

import "math"

// 124. Binary TreeNode Maximum Path Sum
// 定义：给你一个二叉树的根节点 root ，返回其 最大路径和。
// 路径和 是路径中各节点值的总和。
// 递归-分解问题：time O(N), space O(N)
var maxPathSumRes = math.MinInt16

func maxPathSum(root *TreeNode) int {
	maxPathSumOneSide(root)
	return maxPathSumRes
}

func maxPathSumOneSide(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 取得左右子树的最大单边路径和。
	left := maxValue(maxPathSumOneSide(root.Left), 0)
	right := maxValue(maxPathSumOneSide(root.Right), 0)

	// 后序遍历位置，顺便更新最大路径和
	sum := left + root.Val + right
	maxPathSumRes = maxValue(maxPathSumRes, sum)

	// 实现函数定义，左右子树的最大单边路径和加上根节点的值
	// 就是从根节点 root 为起点的最大单边路径和
	return maxValue(left, right) + root.Val
}
