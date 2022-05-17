package tree

// 701. Insert into a Binary Search TreeNode
// Recursion: time O(H), space O(H)
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	// 找到位置，插入新节点
	if root == nil {
		return &TreeNode{Val: val}
	}
	// 如果插入值小于根节点，插入到左子树
	if val < root.Val {
		root.Left = insertIntoBST(root.Left, val)
	}
	// 如果插入值，大于根节点，插入到右子树
	if val > root.Val {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}
