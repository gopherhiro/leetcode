package offer

// 剑指 Offer 32 - I. 从上到下打印二叉树
// 思路：利用队列先进先出特性，BFS
// time O(n) space O(n)
func levelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	answer := make([]int, 0)
	queue := []*TreeNode{root}
	// 控制从上向下一层一层遍历
	for len(queue) > 0 {
		// 控制从左向右一个一个遍历
		row := len(queue)
		for i := 0; i < row; i++ {
			curr := queue[0]
			queue = queue[1:]

			// collect answer
			answer = append(answer, curr.Val)

			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
	}
	return answer
}
