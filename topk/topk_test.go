package topk

import (
	"fmt"
	"testing"
)

func TestKthLargest(t *testing.T) {
	nums := []int{4, 5, 8, 2}
	k := 3
	obj := Constructor(k, nums)
	fmt.Println(obj.Add(3))
	fmt.Println(obj.Add(5))
	fmt.Println(obj.Add(10))
	fmt.Println(obj.Add(9))
}
