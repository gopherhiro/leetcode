package tree

// 559. Maximum Depth of N-ary TreeNode
// 思路：使用N叉树的遍历框架
// 深度优先遍历（DFS）：time O(N) space O(N)
// Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

func maxDepthOfNTree(root *Node) int {
	// base case
	if root == nil {
		return 0
	}

	// 遍历N叉树，得到子树的最大深度
	var maxChildDepth int
	for _, child := range root.Children {
		childDepth := maxDepthOfNTree(child)
		if childDepth > maxChildDepth {
			maxChildDepth = childDepth
		}
	}

	// N 个子树的最大深度中的最大值为 maxChildDepth
	// 则该 N 叉树的最大深度为 maxChildDepth+1。
	return maxChildDepth + 1
}
