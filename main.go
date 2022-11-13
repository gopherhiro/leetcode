package main

import (
	"fmt"
)

func main() {
	o := Constructor()
	s := []string{"hello", "leetcode"}
	w := "hello"
	o.BuildDict(s)
	r := o.Search(w)
	fmt.Println(r)
}
