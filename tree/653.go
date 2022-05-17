package tree

// 653. 两数之和 IV - 输入 BST
// 653. Two Sum IV - Input is a BST
// 思路：遍历二叉树 + 哈希表
func findTarget(root *TreeNode, k int) (answer bool) {
	if root == nil {
		answer = false
	}
	hashtable := map[int]struct{}{}
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		if _, ok := hashtable[k-root.Val]; ok {
			answer = true
			return
		}
		hashtable[root.Val] = struct{}{}
		traverse(root.Left)
		traverse(root.Right)
	}
	traverse(root)
	return
}
