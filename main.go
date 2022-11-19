package main

import (
	"fmt"
)

func main() {
	s := "ABCCABC"
	k := 2
	r := characterReplacement(s, k)
	fmt.Println("answer:", r)
}
