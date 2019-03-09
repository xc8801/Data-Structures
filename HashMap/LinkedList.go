package HashMap

/*
 * Date: 2019/03/08
 * Email: xc8801@126.com
 *
 * LinkedList Insert, Delete,Search Time Complexity:O(n)
 * Space Complexity:O(n)
 */

type Entry struct {
	key       string
	value     interface{}
	keysIndex uint32
}

type LinkedListNode struct {
	*Entry
	next *LinkedListNode
}

type LinkedList struct {
	head   *LinkedListNode
	lenght uint32
	index  uint32 //Iterator Index
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func NewEntry(key string, value interface{}, idx uint32) *Entry {
	return &Entry{
		key:       key,
		value:     value,
		keysIndex: idx,
	}
}

func NewLinkedListNode(entry *Entry) *LinkedListNode {
	return &LinkedListNode{
		Entry: entry,
	}
}

func (list *LinkedList) HasNext() bool {
	if list.lenght == 0 || list.index >= list.lenght {
		return false
	}
	return true
}

func (list *LinkedList) Next() *LinkedListNode {
	if !list.HasNext() {
		return nil
	}

	currNode := list.head

	for idx := uint32(0); idx <= list.index; idx++ {
		currNode = currNode.next
	}

	list.index++
	return currNode
}

func (list *LinkedList) ResetIndex() {
	list.index = 0
}

func (list *LinkedList) Lenght() uint32 {
	return list.lenght
}

func (list *LinkedList) Remove(key string) (uint32, bool) {
	var ok bool
	var currNode, preNode *LinkedListNode

	if currNode, preNode, ok = list.Search(key); !ok {
		return 0, false
	}

	currNode.key = ""
	currNode.value = nil
	preNode.next = currNode.next
	list.lenght--

	return currNode.keysIndex, true
}

func (list *LinkedList) Change(key string, value interface{}) (*LinkedListNode, bool) {
	if currNode, _, ok := list.Search(key); ok {
		currNode.value = value
		return currNode, true
	}

	return nil, false
}

func (list *LinkedList) Insert(key string, value interface{}, keysIdx uint32) *LinkedListNode {
	if oldNode, ok := list.Change(key, value); ok {
		return oldNode
	}

	entry := NewEntry(key, value, keysIdx)
	node := NewLinkedListNode(entry)

	node.next = list.head
	list.head = node

	list.lenght++
	return node
}

func (list *LinkedList) Get(key string) (interface{}, bool) {
	if currNode, _, ok := list.Search(key); ok {
		return currNode.value, true
	}

	return nil, false
}

func (list *LinkedList) Search(key string) (*LinkedListNode, *LinkedListNode, bool) {
	currNode := list.head
	preNode := list.head

	for currNode != nil {
		if currNode.key == key {
			return currNode, preNode, true
		}
		preNode = currNode
		currNode = currNode.next

	}

	return nil, nil, false
}
