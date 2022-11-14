package main

import (
	"fmt"
)

func main() {
	k := []byte{'a', 'b', 'c', 'd'}
	v := []string{"ei", "zf", "ei", "am"}
	d := []string{"abcd", "acbd", "adbc", "badc", "dacb", "cadb", "cbda", "abad"}
	o := Constructor(k, v, d)
	r := o.Encrypt("abcd")
	fmt.Println("Encrypt abcd:", r)
	count := o.Decrypt("eizfeiam")
	fmt.Println("count:", count)
}
