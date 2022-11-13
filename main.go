package main

import (
	"fmt"
)

func main() {
	d := []string{"cat", "bat", "rat"}
	s := "the cattle was rattled by the battery"
	r := replaceWords(d, s)
	fmt.Println(r)
}
