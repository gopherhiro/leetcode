package linkedlist

import (
	"leetcode/pkg"
	"testing"
)

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
