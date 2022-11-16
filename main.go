package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 0, 1}
	r := longestConsecutive(s)
	fmt.Println("answer:", r)
}
