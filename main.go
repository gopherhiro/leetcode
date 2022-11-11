package main

import "fmt"

func main() {
	obj := Constructor()
	obj.AddWord("bad")
	r := obj.Search("bad")
	fmt.Println("Search bad:", r)

	r = obj.Search("pad")
	fmt.Println("Search pad:", r)

	r = obj.Search(".ad")
	fmt.Println("Search: .ad", r)

	r = obj.Search("b..")
	fmt.Println("Search b..:", r)
}
