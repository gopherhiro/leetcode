package linkedlist

import (
	"leetcode/pkg"
	"testing"
)

func TestGetIntersectionNode(t *testing.T) {
	num1 := []int{1, 2, 3, 4, 5}
	h1 := pkg.CreateListFromTail(num1)
	pkg.PrintLinkedList(h1)

	num2 := []int{5, 6, 7, 8, 9}
	h2 := pkg.CreateListFromTail(num2)
	pkg.PrintLinkedList(h2)

	h3 := getIntersectionNode(h1, h2)
	pkg.PrintLinkedList(h3)
}

func TestSwapPairs(t *testing.T) {
	num := []int{1, 2, 3, 4, 5}
	h1 := pkg.CreateListFromTail(num)
	pkg.PrintLinkedList(h1)

	h2 := swapPairsK(h1)
	pkg.PrintLinkedList(h2)
}

func TestRSwapPairsR(t *testing.T) {
	num := []int{1, 2, 3, 4, 5}
	h1 := pkg.CreateListFromTail(num)
	pkg.PrintLinkedList(h1)

	h2 := swapPairsR(h1)
	pkg.PrintLinkedList(h2)
}

func TestReverseList(t *testing.T) {
	num := []int{1, 2, 3, 4, 5}
	h1 := pkg.CreateListFromTail(num)

	pkg.PrintLinkedList(h1)
	h2 := reverseList(h1)
	pkg.PrintLinkedList(h2)

	h3 := reverseListR(h2)
	pkg.PrintLinkedList(h3)
}

func TestRemoveNthFromEnd(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	h := pkg.CreateListFromTail(nums)
	pkg.PrintLinkedList(h)
	head := removeNthFromEnd(h, 2)
	pkg.PrintLinkedList(head)
}
