package main

import (
	"fmt"

	"leetcode/sort"
)

func main() {
	nums := []int{1, 3, 2, 5, 6}
	r := sort.FindMedian(nums)
	fmt.Println(r)
}
