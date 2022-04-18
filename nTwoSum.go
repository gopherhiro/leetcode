package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{2, 2, 2, 2, 2}
	res := fourSum(nums, 8)
	fmt.Println(res)
}

// 18. 4Sum
// 18. 四数之和
// 思路：sort + two pointer
// time O(N^2) space O(1)
func fourSum(nums []int, target int) [][]int {
	if len(nums) <= 1 {
		return [][]int{}
	}
	// sort
	sort.Ints(nums)
	return nSumTarget(nums, 4, 0, target)
}

// 15. 3Sum
// 15. 三数之和
// 思路：排序 + two pointer
// time O(N^2) space O(N)
func threeSum(nums []int) [][]int {
	if len(nums) <= 1 {
		return [][]int{}
	}
	// sort
	sort.Ints(nums)
	return nSumTarget(nums, 3, 0, 0)
}

// N数之和
// 核心思路：two pointer
// time O(N^2) space O(N)
// 前提条件：nums 数组必须是有序的
// 从指定 start 开始，是为了去掉重复组合
func nSumTarget(nums []int, n, start, target int) [][]int {
	length := len(nums)

	// 至少是2Sum && 数组大小不能小于n
	if n < 2 || length < n {
		return [][]int{}
	}

	ans := make([][]int, 0)

	// base case
	if n == 2 {
		// two Sum
		return twoSumTarget(nums, start, target)
	} else {
		// n > 2 时，递归计算 (n-1)Sum的结果
		// 穷举 nums[i]，找到合法的组合。
		for i := start; i < length; i++ {
			subs := nSumTarget(nums, n-1, i+1, target-nums[i])
			// (n-1)Sum 加上 nums[i]，即是nSum
			for _, sub := range subs {
				sub = append(sub, nums[i])
				ans = append(ans, sub)
			}
			// 去重：如果出现值相等情况，则跳过。
			for i < length-1 && nums[i] == nums[i+1] {
				i++
			}
		}

	}
	return ans
}

// 双指针
func twoSumTarget(nums []int, start, target int) (ans [][]int) {
	n := len(nums)
	lo, hi := start, n-1
	for lo < hi {
		sum := nums[lo] + nums[hi]
		// 遍历 & 去重
		if sum < target || lo > start && nums[lo] == nums[lo-1] {
			lo++
		} else if sum > target || hi < n-1 && nums[hi] == nums[hi+1] {
			hi--
		} else {
			// sum == target
			ans = append(ans, []int{nums[lo], nums[hi]})
			lo++
			hi--
		}
	}
	return
}

// N数之和
// 核心思路：two pointer
// time O(N^2) space O(N)
// 前提条件：nums 数组必须是有序的
func nSumTargetV1(nums []int, n, start, target int) [][]int {
	length := len(nums)

	// 至少是2Sum && 数组大小不能小于n
	if n < 2 || length < n {
		return [][]int{}
	}

	ans := make([][]int, 0)

	// base case
	if n == 2 { // two Sum
		// two pointer
		lo, hi := start, length-1
		for lo < hi {
			left, right := nums[lo], nums[hi]
			sum := nums[lo] + nums[hi]
			if sum < target {
				// 去重：值相同，则跳过
				for lo < hi && nums[lo] == left {
					lo++
				}
			} else if sum > target {
				// 去重：值相同，则跳过
				for lo < hi && nums[hi] == right {
					hi--
				}
			} else {
				// sum == target
				res := []int{left, right}
				ans = append(ans, res)
				// 去重：值相同，则跳过
				for lo < hi && nums[lo] == left {
					lo++
				}
				for lo < hi && nums[hi] == right {
					hi--
				}
			}
		}

	} else {
		// n > 2 时，递归计算 (n-1)Sum的结果
		// 穷举 nums[i]，找到合法的组合。
		for i := start; i < length; i++ {
			subs := nSumTarget(nums, n-1, i+1, target-nums[i])
			// (n-1)Sum 加上 nums[i]，即是nSum
			for _, sub := range subs {
				sub = append(sub, nums[i])
				ans = append(ans, sub)
			}
			// 去重：如果出现值相等情况，则跳过。
			for i < length-1 && nums[i] == nums[i+1] {
				i++
			}
		}

	}
	return ans
}

// 1. 两数之和
// 思路：hash + 遍历
// time O(N), space O(N)
func twoSum(nums []int, target int) []int {
	hash := make(map[int]int, 0)
	for k, v := range nums {
		if i, ok := hash[target-v]; ok {
			return []int{i, k}
		}
		hash[v] = k
	}
	return []int{-1, -1}
}
