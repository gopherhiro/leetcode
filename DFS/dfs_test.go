package DFS

import (
	"fmt"
	"testing"
)

func TestSumPathNumbers(t *testing.T) {
	preorder := []int{1, 7, 9, 2, 9}
	inorder := []int{7, 1, 2, 9, 9}
	root := buildTreeFromPreIn(preorder, inorder)
	r1 := sumPathNumbers(root)
	r2 := sumPathNumber(root)
	fmt.Println(r1, r2)
}
