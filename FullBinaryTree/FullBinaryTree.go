/*
 * @Author: xc8801@126.com
 * @Date: 2019-03-23
 *
 * 满二叉树性质：子节点深度相同
 *			①
 *
 * 		/  		\
 *
 *	②			③
 * 4种遍历:preOder, midOrder,PostOrder, levelOrder
 */

package FullBinaryTree

import "fmt"

type FullBinaryTree struct {
	Head *Node
	Size int
}

func NewBinaryTree() *FullBinaryTree {
	return &FullBinaryTree{}
}

func (bt *FullBinaryTree) AddNode(value interface{}) {
	bt.Size++
	if bt.addNodeOne(value) {
		return
	}

	bt.addNodeMuti(value)
}

func (bt *FullBinaryTree) addNodeMuti(value interface{}) {
	currNode := bt.Head
	queue := NewQueue()
	newNode := NewNode().SetEntry(value)
	for currNode != nil {
		if currNode.Left == nil {
			currNode.SetLeft(newNode)
			break
		}
		if currNode.Right == nil {
			currNode.SetRight(newNode)
			break
		}

		queue.Push(currNode.Left)
		queue.Push(currNode.Right)
		currNode = queue.Pull().(*Node)
	}
}

func (bt *FullBinaryTree) addNodeOne(value interface{}) bool {
	if bt.Head == nil {
		bt.Head = NewNode().SetEntry(value)
		return true
	}
	return false
}

func (bt *FullBinaryTree) GetHeight() int {
	//log(bt.Size)
	x := 0
	for x^2 < bt.Size {
		x++
	}
	return x
}

func (bt *FullBinaryTree) PreOrder(node *Node) {
	if node == nil {
		return
	}

	fmt.Println(node.Entry)
	bt.PreOrder(node.Left)
	bt.PreOrder(node.Right)
}

func (bt *FullBinaryTree) MidOrder(node *Node) {
	if node == nil {
		return
	}

	bt.MidOrder(node.Left)
	fmt.Println(node.Entry)
	bt.MidOrder(node.Right)
}

func (bt *FullBinaryTree) PostOrder(node *Node) {
	if node == nil {
		return
	}

	bt.PostOrder(node.Left)
	bt.PostOrder(node.Right)
	fmt.Println(node.Entry)
}

func (bt *FullBinaryTree) LevelOrder(node *Node) {
	if node == nil {
		return
	}

	currNode := node
	queue := NewQueue()

	for currNode != nil {
		fmt.Println(currNode.Entry)
		queue.Push(currNode.Left)
		queue.Push(currNode.Right)
		currNode = queue.Pull().(*Node)
	}
}
