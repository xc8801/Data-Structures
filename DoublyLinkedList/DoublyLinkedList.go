package DoublyLinkedList

import (
	"fmt"
)

type Node struct {
	data interface{}
	next *Node
	pre  *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}
