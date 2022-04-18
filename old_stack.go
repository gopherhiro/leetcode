package main

import (
	"container/list"
	"fmt"
)

func main() {
	num1 := []int{1, 2, 1}
	res := nextGreaterElementsWithList(num1)
	fmt.Println(res)
}

// 单调栈: stack + hashtable
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	hash := make(map[int]int, len(nums2))
	var stack []int
	for _, num := range nums2 {
		for len(stack) > 0 && stack[len(stack)-1] <= num {
			stackTop := stack[len(stack)-1]
			hash[stackTop] = num
			// a[low : high]
			// This selects a half-open range which includes the first element, but excludes the last one.
			stack = stack[:len(stack)-1] // remove stack top element
		}
		stack = append(stack, num)
	}

	res := make([]int, 0, len(nums1))
	for _, num := range nums1 {
		if value, ok := hash[num]; ok {
			res = append(res, value)
		} else {
			res = append(res, -1)
		}

	}

	return res
}

func nextGreaterElementAgain(nums1 []int, nums2 []int) []int {
	hash := make(map[int]int, len(nums2))
	stack := list.New()
	for _, num := range nums2 {
		for stack.Len() > 0 && stack.Back().Value.(int) <= num {
			hash[stack.Back().Value.(int)] = num
			stack.Remove(stack.Back())
		}
		stack.PushBack(num)
	}

	res := make([]int, 0, len(nums1))
	for _, num := range nums1 {
		if value, ok := hash[num]; ok {
			res = append(res, value)
		} else {
			res = append(res, -1)
		}

	}

	return res
}

func nextGreaterElementsWithList(nums []int) []int {
	n := len(nums)
	res := make([]int, len(nums))
	stack := list.New()
	for i := 2*n - 1; i >= 0; i-- {
		for stack.Len() > 0 && stack.Back().Value.(int) <= nums[i%n] {
			stack.Remove(stack.Back())
		}
		if stack.Len() == 0 {
			res[i%n] = -1
		} else {
			res[i%n] = stack.Back().Value.(int)
		}
		stack.PushBack(nums[i%n])
	}
	return res
}

func nextGreaterElementsWithSlice(nums []int) []int {
	stack := make([]int, 0)
	res := make([]int, len(nums))
	l := len(nums)
	for i := l*2 - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] <= nums[i%l] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			res[i%l] = stack[len(stack)-1]
		} else {
			res[i%l] = -1
		}
		stack = append(stack, nums[i%l])
	}
	return res
}
