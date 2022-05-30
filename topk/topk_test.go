package topk

import (
	"fmt"
	"testing"
)

func TestFindClosestElements(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	k := 4
	x := 3
	r := findClosestElements(arr, k, x)
	fmt.Println(r)
}

func TestKthLargest(t *testing.T) {
	nums := []int{4, 5, 8, 2}
	k := 3
	obj := Constructor(k, nums)
	fmt.Println(obj.Add(3))
	fmt.Println(obj.Add(5))
	fmt.Println(obj.Add(10))
	fmt.Println(obj.Add(9))
}
