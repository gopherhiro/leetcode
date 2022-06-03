package DFS

// 129. Sum Root to Leaf Numbers
// 129. 求根节点到叶节点数字之和
// 思路：DFS + 回溯。
// time O(N), space O(N)
func sumNumbers(root *TreeNode) (ans int) {
	if root == nil {
		return
	}

	path := make([]int, 0)

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		// 为了维护经过的路径，在进入递归时，向 path 列表添加节点。
		// 在结束递归时，删除节点(即回溯到上一个节点)。
		path = append(path, root.Val)
		defer func() {
			path = path[:len(path)-1]
		}()

		if root.Left == nil && root.Right == nil {
			ans += encode(path)
			return
		}
		dfs(root.Left)
		dfs(root.Right)
	}

	dfs(root)
	return
}

func encode(p []int) int {
	if len(p) == 0 {
		return 0
	}
	sum, n := 0, 1
	for i := len(p) - 1; i >= 0; i-- {
		sum += p[i] * n
		n *= 10
	}
	return sum
}

// 129. Sum Root to Leaf Numbers
// 129. 求根节点到叶节点数字之和
// 思路：DFS + 回溯。
// time O(N), space O(N)
func sumPathNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	totalSum, pathSum := 0, 0
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		// 计算路径和
		pathSum = pathSum*10 + root.Val
		// 回溯路径和
		defer func() {
			pathSum = (pathSum - root.Val) / 10
		}()
		// 达到叶子结点
		if root.Left == nil && root.Right == nil {
			totalSum += pathSum
			return
		}
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return totalSum
}

// 129. Sum Root to Leaf Numbers
// 129. 求根节点到叶节点数字之和
// 思路：DFS，分解问题。
// time O(N), space O(N)
func sumNumbersII(root *TreeNode) (ans int) {
	if root == nil {
		return
	}
	var dfs func(root *TreeNode, prevSum int) int
	dfs = func(root *TreeNode, pathSum int) int {
		// base case
		if root == nil {
			return 0
		}

		// 前序位置
		pathSum = pathSum*10 + root.Val
		if root.Left == nil && root.Right == nil {
			return pathSum
		}

		left := dfs(root.Left, pathSum)
		right := dfs(root.Right, pathSum)

		// 后序位置
		return left + right
	}
	ans = dfs(root, 0)
	return
}

// 129. Sum Root to Leaf Numbers
// 129. 求根节点到叶节点数字之和
// 思路：DFS，分解问题。
// time O(N), space O(N)
func sumPathNumber(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var dfs func(root *TreeNode, pathSum int) int
	dfs = func(root *TreeNode, pathSum int) int {
		if root == nil {
			return 0
		}
		pathSum = 10*pathSum + root.Val
		if root.Left == nil && root.Right == nil {
			return pathSum
		}
		return dfs(root.Left, pathSum) + dfs(root.Right, pathSum)
	}
	return dfs(root, 0)
}
