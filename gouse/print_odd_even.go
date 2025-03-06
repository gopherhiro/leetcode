package main

import (
	"fmt"
	"sync"
)

// 题目描述：使用两个协程按顺序交替打印 1 - n 的奇偶数
// 思路：使用两个缓冲为1的 channel 来控制 goroutine 的协作。
// 为什么是用有缓冲的 channel?如果不使用有缓冲的channel会导致死锁问题。
// 原因：对于无缓冲的通道，能够向通道写入数据的前提是必须有另外一个协程在读取通道。
// 在我们的题目中：当其中一个协程打印结束后，即执行完以后，另外一个协程就不能在向channel写数据了，导致死锁。
// 比如： n = 3 时，
// 1. 打印奇数的协程先打印 1
// 2. 打印偶数的协程打印2，然后该协程执行完了
// 3. 打印奇数的协程打印3，然后再向channel写数据时，报错：fatal error: all goroutines are asleep - deadlock!
// 所以，使用有缓冲的通道可以规避这个问题。
// 如果还需要使用无缓冲的通道来实现的话，注意边界条件即可：即小于 n 的时候，才输出对应的数据即可。

func printOddEven(n int) {
	var wg sync.WaitGroup
	odd := make(chan bool, 1)
	even := make(chan bool, 1)

	// print odd number
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= n; i += 2 {
			<-odd
			fmt.Println(i)
			even <- true
		}
	}()

	// print even number
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 2; i <= n; i += 2 {
			<-even
			fmt.Println(i)
			odd <- true
		}
	}()

	odd <- true

	wg.Wait()
	close(odd)
	close(even)
}

func printOddEvenMy(n int) {
	var wg sync.WaitGroup

	odd := make(chan int, 1)
	even := make(chan int, 1)

	odd <- 0

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			i, ok := <-odd
			if !ok {
				return
			}
			fmt.Printf("%d ", i)
			if i >= n {
				close(even) // 重点（关闭另一个通道）：已经打印完毕，后续不再需要这个 chan，所以需要关闭
				return
			}
			even <- i + 1
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			i, ok := <-even
			if !ok {
				return
			}
			fmt.Printf("%d ", i)
			if i >= n {
				close(odd) // 重点（关闭另一个通道）：已经打印完毕，后续不再需要这个 chan，所以需要关闭
				return
			}
			odd <- i + 1
		}
	}()

	wg.Wait()
}

func main() {
	printOddEven(3)
}
