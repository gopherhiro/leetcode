package offer

// 剑指 Offer 28. 对称的二叉树
// 101. 对称二叉树
// 思路：left,right 左右节点对称。
// time O(n), space O(n)
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var helper func(L, R *TreeNode) bool
	helper = func(L, R *TreeNode) bool {
		if L == nil && R == nil {
			return true
		}
		if L == nil || R == nil {
			return false
		}
		if L.Val != R.Val {
			return false
		}
		return helper(L.Left, R.Right) && helper(L.Right, R.Left)
	}
	return helper(root, root)
}
