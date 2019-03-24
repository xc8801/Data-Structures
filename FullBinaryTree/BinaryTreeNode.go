/*
 * @Author: xc8801@126.com
 * @Date: 2019-03-24 12:11:08
 */

package FullBinaryTree

type Node struct {
	Left  *Node
	Right *Node
	Entry interface{}
}

func NewNode() *Node {
	return &Node{}
}

func (n *Node) SetEntry(entry interface{}) *Node {
	n.Entry = entry
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
