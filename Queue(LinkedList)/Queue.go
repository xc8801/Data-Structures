package QueueLinkedList

/*
 * Date: 2019/03/02
 * Email: xc8801@126.com
 *
 * Push Time Complexity:O(1), Pop Time Complexity:O(1)
 * Space Complexity:O(n)
 */

type Node struct {
	data interface{}
	pre  *Node
	next *Node
}

type QueueLinkedList struct {
	head   *Node
	tail   *Node
	lenght uint32
}

func NewQueue() *QueueLinkedList {
	return &QueueLinkedList{}
}

func (queue *QueueLinkedList) GetSize() uint32 {
	return queue.lenght
}

func (queue *QueueLinkedList) Empty() bool {
	return queue.head == nil
}

func (queue *QueueLinkedList) Push(data interface{}) {
	node := &Node{
		data: data,
		next: queue.head,
	}

	if queue.tail == nil {
		queue.tail = node
	} else {
		queue.head.pre = node
	}

	queue.head = node
	queue.lenght++
}

func (queue *QueueLinkedList) Pop() (interface{}, bool) {
	if queue.Empty() {
		return nil, false
	}

	data := queue.tail.data
	queue.tail.data = nil

	if queue.tail == queue.head {
		queue.head = nil
	}

	queue.tail = queue.tail.pre
	queue.lenght--

	return data, true
}
