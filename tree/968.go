package tree

// 968. Binary TreeNode Cameras
// DFS post order traversal, time O(N), space O(N)
// 思路：后序位置可以接收到子树的信息，同时也可以通过函数参数接收到父节点传递的信息。
func minCameraCover(root *TreeNode) (answer int) {
	if root == nil {
		answer = 0
	}

	const NodeIsNIL = -1 // 空节点
	const (
		Uncover = iota // 监控未覆盖
		Cover          // 监控已覆盖
		Set            // 安装摄像头
	)

	var setCamera func(root *TreeNode, hasParent bool) int
	setCamera = func(root *TreeNode, hasParent bool) int {
		// base case
		if root == nil {
			return NodeIsNIL
		}
		// 获取左右子节点的情况
		left := setCamera(root.Left, true)
		right := setCamera(root.Right, true)

		// 根据左右子节点的情况和父节点的情况
		// 判断当前节点应该做的事情
		if left == NodeIsNIL && right == NodeIsNIL {
			// 当前节点是叶子节点
			if hasParent {
				// 有父节点的话，让父节点来 Cover 自己
				return Uncover
			}
			// 没有父节点的话，自己 Set （安装）一个摄像头
			answer++
			return Set
		}

		// 左右子树存在没有被 cover 的
		// 必须在当前节点 set 一个摄像头
		if left == Uncover || right == Uncover {
			answer++
			return Set
		}

		// 左右子树只要有一个 Set 了摄像头
		// 当前节点就已经是 Cover 状态了
		if left == Set || right == Set {
			return Cover
		}

		// 最后一种情况：left == Cover && right == Cover
		// 即当前节点的左右子节点都被 Cover
		if hasParent {
			// 有父节点，可以等父节点 Cover 自己
			return Uncover
		}
		// 没有父节点，只能自己 Set 一个摄像头
		answer++
		return Set
	}

	setCamera(root, false)
	return
}
