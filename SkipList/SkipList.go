package SkipList

/*
 * Date: 2019/03/02
 * Email: xc8801@126.com
 *
 * Insert,Find, Delete AVG Time Complexity:O(log(n)), Wrost Time Complexity:O(n)
 * Space Complexity:O(n)
 */

import (
	"fmt"
	"math/rand"
	"time"
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
	DEFAULT_P = 0.5 // p^10
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

func (list *SkipList) randLevel() uint16 {
	level := uint16(1)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for r.Float64() < list.p {
		level++
	}

	if level > list.maxLevel {
		level = list.maxLevel
	}

	return level
}

func (list *SkipList) Insert(newScore int, newValue interface{}) {
	var next *Node
	var prevs = list.head

	newLevel := list.randLevel() //[1~maxLevel]
	newNode := NewNode(newLevel, newScore, newValue)

	for level := list.maxLevel - 1; level <= list.maxLevel; level-- {
		next = prevs[level]

		for next != nil && next.score <= newScore {
			prevs = next.forward
			next = next.forward[level]
		}

		if level <= newLevel-1 {
			newNode.forward[level] = prevs[level]
			prevs[level] = newNode
		}
	}

	list.lenght++
}

func (list *SkipList) Display() {

	var idx = 0
	next := list.head[0]

	for next != nil {
		fmt.Println(next, idx)
		next = next.forward[0]
		idx++
	}
}
