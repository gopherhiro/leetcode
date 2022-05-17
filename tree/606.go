package tree

import "fmt"

// 606. Construct String from Binary TreeNode
// [root,#,#] => root
// [root,#,right] => root()(right)
// [root,left,#] => root(left)
// [root,left,right] => root(left)(right)
// Approach: Recursion, 分解问题: time O(N), space O(N)
func tree2str(root *TreeNode) string {
	// base case
	if root == nil {
		return ""
	}
	// [root,#,#] => root
	if root.Left == nil && root.Right == nil {
		return fmt.Sprintf("%d", root.Val) + ""
	}

	leftStr := tree2str(root.Left)
	rightStr := tree2str(root.Right)
	// 后序位置
	// [root,#,right] => root()(right)
	if root.Left == nil && root.Right != nil {
		return fmt.Sprintf("%d", root.Val) + "()" + "(" + rightStr + ")"
	}

	// [root,left,#] => root(left)
	if root.Left != nil && root.Right == nil {
		return fmt.Sprintf("%d", root.Val) + "(" + leftStr + ")"
	}

	// [root,left,right] => root(left)(right)
	return fmt.Sprintf("%d", root.Val) + "(" + leftStr + ")" + "(" + rightStr + ")"
}
