package DFS

// 112. Path Sum
// 112. 路径总和
// 思路：DFS 遍历二叉树
// 1、遍历每一个节点，使用 targetSum 减去节点值
// 2、判断 targetSum 是否为0，并且是否达到叶子节点
// time O(n) space O(n)
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	found := false

	var dfs func(root *TreeNode, remain int)
	dfs = func(root *TreeNode, remain int) {
		if root == nil {
			return
		}
		remain -= root.Val
		if remain == 0 && root.Left == nil && root.Right == nil {
			found = true
			return
		}
		dfs(root.Left, remain)
		dfs(root.Right, remain)
	}

	dfs(root, targetSum)
	return found
}

// 思路：DFS
// 遍历二叉树，维护一个当前节点和，与targetSum比较即可。
func hasPathSumM(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	found := false
	curSum := 0

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		// 前序位置
		curSum += root.Val
		if root.Left == nil && root.Right == nil {
			if curSum == targetSum {
				found = true
			}
		}

		dfs(root.Left)
		dfs(root.Right)

		// 后序位置
		curSum -= root.Val
	}

	dfs(root)
	return found
}

func hasPathSumN(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}
	targetSum -= root.Val
	return hasPathSumN(root.Left, targetSum) || hasPathSumN(root.Right, targetSum)
}
