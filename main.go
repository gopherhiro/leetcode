package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	str := "abc"
	fmt.Println(str)
	r := getMD5String(str)
	fmt.Println(r)
}

func getMD5String(text string) string {
	h := md5.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}
