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
// 思路：DFS，分解问题。
// time O(N), space O(N)
func sumNumbersII(root *TreeNode) (ans int) {
	if root == nil {
		return
	}
	var dfs func(root *TreeNode, prevSum int) int
	dfs = func(root *TreeNode, prevSum int) int {
		// base case
		if root == nil {
			return 0
		}

		// 前序位置
		sum := prevSum*10 + root.Val
		if root.Left == nil && root.Right == nil {
			return sum
		}

		left := dfs(root.Left, sum)
		right := dfs(root.Right, sum)

		// 后序位置
		return left + right
	}
	ans = dfs(root, 0)
	return
}
