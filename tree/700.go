package tree

// 700. Search in a Binary Search TreeNode
// Recursion, time O(H), space O(H)
func searchBST(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}
	// 去左子树搜索
	if target < root.Val {
		return searchBST(root.Left, target)
	}
	// 去右子树搜索
	if target > root.Val {
		return searchBST(root.Right, target)
	}
	return root
}

// 700. Search in a Binary Search TreeNode
// Iteration, time O(N), space O(1)
func searchBSTIter(root *TreeNode, target int) *TreeNode {
	for root != nil {
		if root.Val == target {
			return root
		}
		// 特别注意这里：要么走左子树，要么走右子树。
		if target < root.Val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return nil
}
