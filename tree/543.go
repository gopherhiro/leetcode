package tree

// 543. Diameter of Binary TreeNode
// 思路：所谓二叉树的直径，就是左右子树的最大深度之和，
// 那么直接的想法是对每个节点计算左右子树的最大高度，
// 得出每个节点的直径，从而得出最大的那个直径。
// time O(N), space O(N)
func diameterOfBinaryTree(root *TreeNode) int {
	var maxDiameter int

	var maxDepth func(*TreeNode) int
	maxDepth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := maxDepth(root.Left)
		right := maxDepth(root.Right)
		// 二叉树的直径：左右子树的最大深度之和，比较&获取最大直径。
		maxDiameter = maxValue(maxDiameter, left+right)
		return maxValue(left, right) + 1
	}
	maxDepth(root)

	return maxDiameter
}
