package tree

// 538. Convert BST to Greater TreeNode
// 538. 把二叉搜索树转换为累加树
// Approach #1 Recursion, time O(N), space O(N)
// 思路：采用中序降序遍历，可以得到该节点右边的所有节点值的和。
// 从而的到Greater Sum TreeNode
func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var sum int
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		traverse(root.Right)

		// 维护累加和 & 更新节点值
		sum += root.Val
		root.Val = sum

		traverse(root.Left)
	}
	traverse(root)

	return root
}
