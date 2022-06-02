package DFS

// 1430. Path With Given Sequence
// 1430. 判断给定的序列是否是二叉树从根到叶的路径
// 策略：DFS
// time O(N) space O(N)
func isValidSequence(root *TreeNode, arr []int) bool {
	if root == nil {
		return false
	}
	var dfs func(root *TreeNode, i int) bool
	dfs = func(root *TreeNode, i int) bool {
		if root == nil || i == len(arr) {
			return false
		}
		// 节点值对应数组下标值不等，则非对应路径。
		if root.Val != arr[i] {
			return false
		}
		// 到达叶子结点，判断是否同时到达数组末端
		// 是，则存在对应路径
		if root.Left == nil && root.Right == nil {
			return i == len(arr)-1
		}
		// 递归遍历左右子树，只要其中一颗子树找到对应路径，则满足条件。
		return dfs(root.Left, i+1) || dfs(root.Right, i+1)
	}
	return dfs(root, 0)
}
