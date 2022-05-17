package DFS

// 437. Path Sum III
// 437. 路径总和 III
// 思路：DFS：穷举所有的可能。
// 访问每一个节点 node，检测以 node 为起始节点且向下延深的路径有多少种。
// time O(N^2)，space O(N^2)
func pathSumIII(root *TreeNode, targetSum int) (ans int) {
	if root == nil {
		return
	}
	ans = rootSum(root, targetSum)
	ans += pathSumIII(root.Left, targetSum)
	ans += pathSumIII(root.Right, targetSum)
	return
}

func rootSum(root *TreeNode, targetSum int) (res int) {
	if root == nil {
		return
	}
	if root.Val == targetSum {
		res++
	}
	res += rootSum(root.Left, targetSum-root.Val)
	res += rootSum(root.Right, targetSum-root.Val)
	return
}
