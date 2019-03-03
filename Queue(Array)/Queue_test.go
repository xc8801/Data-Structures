package QueueArray

import (
	"testing"
)

func TestQueueArrayCase1(t *testing.T) {
	queue := NewQueueArray()

	testList := []int{1, 2, 3, 4, 5}

	for _, item := range testList {
		queue.Push(item)
	}

	if queue.GetSize() != len(testList) {
		t.Errorf("QueueArray lenght:%d", queue.GetSize())
	}

	for _, item := range testList {
		node, ok := queue.PopTop()

		if !ok || node != item {
			t.Errorf("QueueArray Node:%d, Item:%d, Lenght:%d", node, item, queue.GetSize())
		}
	}
}
