package main

import "fmt"

func main() {
	r := smallestEvenMultiple(6)
	fmt.Println(r)
}

func smallestEvenMultiple(n int) int {
	isOdd := n%2 != 0
	if isOdd {
		return 2 * n
	}
	return n
}
