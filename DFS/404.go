package DFS

// 404. Sum of Left Leaves
// 404. 左叶子之和
// 思路：DFS-前序遍历。遍历二叉树，累计左子节点之和即可。
// 左子节点：一个节点是左节点，并且是叶子节点。
// time O(N) space O(N)
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 判断是否是叶子节点
	isLeaf := func(root *TreeNode) bool {
		return root.Left == nil && root.Right == nil
	}

	sum := 0
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		// 是左叶子节点
		if root.Left != nil && isLeaf(root.Left) {
			sum += root.Left.Val
		}
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)

	return sum
}

func sumOfLeftLeavesM(root *TreeNode) int {
	if root == nil {
		return 0
	}

	sum := 0
	var dfs func(root *TreeNode, direction int)
	dfs = func(root *TreeNode, direction int) {
		if root == nil {
			return
		}
		// 是左叶子节点
		if root.Left == nil && root.Right == nil && direction == 0 {
			sum += root.Val
		}
		dfs(root.Left, 0)  // 0 means left direction
		dfs(root.Right, 1) // 1 means right direction
	}
	dfs(root, -1)

	return sum
}
