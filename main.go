package main

import (
	"fmt"
	"math/rand"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	arr := []int{4, 5, 1, 6, 2, 7, 3, 8}
	r := sortArray(arr)
	fmt.Println(r)
}

func sortArray(nums []int) []int {
	n := len(nums)
	if n == 0 || n == 1 {
		return nums
	}

	var sort func(nums []int, lo, hi int)
	sort = func(nums []int, lo, hi int) {
		if lo >= hi {
			return
		}
		p := partition(nums, lo, hi)
		sort(nums, lo, p-1)
		sort(nums, p+1, hi)
	}

	shuffle(nums)

	sort(nums, 0, n-1)
	return nums
}

func partition(nums []int, lo, hi int) int {
	povit := nums[hi]

	i := lo
	for j := lo; j < hi; j++ {
		if nums[j] < povit {
			swap(nums, i, j)
			i++
		}
	}
	swap(nums, i, hi)
	return i
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func shuffle(nums []int) {
	rand.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})
}
