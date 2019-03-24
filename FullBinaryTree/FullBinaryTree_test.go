package FullBinaryTree

import (
	"fmt"
	"testing"
)

func TestCase1(t *testing.T) {
	bt := NewBinaryTree()

	for i := 0; i < 10; i++ {
		bt.AddNode(i)
	}

	fmt.Println(bt.Size)
	bt.PreOrder(bt.Head)
	bt.MidOrder(bt.Head)
	bt.PostOrder(bt.Head)
	bt.LevelOrder(bt.Head)

}
