package gouse

import (
	"fmt"
	"sync"
)

// 题目描述：给出两个列表：
// 列表1：a,b,c,d,e,f,g,h
// 列表2：1,2,3,4,5,6,7,8
// 使用两个 goroutine 交替打印数字和字母，比如：ab12cd34ef56gh78
// 思路：用两个无缓冲的 channel 来控制 goroutine 的协作。

func PrintNumberAndLetter() {
	// number 的 channel 用来通知打印数字的 goroutine.
	// letter 的 channel 用来通知打印字母的 goroutine.
	number, letter := make(chan bool), make(chan bool)
	go func() {
		i := 1
		for {
			<-number // 等待「打印数字」通知，打印数字，然后通知 letter 的 channel
			fmt.Printf("%d%d", i, i+1)
			i += 2
			letter <- true
		}
	}()

	// 加上 syn.WaitGroup 来让主协程等待打印协程的工作结束
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		i := 0
		for {
			<-letter // 等待「打印字母」通知，打印字母，然后通知 number 的 channel
			// 结束条件：到达字母字符串的末尾时，结束打印。
			if i >= len(str) {
				wg.Done()
				return
			}
			fmt.Print(str[i : i+2])
			i += 2
			number <- true
		}
	}()

	number <- true

	wg.Wait()
	fmt.Println()
	fmt.Println("Print DONE!")
}
