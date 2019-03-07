package SkipList

type ForwardNode []*Node

type Node struct {
	forward ForwardNode
	score   int
	value   interface{}
}

func (n *Node) Clear() {
	n.score = 0
	n.value = nil
	n.forward = nil
}

func (n *Node) GetLevel() uint16 {
	return uint16(len(n.forward))
}

func NewForwardNode(level uint16) ForwardNode {
	return make([]*Node, level)
}

func NewNode(level uint16, score int, value interface{}) *Node {
	return &Node{
		score:   score,
		value:   value,
		forward: NewForwardNode(level),
	}
}
