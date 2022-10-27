package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 2, 1, 5, 4}
	r := countKDifference(nums, 2)
	fmt.Println(r)
}
