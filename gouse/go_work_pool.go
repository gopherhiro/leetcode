package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//题目：实现一个批处理任务，list=[1,2,3,4,5,6,7,8]，对list中的元素进行平方运算。
//1、输出结果按照list顺序输出对应结果；
//2、要求goroutine数量为3个；
//3、当运算时间超过2s还未结束，主动取消本次批处理任务；

//特性总结
//并发限制：最多3个 goroutine（通过 semaphore）。
//分批处理：任务缓冲区固定为100（jobBufferSize），适合大列表。
//超时控制：2秒超时（通过 context）。
//精确等待：使用 WaitGroup 确保任务完成。
//结果顺序：按原列表顺序输出。

// Task 定义任务结构体
type Task struct {
	index int // 改为小写
	value int // 改为小写
}

// WorkerPool 定义工作池结构
type WorkerPool struct {
	semaphore chan struct{}
	jobs      chan Task
	results   []int
	wg        sync.WaitGroup
}

// NewWorkerPool 创建新的工作池
func NewWorkerPool(maxWorkers, resultSize int) *WorkerPool {
	const jobBufferSize = 100
	return &WorkerPool{
		semaphore: make(chan struct{}, maxWorkers),
		jobs:      make(chan Task, jobBufferSize),
		results:   make([]int, resultSize),
	}
}

// SubmitTasks 分批提交任务到工作池
func (wp *WorkerPool) SubmitTasks(ctx context.Context, list []int) {
	go func() {
		defer close(wp.jobs)
		for i, num := range list {
			select {
			case wp.jobs <- Task{index: i, value: num}:
			case <-ctx.Done():
				return
			}
		}
	}()
}

// ProcessTasks 处理任务
func (wp *WorkerPool) ProcessTasks(ctx context.Context) {
	for task := range wp.jobs {
		wp.semaphore <- struct{}{}
		wp.wg.Add(1)
		go func(t Task) {
			defer func() { <-wp.semaphore }()
			defer wp.wg.Done()
			select {
			case <-ctx.Done():
				return
			default:
				wp.results[t.index] = t.value * t.value
			}
		}(task)
	}
}

// WaitForCompletion 等待任务完成或超时
func (wp *WorkerPool) WaitForCompletion(ctx context.Context) []int {
	done := make(chan struct{})
	go func() {
		wp.wg.Wait()
		close(done)
	}()

	select {
	case <-ctx.Done():
		fmt.Println("处理超时，已取消")
	case <-done:
	}
	return wp.results
}

func processBatch(list []int) []int {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	pool := NewWorkerPool(3, len(list))
	pool.SubmitTasks(ctx, list)
	pool.ProcessTasks(ctx)
	return pool.WaitForCompletion(ctx)
}

func main() {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8}
	result := processBatch(list)
	fmt.Println("处理结果:", result)
}
