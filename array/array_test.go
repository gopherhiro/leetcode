package array

import (
	"fmt"
	"testing"
)

func TestKClosestQuickSelect(t *testing.T) {
	points := [][]int{
		{3, 3},
		{5, -1},
		{-2, 4},
	}
	k := 1
	r := kClosestQuickSelect(points, k)
	fmt.Println(r)
}

func TestKClosest(t *testing.T) {
	points := [][]int{
		{3, 3},
		{5, -1},
		{-2, 4},
	}
	k := 1
	r := kClosest(points, k)
	fmt.Println(r)
}

func TestFindKthLargest(t *testing.T) {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	fmt.Println("before:", nums)
	r := findKthLargest(nums, k)
	fmt.Println("after:", nums)
	fmt.Println(r)
}

func TestFindKthSmallest(t *testing.T) {
	nums := []int{3, 2, 4, 5, 1}
	k := 3
	fmt.Println("before:", nums)
	r := findKthSmallest(nums, k)
	fmt.Println("after:", nums)
	fmt.Println(r)
}

func TestMissingNumber(t *testing.T) {
	n := []int{1, 3, 2, 5, 0}
	r1 := missingNumber(n)
	fmt.Println(r1)

	r2 := missingNumberB(n)
	fmt.Println(r2)
}

func TestTopKFrequent(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	r := topKFrequent(nums, k)
	fmt.Println(r)
}
