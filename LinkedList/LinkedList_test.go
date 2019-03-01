package LinkedList

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	ll := NewLinkedList()

	ll.InsertLast(2)
	ll.InsertFirst(1)
	ll.InsertLast(3)

	if ll.Lenght() != 3 {
		t.Errorf("LinkList lenght:%d", ll.Lenght())
	}
	if data, _ := ll.GetLast(); data != 3 {
		t.Errorf("LinkList Getlast:%d", data)
	}

	ll.RemoveLast()
	ll.RemoveLast()
	if data, _ := ll.GetFirst(); data != 1 {
		t.Errorf("LinkList Getlast:%d", data)
	}

	ll.RemoveFirst()

	if ll.Lenght() != 0 {
		t.Errorf("LinkList lenght:%d", ll.Lenght())
	}
}
