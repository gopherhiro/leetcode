package main

import "fmt"

func main() {
	board := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}
	fmt.Println("b:board:", board)
	gameOfLife(board)
	fmt.Println("a:board:", board)
}
