package offer

// 剑指 Offer 32 - III. 从上到下打印二叉树 III
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	answer := make([][]int, 0)

	// true: towards the right
	// false: towards the left
	flag := true

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		rowSize := len(queue)
		row := make([]int, 0)
		for i := 0; i < rowSize; i++ {
			curr := queue[0]
			queue = queue[1:]

			if flag {
				row = append(row, curr.Val)
			} else {
				row = append([]int{curr.Val}, row...)
			}

			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
		// 切换方向
		flag = !flag
		answer = append(answer, row)
	}
	return answer
}
