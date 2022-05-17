package tree

// 110. Balanced Binary Tree
// DFS post order traverse, time O(N), space O(N)
func isBalanced(root *TreeNode) bool {
	answer := true

	var maxDepth func(root *TreeNode) int
	maxDepth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftMaxDepth := maxDepth(root.Left)
		rightMaxDepth := maxDepth(root.Right)
		// 后序遍历位置
		// 如果左右子树最大深度大于 1，就不是平衡二叉树
		if absv(leftMaxDepth-rightMaxDepth) > 1 {
			answer = false
		}
		return 1 + maxValue(leftMaxDepth, rightMaxDepth)
	}

	maxDepth(root)

	return answer
}

func absv(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
