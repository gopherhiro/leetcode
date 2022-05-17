package DFS

import (
	"fmt"
	"strings"
)

// 257. Binary Tree Paths
// 257. 二叉树的所有路径
// 思路：DFS，递归遍历，收集路径。
// time O(N), space O(N)
func binaryTreePaths(root *TreeNode) (ans []string) {
	if root == nil {
		return
	}

	path := make([]string, 0)

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		// 为了维护经过的路径
		// 在进入递归时，向 path 列表添加节点。
		// 在结束递归时，删除节点(即回溯到上一个节点)。
		path = append(path, fmt.Sprintf("%d", root.Val))
		defer func() {
			path = path[:len(path)-1]
		}()

		if root.Left == nil && root.Right == nil {
			r := strings.Join(path, "->")
			ans = append(ans, r)
			return
		}
		dfs(root.Left)
		dfs(root.Right)
	}

	dfs(root)
	return
}
