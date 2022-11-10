package main

import (
	"fmt"
)

func main() {
	b := [][]byte{
		{'C', 'A', 'A'},
		{'A', 'A', 'A'},
		{'B', 'C', 'D'},
	}
	w := "AAB"
	r := exist(b, w)
	fmt.Println(r)
}
