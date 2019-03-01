package LinkedList

/*
 * Date: 2019/03/01
 * Email: xc8801@126.com
 *
 * Time Complexity:O(n), Space Complexity:O(n)
 */

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
	var index uint32
	current := list.head

	if current == nil {
		return index
	}

	for current.next != nil {
		index++
		current = current.next
	}
	return index + 1
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

	len := list.Lenght()

	if len == 1 {
		list.head.data = nil
		list.tail = nil
		list.head = nil
		return true
	}

	list.head.data = nil
	list.head = list.head.next

	return true
}

func (list *LinkedList) RemoveLast() bool {
	if list.tail == nil {
		return false
	}

	len := list.Lenght()
	current := list.head

	if len == 1 {
		list.head.data = nil
		list.tail = nil
		list.head = nil
		return true
	}

	for i := uint32(1); i < len-1; i++ {
		current = current.next
	}

	list.tail.data = nil
	list.tail = current
	list.tail.next = nil

	return true
}

func (list *LinkedList) GetFirst() (interface{}, bool) {
	if list.head == nil {
		return nil, false
	}

	return list.head.data, true
}

func (list *LinkedList) GetLast() (interface{}, bool) {
	if list.tail == nil {
		return nil, false
	}

	return list.tail.data, true
}
