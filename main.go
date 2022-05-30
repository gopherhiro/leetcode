package main

import (
	"fmt"
	"strings"
)

func main() {
	r := rr('a', 5)
	fmt.Println(string(r))
}

func rr(ch rune, count int) string {
	var str strings.Builder
	i := 0
	for i < count {
		str.WriteString(string(ch))
		i++
	}
	return str.String()
}

func repeat(b []rune, count int) []rune {
	nb := make([]rune, len(b)*count)
	bp := copy(nb, b)
	for bp < len(nb) {
		copy(nb[bp:], nb[:bp])
		bp *= 2
	}
	return nb
}
