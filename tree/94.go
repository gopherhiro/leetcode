package tree

// 94. Binary TreeNode Inorder Traversal
// 94. 二叉树的中序遍历
// 递归方案
// time O(N)
// space O(N)
type InorderTraversal struct {
	inorderRes []int
}

func NewInorder() *InorderTraversal {
	return &InorderTraversal{inorderRes: []int{}}
}

func (it *InorderTraversal) inorderTraversal(root *TreeNode) []int {
	it.inorderTraversalHelper(root)
	return it.inorderRes
}

func (it *InorderTraversal) inorderTraversalHelper(root *TreeNode) {
	if root == nil {
		return
	}
	it.inorderTraversalHelper(root.Left)
	it.inorderRes = append(it.inorderRes, root.Val)
	it.inorderTraversalHelper(root.Right)
}

// 迭代方案：使用栈结构，模拟函数递归调用栈。
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			// 取栈顶元素并出栈
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, node.Val)
			root = node.Right
		}

	}
	return res
}

// 分解问题
// 动态规划方案：牛！
func inorderTraversN(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	left := inorderTraversN(root.Left)
	res = append(res, left...)

	res = append(res, root.Val)

	right := inorderTraversN(root.Right)
	res = append(res, right...)

	return res
}

func preorderTraversN(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	res = append(res, root.Val)

	left := preorderTraversN(root.Left)
	res = append(res, left...)

	right := preorderTraversN(root.Right)
	res = append(res, right...)

	return res
}

func postorderTraversN(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	left := postorderTraversN(root.Left)
	res = append(res, left...)

	right := postorderTraversN(root.Right)
	res = append(res, right...)

	res = append(res, root.Val)

	return res
}
