package DFS

// 199. Binary Tree Right Side View
// 199. 二叉树的右视图
// 思路：BFS-层序遍历，收集每层最后一个节点即可。
//  time O(N) space O(M)
func rightSideView(root *TreeNode) (ans []int) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		// 收集每一层最后一个元素
		ans = append(ans, queue[levelSize-1].Val)
		for i := 0; i < levelSize; i++ {
			cur := queue[0]
			queue = queue[1:]

			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}

	}
	return
}
