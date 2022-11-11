package main

import "fmt"

func main() {
	words := []string{"appe", "apple"}
	o := Constructor(words)
	r := o.F("a", "ep")
	fmt.Println(r)
}
