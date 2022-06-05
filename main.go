package main

import "fmt"

func main() {
	n := []int{1, 2, 3, 4, 5, 6}
	m := []int{1, 2, 3}
	n = m
	fmt.Printf("n : %T %v\n", n, n)
	fmt.Printf("m : %T %v\n", m, m)
}
