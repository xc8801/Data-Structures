package HashTable

/*
 * Date: 2019/03/08
 * Email: xc8801@126.com
 *
 * Insert, Delete,Search Time Complexity:O(n)
 * Space Complexity:O(n)
 */

type Entry struct {
	key   string
	value interface{}
}

type LinkedListNode struct {
	Entry
	next *LinkedListNode
	prev *LinkedListNode
}

type LinkedList struct {
	head   *LinkedListNode
	lenght uint32
	index  uint32 //Iterator
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func NewEntry(key string, value interface{}) *Entry {
	return &Entry{
		key:   key,
		value: value,
	}
}

func NewLinkedListNode(entry Entry) *LinkedListNode {
	return &LinkedListNode{
		next:  nil,
		entry: entry,
	}
}

func (list *LinkedList) Lenght() uint32 {
	return list.lenght
}

func (list *LinkedList) hasNext() bool {
	return list.next != nil
}

func (list *LinkedList) next() *LinkedListNode {
	if !hasNext() {
		return nil
	}

	list.index++
	return list.next
}

func (list *LinkedList) Insert(key string, value interface{}) {
	if list.Change(key, value) {
		return
	}

	entry := NewEntry(key, value)
	node := NewLinkedListNode(entry)

	list.head.prev = node
	node.next = list.head
	list.head = node

	list.lenght++
}

func (list *LinkedList) Remove(key string) bool {
	var ok bool
	var currNode *LinkedListNode

	if currNode, ok = list.Search(key); !ok {
		return false
	}

	currNode.Entry = nil
	currNode.prev = currNode.next
	list.lenght--

	return true
}

func (list *LinkedList) Change(key string, value interface{}) bool {
	if currNode, ok := list.Search(key); ok {
		currNode.value = value
		return true
	}

	return false
}

func (list *LinkedList) Get(key string) (interface{}, bool) {
	if currNode, ok := list.Search(key); ok {
		return currNode.value, true
	}

	return nil, false
}

func (list *LinkedList) Search(key string) (*LinkedListNode, bool) {
	currNode := list.head

	for currNode != nil {
		if currNode.key == key {
			return currNode, true
		}
		currNode = currNode.next
	}

	return nil, false
}
