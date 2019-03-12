package BTree

/*
 * Date: 2019/03/10
 * Email: xc8801@126.com
 *
 * Insert Time Complexity:O(logN), Find Time Complexity:O(logN)
 * Space Complexity:O(n)
 * Wiki: https://en.wikipedia.org/wiki/B-tree
 */

import (
	"math"
)

type Node struct {
	size      uint32
	childSize uint32

	maxDegree uint32

	prev  *Node
	child []*Node
	entry []*Entry
}

type Entry struct {
	id    uint32
	value interface{}
}

func NewNode(degree uint32, prev *Node) *Node {
	return &Node{
		size:      1,
		childSize: 0,
		prev:      prev,
		maxDegree: degree,
		child:     make([]*Node, degree),
		entry:     make([]*Entry, degree),
	}
}

func NewEntry(id uint32, value interface{}) *Entry {
	return &Entry{
		id:    id,
		value: value,
	}
}

func (node *Node) Full() bool {
	return node.size == node.maxDegree
}

func (node *Node) Empty() bool {
	return node.size == 0
}

//TODO:quickSort or LinkList
func (node *Node) Sort() {
	var temp *Entry

	for idx := 0; idx <= node.size; idx++ {
		for subIdx := 0; subIdx < node.size-idx; subIdx++ {
			if node.entry[subIdx] > node.entry[subIdx+1] {

				temp = node.entry[subIdx]
				node.entry[subIdx] = node.entry[subIdx+1]
				node.entry[subIdx+1] = temp
			}
		}
	}
}

func (node *Node) Insert(entry *Entry) bool {
	if node.Full {
		return false
	}

	node.size++
	node.entry[node.size] = entry
	node.Sort()

	return true
}

func (node *Node) Remove(ID uint32) bool {
	for idx := 0; idx <= node.size; idx++ {
		if node.entry[idx].id == ID {
			node.size--
			node.entry[idx] == nil
			node.Sort()
			return true
		}
	}
	return false
}

func (node *Node) Query(ID uint32) (interface{}, bool) {
	for idx := 0; idx <= node.size; idx++ {
		if node.entry[idx].id == ID {
			return node.entry[idx].vakue, true
		}
	}
	return nil, false
}

func (node *Node) GetTransferNode() *Node {
	idx := (node.maxDegree % 2) + math.Floor(node.maxDegree/2) - 1

}
