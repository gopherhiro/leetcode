package main

import "fmt"

func main() {
	board := [][]int{
		{50},
	}
	t := diagonalSum(board)
	fmt.Println(t)
}
