package tree

import (
	"fmt"
	"testing"
)

func TestIsCompleteTree(t *testing.T) {
	preorder := []int{1, 2, 4, 5, 3, 6}
	inorder := []int{4, 2, 5, 1, 6, 3}
	root := buildTreeFromPreIn(preorder, inorder)
	layerOrderTraverse(root)
	r := isCompleteTree(root)
	fmt.Println(r)
}
