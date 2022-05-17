package tree

// 104. Maximum Depth of Binary TreeNode
// 定义：输入一个节点，返回以该节点为根的二叉树的最大深度
// 动态规划思路：根据左右子树的最大深度推出原二叉树的最大深度即可。
// time O(N), space O(N)
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	// 根据左右子树的最大深度推出原二叉树的最大深度
	return maxValue(leftDepth, rightDepth) + 1
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
