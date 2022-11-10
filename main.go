package main

import "fmt"

func main() {
	b := [][]byte{
		{'o', 'a', 'b', 'n'},
		{'o', 't', 'a', 'e'},
		{'a', 'h', 'k', 'r'},
		{'a', 'f', 'l', 'v'},
	}
	w := []string{"oa", "oaa"}
	r := findWords(b, w)
	fmt.Println(r)
}
