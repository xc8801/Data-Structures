package DoublyLinkedList

import (
	"testing"
)

func TestLinkedListCase(t *testing.T) {
	ll := NewDoublyLinkedList()

	// Check InsertLast() And GetLast()
	list := []int{1, 2, 3, 4, 5}
	for _, item := range list {
		ll.InsertFirst(item)
	}

	for _, item := range list {
		data, ok := ll.GetLast()

		if !ok || data != item {
			t.Errorf("DoublyLinkedList Getlast:%d, Item:%d, Result:%t", data, item, ok)
		}

		ll.RemoveLast()
	}
}
