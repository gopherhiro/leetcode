package sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	nums := []int{5, 1, 1, 2, 0, 0}
	fmt.Println(nums)
	shuffle(nums)
	fmt.Println(nums)
	r := quickSort(nums)
	fmt.Println(r)
}

func TestMergeSort(t *testing.T) {
	nums := []int{5, 4, 3, 2, 1}
	fmt.Println(nums)
	r := mergeSort(nums)
	fmt.Println(r)
}
