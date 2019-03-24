/*
 * @Author: xc8801@126.com
 * @Date: 2019-03-24 13:17:52
 * Copy FullBinaryTree/Queue(LinkedList).go
 */
package BinarySearchTree

// Node
type QueueNode struct {
	Next *QueueNode
	Data interface{}
}

func NewQueueNode() *QueueNode {
	return &QueueNode{}
}

func (node *QueueNode) SetData(data interface{}) *QueueNode {
	node.Data = data
	return node
}

func (node *QueueNode) SetNext(q *QueueNode) *QueueNode {
	node.Next = q
	return node
}

func (node *QueueNode) Clear() {
	node.Next = nil
	node.Data = nil
}

// Queue
type Queue struct {
	Head   *QueueNode
	Lenght uint32
}

func NewQueue() *Queue {
	return &Queue{
		Head: NewQueueNode(),
	}
}
func (queue *Queue) IsEmpty() bool {
	return queue.Lenght == 0
}

func (queue *Queue) Push(data interface{}) {
	preNode := queue.Head
	currNode := queue.Head.Next

	node := NewQueueNode().SetData(data)

	for currNode != nil {
		preNode = currNode
		currNode = currNode.Next
	}

	queue.Lenght++
	preNode.SetNext(node)
}

func (queue *Queue) Pull() interface{} {
	Head := queue.Head
	currNode := queue.Head.Next
	if currNode != nil {
		Head.Next = currNode.Next
	}

	queue.Lenght--
	defer currNode.Clear()
	return currNode.Data
}
