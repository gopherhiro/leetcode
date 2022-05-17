package tree

// 235. Lowest Common Ancestor of a Binary Search Tree
// 235. 二叉搜索树的最近公共祖先
// time O(logN) space O(N)
func lowestCommonAncestorII(root, p, q *TreeNode) *TreeNode {
	// base case
	if root == nil {
		return nil
	}

	if p.Val > q.Val {
		// 保证 p <= q，方便下面二分。
		return lowestCommonAncestorII(root, q, p)
	}

	// p <= root <= q
	if root.Val >= p.Val && root.Val <= q.Val {
		return root
	}

	// p,q < root
	if q.Val < root.Val {
		return lowestCommonAncestorII(root.Left, p, q)
	}
	// p,q > root
	return lowestCommonAncestorII(root.Right, p, q)
}
