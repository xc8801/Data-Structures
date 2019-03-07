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

type SkipList struct {
	p        float64
	head     ForwardNode
	lenght   uint32
	maxLevel uint16
}

const (
	MAX_LEVEL = 32
	DEFAULT_P = 0.6 // p^10
)

func NewSkipList(level uint16) *SkipList {
	if level > MAX_LEVEL {
		level = MAX_LEVEL
	}

	return &SkipList{
		p:        DEFAULT_P, // range (0~1)
		head:     NewForwardNode(level),
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

func (list *SkipList) GetPreNode(score int) (*Node, []ForwardNode) {
	var next *Node
	var prevs = list.head
	var updates = make([]ForwardNode, list.maxLevel)

	for level := list.maxLevel - 1; level <= list.maxLevel; level-- {
		next = prevs[level]

		for next != nil && next.score < score {
			prevs = next.forward
			next = next.forward[level]
		}

		updates[level] = prevs
	}

	return next, updates
}

func (list *SkipList) Insert(newScore int, newValue interface{}) {

	newLevel := list.randLevel() //[1~maxLevel]
	newNode := NewNode(newLevel, newScore, newValue)

	currNode, updateNodes := list.GetPreNode(newScore)

	if currNode != nil && currNode.score == newScore {
		currNode.value = newValue
		return
	}

	for level := newLevel - 1; level <= newLevel; level-- {
		newNode.forward[level] = updateNodes[level][level]
		updateNodes[level][level] = newNode
	}

	list.lenght++
}

func (list *SkipList) Remove(score int) bool {
	currNode, updateNodes := list.GetPreNode(score)

	if currNode == nil || currNode.score != score {
		return false
	}

	list.lenght--
	currLevel := currNode.GetLevel()

	for level := currLevel - 1; level <= currLevel; level-- {
		updateNodes[level][level] = currNode.forward[level]
	}

	currNode.Clear()
	return true
}

func (list *SkipList) Get(score int) (interface{}, bool) {
	currNode, _ := list.GetPreNode(score)

	if currNode == nil || currNode.score != score {
		return nil, false
	}

	return currNode.value, true
}

func (list *SkipList) Display() {
	next := list.head

	for next != nil && next[0] != nil {
		fmt.Println("Display:", next[0].score, next[0])
		next = next[0].forward
	}

	time.Sleep(time.Second)
}
