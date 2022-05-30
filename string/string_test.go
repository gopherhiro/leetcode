package string

import (
	"fmt"
	"testing"
)

func TestFrequencySort(t *testing.T) {
	s := "eeeee"
	r := frequencySortN(s)
	fmt.Println(r)
}

func TestRearrangeCharacters(t *testing.T) {
	s := "aaa"
	target := "ae"
	r := rearrangeCharactersN(s, target)
	fmt.Println(r)
}
