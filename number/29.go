package number

// Given integer operands a and n,
// a << n; shifts all bits in a to the left n times
// a >> n; shifts all bits in a to the right n times

// 29. Divide Two Integers
// 29. 两数相除
// 思路：bit operation
// time O(logn^2) space O(1)
func divide(a int, b int) int {
	// 预先进行：溢出处理
	if a == -1<<31 && b == -1 {
		return 1<<31 - 1
	}
	if a == 1<<31 && b == 1 {
		return 1<<31 - 1
	}

	// 预先确定结果正与负
	sign := -1
	if a > 0 && b > 0 || a < 0 && b < 0 {
		sign = -1 * sign
	}

	// 开始位运算
	a, b = abs(a), abs(b)
	ans := 0
	// 在 a - b >= 0 的情况下，进行循环。
	for a-b >= 0 {
		// 计算左移次数
		// 更新迭代数据
		shift := 0 // times of left shift
		// 第一种写法：
		// b<<shift<<1, 计算步骤：((b<<shift)<<1), 这里 <<1：是为了再走一步，确保 x 能够得到正确的左移次数。
		// 也可写为：b<<(shift+1)
		for a-(b<<shift<<1) >= 0 {
			shift++
		}
		// 因为直接计算，左移次数会多一。所以，第二种写法为：在退出循环后，对统计的左移次数减一即可。
		// for a-(b<<shift) >= 0 {
		// 	shift++
		// }
		// shift = shift - 1

		ans += 1 << shift
		a -= b << shift
	}
	return sign * ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
