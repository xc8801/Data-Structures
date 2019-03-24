package CompleteBinaryTree

import (
	"fmt"
	"testing"
)

func TestCase1(t *testing.T) {
	bt := NewCompleteBinaryTree().SetMaxLenght(11)

	for i := 0; i < 10; i++ {
		bt.AddNode(i)
	}

	fmt.Println(bt.GetLenght())
	bt.LevelOrder()
}
