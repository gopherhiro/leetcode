package main

import (
	"fmt"

	"leetcode/sort"
)

func main() {
	nums := []int{11, 8, 3, 9, 7, 1, 2, 5}
	r := sort.FindMedian(nums)
	fmt.Println(r)
}
