package offer

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 剑指 Offer 27. 二叉树的镜像
// 226. Invert Binary TreeNode
// 思路：遍历二叉树的所有节点，并且对每个节点实施「交换子节点」的操作。
// time O(n): n 为二叉树节点的数目。
// space O(n): 使用的空间由递归栈的深度决定。
func mirrorTree(root *TreeNode) *TreeNode {
	// base case
	if root == nil {
		return nil
	}

	// exchange left,right child node of the root
	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp

	// continue exchange left,right node
	mirrorTree(root.Left)
	mirrorTree(root.Right)

	return root
}
