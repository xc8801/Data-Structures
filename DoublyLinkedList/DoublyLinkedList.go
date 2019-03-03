package DoublyLinkedList

/*
 * Date: 2019/03/02
 * Email: xc8801@126.com
 *
 * Insert Time Complexity:O(1), Find Time Complexity:O(n)
 * Space Complexity:O(n)
 */

type Node struct {
	data interface{}
	next *Node
	pre  *Node
}

type DoublyLinkedList struct {
	head   *Node
	tail   *Node
	lenght uint32
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (list *DoublyLinkedList) Lenght() uint32 {
	return list.lenght
}

func (list *DoublyLinkedList) IsEmpty() bool {
	return list.lenght == 0
}

func (list *DoublyLinkedList) clearNode(node *Node) {
	node = &Node{}
}

func (list *DoublyLinkedList) insertEmptyList(data interface{}) {
	node := &Node{
		data: data,
	}

	list.head = node
	list.tail = node

	node.pre = node
	node.next = node

	list.lenght++
}

func (list *DoublyLinkedList) InsertFirst(data interface{}) {
	if list.IsEmpty() {
		list.insertEmptyList(data)
		return
	}

	node := &Node{
		data: data,
		next: list.head,
		pre:  nil,
	}

	list.head.pre = node
	list.head = node
	list.lenght++
}

func (list *DoublyLinkedList) RemoveFirst() bool {
	if list.IsEmpty() {
		return false
	}

	headNode := list.head
	list.head = list.head.next
	if list.head != nil {
		list.head.pre = nil
	}

	list.clearNode(headNode)
	list.lenght--

	return true
}

func (list *DoublyLinkedList) GetFirst() (interface{}, bool) {
	if list.IsEmpty() {
		return nil, false
	}

	return list.head.data, true
}

func (list *DoublyLinkedList) InsertLast(data interface{}) {
	if list.IsEmpty() {
		list.insertEmptyList(data)
		return
	}

	node := &Node{
		data: data,
		pre:  list.tail,
	}

	list.tail.next = node
	list.tail = node
	node.next = node
	list.lenght++
}

func (list *DoublyLinkedList) RemoveLast() bool {
	if list.IsEmpty() {
		return false
	}

	lastNode := list.tail
	list.tail = list.tail.pre
	if list.tail != nil {
		list.tail.next = nil
	}

	list.clearNode(lastNode)
	list.lenght--

	return true
}

func (list *DoublyLinkedList) GetLast() (interface{}, bool) {
	if list.IsEmpty() {
		return nil, false
	}

	return list.tail.data, true
}
