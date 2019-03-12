package TrieTree

/*
 * Date: 2019/03/11
 * Email: xc8801@126.com
 *
 * Insert,Delete,Search Time Complexity:O(n)
 * Space Complexity:O(n)
 */

type Call func(msg string)

type Node struct {
	call Call
	next map[string]*Node
}

func NewNode() *Node {
	return &Node{
		call: nil,
		next: make(map[string]*Node),
	}
}

func (node *Node) SetCall(call Call) *Node {
	node.call = call
	return node
}

func (node *Node) ExitsByKey(key string) bool {
	_, ok := node.next[key]
	return ok
}

func (node *Node) HasNext() bool {
	return node.next == nil && len(node.next) != 0
}

func (node *Node) GetNext(key string) *Node {
	next, _ := node.next[key]
	return next
}

func (node *Node) SetNext(key string, newNode *Node) *Node {
	if !node.ExitsByKey(key) {
		node.next[key] = newNode
	}
	return node
}
