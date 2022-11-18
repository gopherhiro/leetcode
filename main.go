package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	k := 2
	r := longestOnes(nums, k)
	fmt.Println("answer:", r)
}
