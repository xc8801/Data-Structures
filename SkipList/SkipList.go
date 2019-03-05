package SkipList

/*
 * Date: 2019/03/02
 * Email: xc8801@126.com
 *
 * Insert,Find, Delete AVG Time Complexity:O(log(n)), Wrost Time Complexity:O(n)
 * Space Complexity:O(n)
 */

import (
	"math/rand"
)

type Node struct {
	forward []*Node
	score   int
	value   interface{}
}

func NewNode(level uint16, score int, value interface{}) *Node {
	return &Node{
		score:   score,
		value:   value,
		forward: make([]*Node, level),
	}
}

func (n *Node) GetLevel() int {
	return len(n.forward)
}

func (n *Node) hasNext() bool {
	return n.GetLevel() != 0
}

type SkipList struct {
	p        float64
	head     []*Node
	lenght   uint32
	maxLevel uint16
}

const (
	MAX_LEVEL = 32
	DEFAULT_P = 0.25 // p^10
)

func NewSkipList(level uint16) *SkipList {
	if level > MAX_LEVEL {
		level = MAX_LEVEL
	}

	return &SkipList{
		p:        DEFAULT_P, // range (0~1)
		head:     make([]*Node, level),
		lenght:   0,
		maxLevel: level,
	}
}

func (list *SkipList) SetP(p float64) bool {
	if p > 1 || p < 0 {
		return false
	}
	list.p = p
	return true
}

func (list *SkipList) randLevel() (level uint16) {
	for level = uint16(0); level < list.maxLevel && rand.Float64() < list.p; {
		level++
	}
	return
}

func (list *SkipList) Insert(newScore int, newValue interface{}) {
	var forwardNode *Node
	var backwardNodes = list.head

	newLevel := list.randLevel() //[1~maxLevel]
	newNode := NewNode(newLevel, newScore, newValue)

	for level := list.maxLevel - 1; level >= 0; level-- {
		forwardNode = backwardNodes[level]

		for forwardNode.score <= newScore {
			backwardNodes = forwardNode.forward
			forwardNode = forwardNode.forward[level]
		}

		if level <= newLevel {
			newNode.forward[level] = backwardNodes[level]
			backwardNodes[level] = newNode
		}
	}

	list.lenght++
}
