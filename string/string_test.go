package string

import (
	"fmt"
	"testing"
)

func TestRearrangeCharacters(t *testing.T) {
	s := "aaaddaa"
	target := "a"
	r := rearrangeCharacters(s, target)
	fmt.Println(r)
}
