package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 1, 2, 4}
	r := sumSubarrayMins(nums)
	fmt.Println(r)
}
