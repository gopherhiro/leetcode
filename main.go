package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 0}
	sortColors(nums)
	fmt.Println(nums)
}

type ListNode struct {
	Val  int
	Next *ListNode
}
