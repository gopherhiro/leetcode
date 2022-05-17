package tree

// 102. Binary TreeNode Level Order Traversal
// time O(N)
// space O(N)
func levelOrderTraverse(t *TreeNode) [][]int {
	if t == nil {
		return [][]int{}
	}

	res := make([][]int, 0)
	queue := make([]*TreeNode, 0)

	// 根节点入队
	queue = append(queue, t)

	// 第一层循环控制：从上向下层层遍历
	for len(queue) > 0 {
		// 记录这一层的节点值
		row := make([]int, 0)
		rowSize := len(queue)

		// 第二层循环控制每一层：从左向右遍历
		for i := 0; i < rowSize; i++ {
			// 从队列中弹出一个节点处理
			cur := queue[0]
			queue = queue[1:]
			row = append(row, cur.Val)

			// 其有左右子节点，则左右子节点入队
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}

			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}

		res = append(res, row)
	}

	return res
}
