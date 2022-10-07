package main

import (
	"fmt"
)

func main() {
	flowers := []int{1, 2, 3, 2}
	cnt := 1
	r := beautifulBouquet(flowers, cnt)
	fmt.Println(r)
}
