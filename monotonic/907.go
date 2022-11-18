package monotonic

// 907. Sum of Subarray Minimums
// 907. 子数组的最小值之和
// 思路：单调栈
// time O(n) time O(n)
func sumSubarrayMins(arr []int) int {
	const MOD int = 1e9 + 7
	stack := make([]int, 0)
	sum := 0
	for i := 0; i <= len(arr); i++ {
		// when i reaches the array length, it is an indication that
		// all the elements have been processed, and the remaining
		// elements in the stack should now be popped out.
		for len(stack) != 0 && (i == len(arr) || arr[stack[len(stack)-1]] >= arr[i]) {
			// Notice the sign ">=", This ensures that no contribution
			// is counted twice. rightBoundary takes equal or smaller
			// elements into account while leftBoundary takes only the
			// strictly smaller elements into account
			mid := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// left boundary and right boundary
			left := -1
			if len(stack) != 0 {
				left = stack[len(stack)-1]
			}
			right := i

			// count of subarrays where mid is the minimum element
			count := (mid - left) * (right - mid)
			sum += count * arr[mid]
		}
		stack = append(stack, i)
	}
	return sum % MOD
}

// 907. Sum of Subarray Minimums
// 907. 子数组的最小值之和
// 思路：暴力破解
// time O(n^2) space O(1)
// 备注：超时
func sumSubarrayMins2(arr []int) int {
	n := len(arr)
	sumMin := 0
	for start := 0; start < n; start++ {
		curMin := arr[start]
		for end := start; end < n; end++ {
			curMin = min(curMin, arr[end])
			sumMin += curMin
		}
	}
	return sumMin
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
