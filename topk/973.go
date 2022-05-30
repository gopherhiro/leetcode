package topk

import (
	"math/rand"
	"sort"
)

// 973. K Closest Points to Origin
// 973. 最接近原点的 K 个点
// 策略：排序后，取前 k 个点。
// time O(NlogN) space O(logN)
func kClosest(points [][]int, k int) [][]int {
	sort.Slice(points, func(i, j int) bool {
		p, q := points[i], points[j]
		return p[0]*p[0]+p[1]*p[1] < q[0]*q[0]+q[1]*q[1]
	})
	return points[:k]
}

// 973. K Closest Points to Origin
// 973. 最接近原点的 K 个点
// 策略：基于快排的选择算法
// time O(NlogN) space O(1)
func kClosestQuickSelect(points [][]int, k int) [][]int {
	// 为了避免出现耗时的极端情况，先随机打乱
	Shuffle(points)

	lo, hi := 0, len(points)-1
	for lo <= hi {
		p := Partition(points, lo, hi)
		if p < k {
			// 第 k 大元素在 points[p+1..hi] 中
			lo = p + 1
		} else if p > k {
			// 第 k 大元素在 points[lo..p-1] 中
			hi = p - 1
		} else {
			// p == k，找到第 k 元素。
			// 前 k 个元素点，已选出，跳出循环。
			break
		}
	}
	// 前 k 个小的点，已排好序。
	// 顺序取出即是结果。
	return points[:k]
}

// 分区的过程，就是排序的过程。
func Partition(points [][]int, lo, hi int) int {
	// 选择最后一个元素作为分区点
	pivot := points[hi]
	// i 是分区点 pivot 的下标
	// 已处理区间:[lo, i-1] 都是小于 pivot 的
	// 未处理区间:[i+1, hi-1] 都是大于 pivot 的
	i := lo
	for j := lo; j < hi; j++ {
		if Less(points[j], pivot) {
			Swap(points, i, j)
			i++
		}
	}
	// 将 pivot 放到对应的位置
	// 即 pivot 左边元素较小，右边元素较大
	Swap(points, i, hi)

	// 返回分区点下标
	return i
}

func Less(p, q []int) bool {
	return p[0]*p[0]+p[1]*p[1] < q[0]*q[0]+q[1]*q[1]
}

func Swap(points [][]int, i, j int) {
	points[i], points[j] = points[j], points[i]
}

func Shuffle(points [][]int) {
	rand.Shuffle(len(points), func(i, j int) {
		points[i], points[j] = points[j], points[i]
	})
}
