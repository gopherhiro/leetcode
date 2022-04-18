package main

import (
	"fmt"
	"strings"
)

// 752. Open the Lock
// 思路：BFS
func openLock(deadends []string, target string) int {
	if len(deadends) == 0 || len(target) == 0 {
		return -1
	}
	// 把deadends存入hashtable中，方便后续查找。
	dead := make(map[string]bool, len(deadends))
	for _, v := range deadends {
		dead[v] = true
	}

	// 记录已访问过的密码，防止走回头路。
	visited := make(map[string]bool)
	queue := make([]string, 0)

	step := 0
	start := "0000"

	queue = append(queue, start)
	visited[start] = true

	for len(queue) > 0 {
		rowSize := len(queue)
		for i := 0; i < rowSize; i++ {
			current := queue[0]
			queue = queue[1:]

			// 判断密码是否合法，不合法跳过
			if _, ok := dead[current]; ok {
				continue
			}
			// 判断是否找到目标密码，是的话，则返回步数。
			if current == target {
				return step
			}

			// 将节点未遍历的相邻节点加入队列。
			for j := 0; j < 4; j++ {
				up := upOne(current, j)
				if _, ok := visited[up]; !ok {
					queue = append(queue, up)
					visited[up] = true
				}

				down := downOne(current, j)
				if _, ok := visited[down]; !ok {
					queue = append(queue, down)
					visited[down] = true
				}
			}
		}
		step++
	}
	// 如果穷举完所有可能的密码，都没有找到目标密码，则无法解锁。
	return -1
}

func upOne(s string, j int) string {
	runes := []rune(s)
	if s[j] == '9' {
		runes[j] = 48 // 48是'0'的ASCII值
	} else {
		runes[j]++
	}
	return string(runes)
}

func downOne(s string, j int) string {
	runes := []rune(s)
	if s[j] == '0' {
		runes[j] = 57 // 57是'9'的ASCII值
	} else {
		runes[j]--
	}
	return string(runes)
}


// 112. Path Sum
// 112. 路径总和
// 思路：DFS 遍历二叉树
// 1、遍历每一个节点，使用targetSum减去节点值
// 2、判断targetSum是否为0，并且是否达到叶子节点
// time O(n) space O(n)
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	ans := false

	var dfs func(root *TreeNode, remain int)
	dfs = func(root *TreeNode, remain int) {
		if root == nil {
			return
		}
		remain -= root.Val
		if remain == 0 && root.Left == nil && root.Right == nil {
			ans = true
		}
		dfs(root.Left, remain)
		dfs(root.Right, remain)
	}

	dfs(root, targetSum)
	return ans
}

// 思路：DFS
// 遍历二叉树，维护一个当前节点和，与targetSum比较即可。
func hasPathSumM(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	found := false
	curSum := 0

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		// 前序位置
		curSum += root.Val
		if root.Left == nil && root.Right == nil {
			if curSum == targetSum {
				found = true
			}
		}

		dfs(root.Left)
		dfs(root.Right)

		// 后序位置
		curSum -= root.Val
	}

	dfs(root)
	return found
}

func hasPathSumN(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}
	targetSum -= root.Val
	return hasPathSumN(root.Left, targetSum) || hasPathSumN(root.Right, targetSum)
}

// 思路：BFS
// 记录从根节点到当前节点的路径和。
// 使用两个队列，
// 一个存储将要遍历的节点
// 一个存储根节点到当前节点的路径和
// time O(n) space O(n)
func hasPathSumBFS(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	qNode := []*TreeNode{root} // 存储将要遍历的节点
	qSum := []int{root.Val}    // 存储根节点到当前节点的路径和

	for len(qNode) > 0 {
		cur := qNode[0]
		qNode = qNode[1:]

		curSum := qSum[0]
		qSum = qSum[1:]

		if cur.Left == nil && cur.Right == nil {
			if curSum == targetSum {
				return true
			}
		}

		if cur.Left != nil {
			qNode = append(qNode, cur.Left)
			qSum = append(qSum, curSum+cur.Left.Val)

		}
		if cur.Right != nil {
			qNode = append(qNode, cur.Right)
			qSum = append(qSum, curSum+cur.Right.Val)
		}

	}

	return false
}

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

// 257. Binary Tree Paths
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

// 107. Binary Tree Level Order Traversal II
// 107. 二叉树的层序遍历 II
// 思路：层序遍历框架
// time O(N), space O(M)
func levelOrderBottom(root *TreeNode) (ans [][]int) {
	if root == nil {
		return
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		level := make([]int, 0)
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			cur := queue[0]
			queue = queue[1:]

			level = append(level, cur.Val)

			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}

			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		// 把遍历得到的最新一层放到头部，即是自底向上的遍历。
		ans = append([][]int{level}, ans...)
	}
	return
}

// 107. Binary Tree Level Order Traversal II
// 107. 二叉树的层序遍历 II
// 思路：层序遍历框架
// time O(N) space O(M)
func levelOrderBottomN(root *TreeNode) (ans [][]int) {
	if root == nil {
		return
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		level := make([]int, 0)
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			cur := queue[0]
			queue = queue[1:]

			level = append(level, cur.Val)

			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}

			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		ans = append(ans, level)
	}

	// 颠倒一下元素顺序
	for i := 0; i < len(ans)/2; i++ {
		fmt.Println(ans[i], ans[len(ans)-1-i])
		ans[i], ans[len(ans)-1-i] = ans[len(ans)-1-i], ans[i]
	}
	return
}

// 117. Populating Next Right Pointers in Each Node II
// 117. 填充每个节点的下一个右侧节点指针 II
// 思路：BFS-层序遍历
// time O(N) space O(M)
func connectII(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		var prev *TreeNode
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			cur := queue[0]
			queue = queue[1:]

			if prev != nil {
				prev.Next = cur
			}
			prev = cur

			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
	}

	return root
}

// 637. Average of Levels in Binary Tree
// 637. 二叉树的层平均值
// 思路：BFS-层序遍历
// time O(N) space O(N)
func averageOfLevels(root *TreeNode) (ans []float64) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		sum := float64(0)
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			cur := queue[0]
			queue = queue[1:]

			sum += float64(cur.Val)

			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		r := sum / float64(levelSize)
		ans = append(ans, r)

	}
	return
}

// 199. Binary Tree Right Side View
// 199. 二叉树的右视图
// 思路：BFS-层序遍历，收集每层最后一个节点即可。
//  time O(N) space O(M)
func rightSideView(root *TreeNode) (ans []int) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		// 收集每一层最后一个元素
		ans = append(ans, queue[levelSize-1].Val)
		for i := 0; i < levelSize; i++ {
			cur := queue[0]
			queue = queue[1:]

			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}

	}
	return
}

// 404. Sum of Left Leaves
// 404. 左叶子之和
// 思路：DFS-前序遍历。遍历二叉树，累计左子节点之和即可。
// 左子节点：一个节点是左节点，并且是叶子节点。
// time O(N) space O(N)
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 判断是否是叶子节点
	isLeaf := func(root *TreeNode) bool {
		return root.Left == nil && root.Right == nil
	}

	sum := 0
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		// 是左叶子节点
		if root.Left != nil && isLeaf(root.Left) {
			sum += root.Left.Val
		}
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)

	return sum
}

func sumOfLeftLeavesM(root *TreeNode) int {
	if root == nil {
		return 0
	}

	sum := 0
	var dfs func(root *TreeNode, direction int)
	dfs = func(root *TreeNode, direction int) {
		if root == nil {
			return
		}
		// 是左叶子节点
		if root.Left == nil && root.Right == nil && direction == 0 {
			sum += root.Val
		}
		dfs(root.Left, 0)  // 0 means left direction
		dfs(root.Right, 1) // 1 means right direction
	}
	dfs(root, -1)

	return sum
}

// 404. Sum of Left Leaves
// 404. 左叶子之和
// 思路：BFS-层序遍历。遍历二叉树，累计左子节点之和即可。
// 左子节点：一个节点是左节点，并且是叶子节点。
// time O(N) space O(N)
func sumOfLeftLeavesBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 判断是否是叶子节点
	isLeaf := func(root *TreeNode) bool {
		return root.Left == nil && root.Right == nil
	}

	sum := 0

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur.Left != nil && isLeaf(cur.Left) {
			sum += cur.Left.Val
		}

		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}

	}
	return sum
}

// 671. Second Minimum Node In a Binary Tree
// 671. 二叉树中第二小的节点
// 思路：DFS-有条件的前序遍历
// time O(N) space O(N)
func findSecondMinimumValue(root *TreeNode) int {
	// base case
	// 00
	if root.Left == nil && root.Right == nil {
		return -1
	}
	// 根据题目要求，01、10 case 不会存在。
	// 但为了运行，需要做一下处理。
	// 01
	if root.Left == nil && root.Right != nil {
		return root.Right.Val
	}
	// 10
	if root.Left != nil && root.Right == nil {
		return root.Left.Val
	}

	// 11
	left, right := root.Left.Val, root.Right.Val

	// 如果左右子节点都等于 root.val，则去左右子树递归寻找第二小的值
	if root.Val == left {
		left = findSecondMinimumValue(root.Left)
	}
	if root.Val == right {
		right = findSecondMinimumValue(root.Right)
	}

	if left == -1 {
		return right
	}
	if right == -1 {
		return left
	}

	// 如果左右子树都找到一个第二小的值，更小的那个是整棵树的第二小元素
	return min(left, right)
}

func main() {
	path := []int{1, 2, 3}
	ans := make([][]int, 0)
	tmp := append([]int(nil), path...)
	fmt.Println(tmp)
	ans = append(ans, tmp)
	fmt.Println(ans)
}
