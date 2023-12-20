package gouse

import (
	"fmt"
	"sync"
)

func main() {
	mq := Constructor()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		mq.Producer(1)
		mq.Producer(2)
		mq.Producer(3)
		close(mq.ch)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		mq.Consumer(callback)
	}()

	wg.Wait()
	fmt.Println("Done")
}

type MQ struct {
	ch chan int
}

func Constructor() MQ {
	return MQ{
		ch: make(chan int, 1),
	}
}

func (m *MQ) Producer(elem int) {
	m.ch <- elem
}

func (m *MQ) Consumer(callback func(e int)) {
	for {
		e, ok := <-m.ch
		if !ok {
			fmt.Println("all elem has been consumed.")
			return
		}
		callback(e)
	}
}

func callback(e int) {
	fmt.Println("consumer elem:", e)
}
