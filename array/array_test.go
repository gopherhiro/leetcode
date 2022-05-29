package array

import (
	"fmt"
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	r := findKthLargest(nums, k)
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
