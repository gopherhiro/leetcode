package array

import (
	"fmt"
	"testing"
)

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
