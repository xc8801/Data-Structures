package QueueLinkedList

import (
	"testing"
)

func TestQueueLinedListCase1(t *testing.T) {
	queue := NewQueue()

	testList := []int{1, 2, 3, 4, 5}

	for _, item := range testList {
		queue.Push(item)
	}

	if queue.GetSize() != uint32(len(testList)) {
		t.Errorf("QueueLinkedList lenght:%d", queue.GetSize())
	}

	for _, item := range testList {
		node, ok := queue.Pop()

		if !ok || node != item {
			t.Errorf("QueueLinkedList Node:%d, Item:%d, Lenght:%d", node, item, queue.GetSize())
		}
	}
}
