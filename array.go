package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{0, 1, 0, 2, 0, 3}
	ret := moveZero(arr)
	fmt.Println(ret)
	res := make([]int, 7)
	fmt.Println(len(res), cap(res))

}

// 数组是有序的，则使用双指针
func twoSumSorted(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{left + 1, right + 1}
		}
		if sum < target {
			left++
		}
		if sum > target {
			right--
		}
	}
	return []int{-1, -1}
}

// 数组是无序的：方案一，双重for循环，时间复杂度：O(n^2)，不好。
func twoSumDoubleFor(nums []int, target int) []int {
	for index, value := range nums {
		diff := target - value
		for i, v := range nums {
			if diff == v && i != index {
				return []int{index, i}
			}

		}

	}
	return []int{-1, -1}
}

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int, 0)
	for i, x := range nums {
		if p, ok := hash[target-x]; ok {
			return []int{p, i}
		}
		hash[x] = i
	}
	return []int{-1, -1}
}

func threeSum(nums []int) [][]int {
	length := len(nums)
	ans := make([][]int, 0, length)
	if length < 3 {
		return ans
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := 0; i < length; i++ {
		// 如果 nums[i]大于 0，则三数之和必然无法等于 0，结束循环
		if nums[i] > 0 {
			break
		}
		// 如果 nums[i] == nums[i-1]，则说明该数字重复，会导致结果重复，所以应该跳过
		if i > 0 && nums[i] == nums[i-1] { // de-duplication
			continue
		}
		L := i + 1
		R := length - 1

		for L < R {
			sum := nums[i] + nums[L] + nums[R]
			if sum == 0 {
				tmp := []int{nums[i], nums[L], nums[R]}
				ans = append(ans, tmp)

				// 当 sum == 0 时，nums[L] == nums[L+1] 则会导致结果重复，应该跳过，L++
				for L < R && nums[L] == nums[L+1] { // de-duplication
					L++
				}
				// 当 sum == 0 时，nums[R] == nums[R-1] 则会导致结果重复，应该跳过，R--
				for L < R && nums[R] == nums[R-1] { // de-duplication
					R--
				}

				L++
				R--
			}
			if sum < 0 {
				L++
			}
			if sum > 0 {
				R--
			}
		}

	}

	return ans
}

// 双指针解法
// 快指针fast走在前面探路
// 慢指针slow走在后面
// 找到一个不重复的元素就告诉slow并让slow前进一步，即扩充不重复元素区域
// 这样当fast指针遍历完整个数组nums后，nums[0..slow]就是不重复元素。
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	slow, fast := 0, 0
	for fast < n {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}

// 双指针解法
// 快指针fast走在前面探路
// 慢指针slow走在后面
// 找到一个非去除的元素就告诉slow并让slow前进一步，即扩充非去除元素区域
// 当fast指针遍历完整个数组nums后，nums[0..slow-1]就是非去除元素。
func removeDuplicatesByVal(nums []int, val int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	slow, fast := 0, 0
	for fast < n {
		// != val 的值，是需要扩充区域来存放的值
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

func removeDuplicatesByValAgain(nums []int, val int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	idx := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[idx] = nums[i]
			idx++
		}
	}
	return idx
}

func moveZeroes(nums []int) []int {
	idx := removeDuplicatesByVal(nums, 0)
	for idx < len(nums) {
		nums[idx] = 0
		idx++
	}
	return nums
}

func removeDuplicatesTwo(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	slow, fast := 2, 2
	for fast < n {
		if nums[fast] != nums[slow-2] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

func moveZero(nums []int) []int {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
	return nums
}
