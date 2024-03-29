package DFS

// 113. Path Sum II
// 113. 路径总和 II
// 思路：DFS 遍历二叉树
// 1、遍历每一个节点，使用targetSum减去节点值
// 2、判断remain是否为0，并且是否达到叶子节点
// 核心：为了维护经过的路径，在进入递归时，向 path 列表添加节点，在结束递归时，删除节点(即回溯到上一个节点)。
// time O(n) space O(n)
func pathSum(root *TreeNode, targetSum int) (ans [][]int) {
	if root == nil {
		return
	}

	// path 是公用数组，在 DFS 过程中需要进行回溯，以保持路径干净。
	path := make([]int, 0)

	var dfs func(root *TreeNode, remain int)
	dfs = func(root *TreeNode, remain int) {
		if root == nil {
			return
		}
		remain -= root.Val
		// 为了维护经过的路径，在进入递归时，向 path 列表添加节点。
		// 在结束递归时，删除节点(即回溯到上一个节点)。
		path = append(path, root.Val)
		defer func() {
			path = path[:len(path)-1]
		}()

		if remain == 0 && root.Left == nil && root.Right == nil {
			// 切片是底层数组的视图。
			// path的值会改变，不能直接append path
			// 需要复制出一个副本。
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}

		dfs(root.Left, remain)
		dfs(root.Right, remain)
	}

	dfs(root, targetSum)
	return
}

// 113. Path Sum II
// 113. 路径总和 II
// 思路：DFS 遍历二叉树
// 1、遍历每一个节点，使用targetSum减去节点值
// 2、判断remain是否为0，并且是否达到叶子节点
// time O(n) space O(n)
func pathSum2(root *TreeNode, targetSum int) (ans [][]int) {
	if root == nil {
		return
	}

	// path 是公用数组，在 DFS 过程中需要进行回溯，以保持路径干净。
	path := make([]int, 0)

	var dfs func(root *TreeNode, remain int, path []int)
	dfs = func(root *TreeNode, remain int, path []int) {
		if root == nil {
			return
		}

		remain -= root.Val
		path = append(path, root.Val)
		if remain == 0 && root.Left == nil && root.Right == nil {
			// 切片是底层数组的视图。
			// path的值会改变，不能直接append path
			// 需要复制出一个副本。
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}

		dfs(root.Left, remain, path)
		dfs(root.Right, remain, path)
	}

	dfs(root, targetSum, path)
	return
}
