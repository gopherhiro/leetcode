package array

import "math/rand"

/********* 如何在无限序列中随机抽取元素 ********/

// 398. Random Pick Index
// 398. 随机数索引
// 思路：水塘抽样算法
// time O(N) space O(1)
type Solution struct {
	nums []int
}

func ConstructorA(nums []int) Solution {
	return Solution{
		nums: nums,
	}
}

func (s *Solution) Pick(target int) int {
	ans, count := -1, 0
	for i, v := range s.nums {
		if v != target {
			continue
		}
		count++
		// 生成 [0,count) 之间整数
		// 该整数等于0的概率是 1/count
		if rand.Intn(count) == 0 {
			ans = i
		}
	}
	return ans
}
