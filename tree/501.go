package tree

// 501. Find Mode in Binary Search TreeNode
// 501. 二叉搜索树中的众数
// dfs inorder traverse + 策略逻辑
// time O(N), space O(N)
// 如果看不懂，那就模拟一下。
func findMode(root *TreeNode) []int {
	var base, count, maxCount int
	answer := make([]int, 0)

	// 一棵二叉搜索树(BST, binary search tree)的中序遍历序列是一个递增的有序序列。
	// 顺序扫描中序遍历序列：
	// 用 base 记录当前的数字，
	// 用 count 记录当前数字重复的次数，
	// 用 maxCount 来维护已经扫描过的数当中出现最多的那个数字的出现次数，
	// 用 answer 数组记录出现的众数。
	update := func(value int) {
		if value == base {
			count++
		} else {
			base, count = value, 1
		}

		if count == maxCount {
			answer = append(answer, base)
		} else if count > maxCount {
			maxCount = count
			answer = []int{base}
		}
	}

	// 考虑不存储中序遍历序列：
	// 如果在递归中序遍历的过程中访问节点，
	// 就可以省去存储中序遍历序列的空间。
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		update(root.Val)
		dfs(root.Right)
	}
	dfs(root)
	return answer
}
