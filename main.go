package main

import (
	"fmt"
)

func main() {
	p := "Bob"
	ban := []string{}
	r := mostCommonWord(p, ban)
	fmt.Println(r)
}
