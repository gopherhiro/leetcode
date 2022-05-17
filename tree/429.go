package tree

// 429. N-ary TreeNode Level Order Traversal
// 思路：参考二叉树层序遍历。
// 使用队列：time O(N), space O(N)
func levelOrderNTree(root *Node) (answer [][]int) {
	if root == nil {
		return
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		row := make([]int, 0)
		rowSize := len(queue)
		for i := 0; i < rowSize; i++ {
			cur := queue[0]
			queue = queue[1:]
			row = append(row, cur.Val)
			if cur.Children != nil {
				queue = append(queue, cur.Children...)
			}
		}
		answer = append(answer, row)
	}
	return
}
