package BFS

// 112. Path Sum
// 112. 路径总和
// 思路：BFS
// 记录从根节点到当前节点的路径和。
// 使用两个队列，
// 一个存储将要遍历的节点
// 一个存储根节点到当前节点的路径和
// time O(n) space O(n)
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	qNode := []*TreeNode{root} // 存储将要遍历的节点
	qSum := []int{root.Val}    // 存储根节点到当前节点的路径和

	for len(qNode) > 0 {
		cur := qNode[0]
		qNode = qNode[1:]

		curSum := qSum[0]
		qSum = qSum[1:]

		if cur.Left == nil && cur.Right == nil {
			if curSum == targetSum {
				return true
			}
		}

		if cur.Left != nil {
			qNode = append(qNode, cur.Left)
			qSum = append(qSum, curSum+cur.Left.Val)

		}
		if cur.Right != nil {
			qNode = append(qNode, cur.Right)
			qSum = append(qSum, curSum+cur.Right.Val)
		}

	}

	return false
}
