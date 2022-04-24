package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := [][]int{
		{3, 1, 2, 4, 5},
		{1, 2, 3, 4},
		{3, 4, 5, 6},
	}
	r := intersection(nums)
	fmt.Println(r)
}

// 2248. Intersection of Multiple Arrays
// 6041. 多个数组求交集
// 策略：哈希表
// space O(MN) space O(M+N)
func intersection(nums [][]int) []int {
	hashtable := make(map[int]int, 0)
	for _, num := range nums {
		for _, v := range num {
			hashtable[v]++ // 统计每一个元素出现的次数
		}
	}
	ans := make([]int, 0)
	n := len(nums)
	for v, count := range hashtable {
		if count == n {
			// 该元素在每一个子数组中都存在。
			ans = append(ans, v)
		}
	}
	sort.Ints(ans)
	return ans
}

// 采集果实
func getMinimumTime(time []int, fruits [][]int, limit int) int {
	if len(fruits) == 0 {
		return 0
	}
	sum := 0
	for _, fruit := range fruits {
		typ := fruit[0]
		num := fruit[1]
		count := doCount(num, limit)
		consume := time[typ]
		sum += consume * count
	}
	return sum
}

func doCount(num, limit int) int {
	if num == 0 {
		return 0
	}
	if num <= limit {
		return 1
	}
	count := 0
	for num > 0 {
		num = num - limit
		count++
	}
	return count
}

// 搭桥过河
//「可以交换河道顺序」版本
func buildBridgeN(num int, wood [][]int) int64 {
	if num == 0 || len(wood) == 0 {
		return 0
	}
	sort.Slice(wood, func(i, j int) bool {
		return wood[i][1] < wood[j][1]
	})
	nature := 0
	xEnd := wood[0][1]
	for _, w := range wood {
		start := w[0]
		if start < xEnd {
			continue
		}
		nature += start - xEnd
		xEnd = w[1]
	}
	return int64(nature)
}

// LCP 50. 宝石补给
// 思路：场景模拟
// time O(N) space O(1)
// N 是 max(len(gem),len(operations))
func giveGem(gem []int, operations [][]int) int {
	if len(gem) == 0 {
		return 0
	}
	for i := 0; i < len(operations); i++ {
		x := operations[i][0]
		y := operations[i][1]

		// gem[x] 的一半
		give := gem[x] / 2

		gem[x] -= give // 授人玫瑰，减少
		gem[y] += give // 得到馈赠，增加
	}
	// 找寻最大、最小值，求差
	mi, ma := gem[0], gem[0]
	for i := 0; i < len(gem); i++ {
		if gem[i] < mi {
			mi = gem[i]
		}
		if gem[i] > ma {
			ma = gem[i]
		}
	}
	return ma - mi
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
