package number

import (
	"strconv"
	"strings"
)

// 412. Fizz Buzz
// 思路：模拟
// time O(n) space O(n)
func fizzBuzz(n int) []string {
	ans := make([]string, n+1)
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			ans[i] = "FizzBuzz"
		} else if i%3 == 0 {
			ans[i] = "Fizz"
		} else if i%5 == 0 {
			ans[i] = "Buzz"
		} else {
			ans[i] = strconv.Itoa(i)
		}
	}
	return ans[1:]
}

// 412. Fizz Buzz
// 思路：模拟 + 字符串拼接
// time O(n) space O(1)
// 更推荐该方法：可读性更好。
func fizzBuzz2(n int) (ans []string) {
	for i := 1; i <= n; i++ {
		sb := &strings.Builder{}
		if i%3 == 0 {
			sb.WriteString("Fizz")
		}
		if i%5 == 0 {
			sb.WriteString("Buzz")
		}
		// 既不是 3 也不是 5 的倍数
		if sb.Len() == 0 {
			sb.WriteString(strconv.Itoa(i))
		}
		ans = append(ans, sb.String())
	}
	return
}
