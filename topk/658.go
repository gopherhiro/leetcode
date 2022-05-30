package topk

// 658. Find K Closest Elements
// 658. 找到 K 个最接近的元素
// 策略：two pointer
// time O(N) space O(1)
func findClosestElements(arr []int, k, x int) []int {
	n := len(arr)
	if n == 0 {
		return []int{}
	}
	if k > n {
		return []int{}
	}

	left, right := 0, n-1
	rem := n - k // 需要移除的个数
	for rem > 0 {
		if abs(x, arr[left]) <= abs(x, arr[right]) {
			right--
		} else {
			left++
		}
		rem--
	}

	return arr[left : right+1]
}

func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}
