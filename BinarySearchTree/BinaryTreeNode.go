/*
 * @Author: xc8801@126.com
 * @Date: 2019-03-24 21:28:05
 */
package BinarySearchTree

type Entry struct {
	Score int
	Data  interface{}
}

func NewEntry(score int, data interface{}) *Entry {
	return &Entry{
		Score: score,
		Data:  data,
	}
}

type Node struct {
	Left  *Node
	Right *Node
	Entry *Entry
}

func NewNode() *Node {
	return &Node{}
}

func (n *Node) Less(newNode *Node) bool {
	if n.Entry.Score < newNode.Entry.Score {
		return true
	}
	return false
}

func (n *Node) More(newNode *Node) bool {
	if n.Entry.Score > newNode.Entry.Score {
		return true
	}
	return false
}

func (n *Node) Match(newNode *Node) bool {
	if n.Entry.Score == newNode.Entry.Score {
		return true
	}
	return false
}

func (n *Node) CloneValue(newNode *Node) {
	n.Entry = newNode.Entry
}

func (n *Node) SetEntry(score int, entry interface{}) *Node {
	n.Entry = NewEntry(score, entry)
	return n
}

func (n *Node) SetLeft(left *Node) *Node {
	n.Left = left
	return n
}

func (n *Node) SetRight(right *Node) *Node {
	n.Right = right
	return n
}

func (n *Node) Clear() {
	n.Entry = nil
	n.Left = nil
	n.Right = nil
}
