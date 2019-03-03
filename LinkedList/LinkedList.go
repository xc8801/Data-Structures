package LinkedList

/*
 * Date: 2019/03/01
 * Email: xc8801@126.com
 *
 * Insert Time Complexity:O(1), Find Time Complexity:O(n)
 * Space Complexity:O(n)
 */

type Node struct {
	data interface{}
	next *Node
}

type LinkedList struct {
	head   *Node
	lenght uint32
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (list *LinkedList) Lenght() uint32 {
	return list.lenght
}

func (list *LinkedList) IsEmpty() bool {
	return list.head == nil
}

func (list *LinkedList) InsertFirst(data interface{}) {
	node := &Node{data: data}

	if list.IsEmpty() {
		list.head = node
		return
	}

	node.next = list.head
	list.head = node
	list.lenght++
}

func (list *LinkedList) InsertLast(data interface{}) {
	node := &Node{data: data}

	current, ok := list.GetNodeByIndex(list.lenght)
	if ok {
		current.next = node
	} else {
		node.next = list.head
		list.head = node
	}
	list.lenght++

}

func (list *LinkedList) RemoveFirst() bool {
	if list.IsEmpty() {
		return false
	}

	list.head.data = nil
	list.head = list.head.next
	list.lenght--

	return true
}

func (list *LinkedList) RemoveLast() bool {
	if list.IsEmpty() {
		return false
	}

	current, ok := list.GetNodeByIndex(list.lenght - 1)
	if !ok {
		return false
	}

	current.next.data = nil
	current.next = nil
	list.lenght--
	return true
}

func (list *LinkedList) GetFirst() (interface{}, bool) {
	if list.head == nil {
		return nil, false
	}

	return list.head.data, true
}

func (list *LinkedList) GetLast() (interface{}, bool) {
	if list.IsEmpty() {
		return nil, false
	}

	node, ok := list.GetNodeByIndex(list.lenght)

	if ok {
		return node.data, ok
	}
	return nil, ok
}

func (list *LinkedList) GetNodeByIndex(index uint32) (*Node, bool) {
	if index > list.lenght || index == 0 {
		return nil, false
	}

	current := list.head
	for count := uint32(1); count < index; count++ {
		current = current.next
	}

	return current, true
}
