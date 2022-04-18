package main

import (
	"container/list"
	"fmt"
)

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	res := maxSlidingWindowAgain(nums, k)
	fmt.Println(res)
}

func maxSlidingWindowAgain(nums []int, k int) []int {
	queue := make([]int, 0)
	res := make([]int, 0)
	for i, val := range nums {
		//step 1: if the element in the queue is less than the curr value then remove it
		//coz we only want to have max elements and 0th position will always have the index of max element
		for len(queue) > 0 && nums[queue[len(queue)-1]] <= val {
			//remove this value
			queue = queue[:len(queue)-1]
		}
		//append the current value
		queue = append(queue, i)

		//if the window is full then add the element at 0th pos to the res
		if i+1 >= k {
			//append the value to res
			res = append(res, nums[queue[0]])
		}
		//if the element at 0th index is less the sliding window index range then remove it
		if queue[0] <= i-k+1 {
			queue = queue[1:]
		}
	}
	return res
}

func maxSlidingWindow(nums []int, k int) []int {
	window := list.New()
	res := make([]int, 0, len(nums))

	for i := 0; i < len(nums); i++ {
		// 先填满窗口的 k - 1
		if i < k-1 {
			push(window, nums[i])
		} else {
			// 窗口向前滑动
			// 加入新数字
			push(window, nums[i])
			// 记录当前窗口的最大值
			res = append(res, max(window))
			// 移出旧数字
			pop(window, nums[i-k+1])
		}
	}
	return res
}

// 单调队列（单调递减）
func push(w *list.List, val int) {
	for w.Len() > 0 && w.Back().Value.(int) < val {
		w.Remove(w.Back())
	}
	w.PushBack(val)
}

func pop(w *list.List, val int) {
	// 队头元素val判等，是可能已经被删除了，即不用再删除。
	// 如果没有删除，操作删除即可。
	if w.Front().Value.(int) == val {
		w.Remove(w.Front())
	}
}

func max(w *list.List) int {
	// 队头元素即是最大的值
	return w.Front().Value.(int)
}
