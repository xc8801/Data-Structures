package SkipList

import (
	"math/rand"
	"testing"
)

func TestSkipListCase1(t *testing.T) {
	maxLen := 10
	testList := make([]int, maxLen)

	for idx := 0; idx < maxLen; idx++ {
		xxx := rand.Intn(maxLen)
		testList[xxx] = xxx
	}

	//fmt.Println(testList)
	skipList := NewSkipList(16)
	/*
		for _, item := range testList {
			skipList.Insert(item, fmt.Sprintf("VALUE-%d", item))
		}
	*/
	skipList.Insert(1, "VALUE-1")

	skipList.Insert(2, "VALUE-2")
	skipList.Insert(3, "VALUE-3")

	skipList.Insert(6, "VALUE-6")
	skipList.Insert(5, "VALUE-5")
	skipList.Insert(4, "VALUE-4")
	skipList.Insert(3, "VALUE-4")

	skipList.Remove(1)
	skipList.Display()
}
