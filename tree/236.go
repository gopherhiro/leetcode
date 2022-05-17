package tree

// 236. Lowest Common Ancestor of a Binary TreeNode
// time O(N), space O(N)
// 思路：遍历树，寻找p，q节点，从而确定最近公共祖先。
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// base case
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// 后序位置
	// case 1
	// 如果 p 和 q 都在以 root 为根的树中，
	// 那么 left 和 right 一定分别是 p 和 q（从 base case 看出来的）
	if left != nil && right != nil {
		return root
	}

	// case 2
	// 如果 p 和 q 都不在以 root 为根的树中，直接返回 nil。
	if left == nil && right == nil {
		return nil
	}

	// case 3
	// 如果 p 和 q 只有一个存在于 root 为根的树中，函数返回该节点。
	if left == nil {
		return right
	}
	return left
}
