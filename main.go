package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	r := spiralOrder(matrix)
	fmt.Println(r)
}
