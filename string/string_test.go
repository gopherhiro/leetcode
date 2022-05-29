package string

import (
	"fmt"
	"testing"
)

func TestRearrangeCharacters(t *testing.T) {
	s := "aaa"
	target := "ae"
	r := rearrangeCharactersN(s, target)
	fmt.Println(r)
}
