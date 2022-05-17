package tree

import "math"

// 530. Minimum Absolute Difference in BST
// DFS-inorder traverse, time O(N), space O(N)
func getMinimumDifference(root *TreeNode) int {
	answer := math.MaxInt32
	var prev *TreeNode
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		// 中序位置，进行处理
		// 如果上一个节点存在，则进行最小差值更新。
		// 否则，只对上一个节点进行记录。
		if prev != nil {
			answer = min(answer, root.Val-prev.Val)
		}
		prev = root // 记录上一个节点
		inorder(root.Right)
	}
	inorder(root)
	return answer
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
