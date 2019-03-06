package SkipList

import (
	"fmt"
	"testing"
)

func TestSkipListCase1(t *testing.T) {
	testList := []int{1, 2, 3}

	skipList := NewSkipList(16)

	for idx, item := range testList {
		skipList.Insert(idx, item)
	}

	fmt.Println(skipList)
	skipList.Display()
}
