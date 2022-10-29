package main

import (
	"fmt"
)

func main() {
	items := [][]string{
		{"phone", "blue", "pixel"}, {"computer", "silver", "lenovo"}, {"phone", "gold", "iphone"},
	}
	r := countMatches(items, "color", "silver")
	fmt.Println(r)
}
