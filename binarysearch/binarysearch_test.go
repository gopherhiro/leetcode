package binarysearch

import (
	"fmt"
	"testing"
)

func TestSearchRotate(t *testing.T) {
	num := []int{4, 5, 6, 7, 0, 1, 2}
	target := 6
	r := searchRotate(num, target)
	fmt.Println(r)
}
