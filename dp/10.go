package dp

import "fmt"

// 10. Regular Expression Matching
// 10. 正则表达式匹配
// 思路：dynamic programming
// dp 数组定义：
// 1、p[0..j] 可以匹配 s[0..i]，则 dp[i,j] = true
// 2、p[0..j] 不能匹配 s[0..i]，则 dp[i,j] = false
// time O(MN) , space O(MN)
func isMatch(s string, p string) bool {
	ms, np := len(s), len(p)
	if ms == 0 && np == 0 {
		return true
	}

	dp := genBoolDp(ms, np)

	// base case
	// 两者都为空，肯定是匹配成功的。
	dp[0][0] = true

	// p 为空串，s 不为空，则不匹配。
	/*	for i := 1; i < ms; i++ {
		dp[i][0] = false
	}*/

	// s 为空串，p不为空。p 有可以匹配空串的情况。
	// 则处理 a* or a*b* or a*b*c* 可以匹配空串情况
	for j := 1; j < np; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	for i := 0; i < ms; i++ {
		for j := 0; j < np; j++ {
			// 匹配
			if s[i] == p[j] || p[j] == '.' {
				dp[i+1][j+1] = dp[i][j]
			}
			if p[j] == '*' {
				// * 跟着他前一个字符走。
				// 前一个能匹配上 s[i]，* 才能有用，可以匹配 0次、一次、多次。
				// 前一个都不能匹配上 s[i]，* 也无能为力，只能让前一个字符消失，也就是匹配 0 次。
				if s[i] == p[j-1] || p[j-1] == '.' {
					zero := dp[i+1][j-1] // 匹配 0 次
					one := dp[i+1][j]    // 匹配 1 次
					many := dp[i][j+1]   // 匹配 多次
					dp[i+1][j+1] = zero || one || many
				} else {
					zero := dp[i+1][j-1] // 匹配 0 次
					dp[i+1][j+1] = zero
				}
			}
		}
	}
	return dp[ms][np]
}

func genBoolDp(m, n int) [][]bool {
	dp := make([][]bool, m+1)
	for i, _ := range dp {
		dp[i] = make([]bool, n+1)
	}
	return dp
}

// 10. Regular Expression Matching
// 10. 正则表达式匹配
// 思路：Recursion + memo
// dp 函数定义：
// 1、p[0..j] 可以匹配 s[0..i]，则 dp(i,j) = true
// 2、p[0..j] 不能匹配 s[0..i]，则 dp(i,j) = false
// time O(MN) , space O(MN)
func isMatchM(s string, p string) bool {
	ms, np := len(s), len(p)
	if ms == 0 && np == 0 {
		return true
	}
	if ms == 0 || np == 0 {
		return false
	}

	// 备忘录
	memo := make(map[string]bool, 0)

	var dp func(i, j int) bool
	dp = func(i, j int) bool {
		// base case
		// p,s 都恰好走到末尾，则匹配成功
		if j == np {
			return i == ms
		}
		// s 走到末尾，p 还未结束。
		// 则 p[j..] 可以匹配空串，则匹配成功。
		if i == ms {
			// 能匹配空串：字符和"*"一起出现
			// 不是成对出现，则匹配失败。
			if (np-j)%2 == 1 {
				return false
			}
			for ; j+1 < np; j += 2 {
				if p[j+1] != '*' {
					return false
				}
			}
			return true
		}
		key := fmt.Sprintf("%d", i) + "," + fmt.Sprintf("%d", j)

		// 查备忘录
		if val, ok := memo[key]; ok {
			return val
		}

		var ret bool

		// s[i] == p[j]
		if s[i] == p[j] || p[j] == '.' {
			// 匹配 0次 或 多次
			if j < np-1 && p[j+1] == '*' {
				zero := dp(i, j+2) // 匹配 0次
				many := dp(i+1, j) // 匹配 多次
				ret = zero || many // 只要其中一个匹配，则是匹配成功。
			} else {
				// 默认匹配 1次
				ret = dp(i+1, j+1)
			}

		} else {
			// s[i] != p[j]
			// 后面没有通配符，则匹配 0 次。
			if j < np-1 && p[j+1] == '*' {
				// 匹配 0 次
				ret = dp(i, j+2)
			} else {
				// 默认后面没有通配符，则不匹配。
				ret = false
			}
		}
		// 结果存储到备忘录
		memo[key] = ret
		return ret
	}
	return dp(0, 0)
}

// 10. Regular Expression Matching
// 10. 正则表达式匹配
// 思路：Recursion
// dp 函数定义：
// 1、p[0..j] 可以匹配 s[0..i]，则 dp(i,j) = true
// 2、p[0..j] 不能匹配 s[0..i]，则 dp(i,j) = false
// time O(2^MN) , space O(2^MN)
func isMatchR(s string, p string) bool {
	ms, np := len(s), len(p)
	if ms == 0 && np == 0 {
		return true
	}
	if ms == 0 || np == 0 {
		return false
	}

	var dp func(i, j int) bool
	dp = func(i, j int) bool {
		// base case
		// p,s 都恰好走到末尾，则匹配成功
		if j == np {
			return i == ms
		}
		// s 走到末尾，p 还未结束。
		// 则 p[j..] 可以匹配空串，则匹配成功。
		if i == ms {
			// 能匹配空串：字符和"*"一起出现
			// 不是成对出现，则匹配失败。
			if (np-j)%2 == 1 {
				return false
			}
			for ; j+1 < np; j += 2 {
				if p[j+1] != '*' {
					return false
				}
			}
			return true
		}
		// s[i] == p[j]
		if s[i] == p[j] || p[j] == '.' {
			// 匹配 0次 或 多次
			if j < np-1 && p[j+1] == '*' {
				zero := dp(i, j+2)  // 匹配 0次
				many := dp(i+1, j)  // 匹配 多次
				return zero || many // 只要其中一个匹配，则是匹配成功。
			}
			// 匹配 1次
			return dp(i+1, j+1)
		}
		// s[i] != p[j]
		if j < np-1 && p[j+1] == '*' {
			// 匹配 0 次
			return dp(i, j+2)
		}
		// s[i] != p[j] 并且后面也没有通配符，则不匹配。
		return false
	}

	return dp(0, 0)
}

func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	quickSortHelper(arr, 0, len(arr)-1)
}

func quickSortHelper(arr []int, low, high int) {
	if low < high {
		p := partition(arr, low, high)
		quickSortHelper(arr, low, p-1)
		quickSortHelper(arr, p+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func heapSort(arr []int) {
	n := len(arr)

	// 构建最大堆
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// 一个个从堆顶取出元素，并调整堆
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0] // 交换
		heapify(arr, i, 0)              // 调整堆
	}
}

func heapify(arr []int, n int, i int) {
	largest := i     // 初始化最大值为根节点
	left := 2*i + 1  // 左子节点
	right := 2*i + 2 // 右子节点

	// 如果左子节点大于根节点
	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	// 如果右子节点大于目前最大值
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	// 如果最大值不是根节点
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i] // 交换

		// 递归地调整受影响的子树
		heapify(arr, n, largest)
	}
}
