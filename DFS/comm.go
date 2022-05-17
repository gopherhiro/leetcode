package DFS

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Left  *TreeNode
	Val   int
	Right *TreeNode
	Next  *TreeNode
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
