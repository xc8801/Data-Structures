package LinkedList

//Date: 2019/03/01
//Email: xc8801@126.com

import (
	"fmt"
)

type Node struct {
	data interface{}
	next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (list *LinkedList) Lenght() uint32 {
	index := 0
	current := list.head

	for current.next != nil {
		index++
		current = current.next
	}

	return index
}

func (list *LinkedList) InsertFirst(data interface{}) {
	node := &Node{data: data}

	if list.head == nil {
		list.head = node
		list.tail = node
		return
	}

	node.next = list.head
	list.head = node
}

func (list *LinkedList) InsertLast(data interface{}) {
	node := &Node{data: data}

	if list.head == nil {
		list.head = node
		list.tail = node
		return
	}

	list.tail.next = node
	list.tail = node
}

func (list *LinkedList) RemoveFirst() bool {
	if list.head == nil {
		return false
	}

	current := list.head
	list.head = list.head.next
	current = &Node{}
}

func (list *LinkedList) RemoveLast() bool {
	if list.tail == nil {
		return false
	}

	len := list.Lenght()
	current := list.head

	for i := 0; i < len; i++ {
		current = current.next
	}

	list.tail = current
	current.next.data = nil
	current.next = nil
}

func (list *ListedList) GetFirst() (interface{}, bool) {
	if list.head == nil {
		return nil, false
	}

	return list.head, true
}

func (list *LinkedList) GetLast() (interface{}, bool) {
	if list.tail == nil {
		return nil, false
	}

	return list.tail, true
}
