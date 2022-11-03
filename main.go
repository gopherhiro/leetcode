package main

import (
	"fmt"
)

func main() {
	grid := [][]int{
		{1, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 2, -1},
	}
	r := uniquePathsIII(grid)
	fmt.Println(r)

}
