package binarytree

import (
	"math/rand"
	"time"
)

// RandGen random generator
type RandGen struct {
	src       rand.Source
	cache     int64
	remaining int
}

func (b *RandGen) Bool() bool {
	if b.remaining == 0 {
		b.cache, b.remaining = b.src.Int63(), 63
	}

	result := b.cache&0x01 == 1
	b.cache >>= 1
	b.remaining--

	return result
}

func NewRand() *RandGen {
	return &RandGen{src: rand.NewSource(time.Now().UnixNano())}
}

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Left  *TreeNode
	Val   int
	Right *TreeNode
	Next  *TreeNode
}

// New returns a new, random binary tree
// holding the values 1k, 2k, ..., nk.
func New(n, k int) *TreeNode {
	var t *TreeNode
	for _, v := range rand.Perm(n) {
		t = insert(t, (1+v)*k)
	}
	return t
}

// insert a value into tree
func insert(t *TreeNode, v int) *TreeNode {
	if t == nil {
		return &TreeNode{nil, v, nil, nil}
	}
	// 随机插入左右子树
	r := NewRand()
	if r.Bool() {
		t.Left = insert(t.Left, v)
		return t
	}
	t.Right = insert(t.Right, v)
	return t
}
