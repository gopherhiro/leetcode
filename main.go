package main

import (
	"fmt"
	"leetcode/interviewed"
)

func main() {
	s := "2020-05-16 19:20:34|user.login|name=Charles&location=Beijing&device=iPhone"
	r, err := interviewed.ParseLog(s)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}
}
