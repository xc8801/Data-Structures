/*
 * @Author: xc8801@126.com
 * @Date: 2019-03-24 21:29:32
 */

package BinarySearchTree

type BinarySearchTree struct {
	Head   *Node
	Lenght int32
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{
		Lenght: 0,
		Head:   NewNode(),
	}
}

func (bst *BinarySearchTree) ADD(score int, value interface{}) bool {
	bst.Lenght++
	if bst.addOne(score, value) {
		return true
	}

	return bst.addMuti(score, value)
}

func (bst *BinarySearchTree) addOne(score int, value interface{}) bool {
	if bst.Head != nil {
		return false
	}

	bst.Head = NewNode().SetEntry(score, value)
	return true
}

func (bst *BinarySearchTree) addMuti(score int, value interface{}) bool {
	newNode := NewNode().SetEntry(score, value)
	node := bst.queryAddPreNode(newNode)

	bst.setCurr(node, newNode)
	return bst.setLeft(node, newNode) != bst.setRight(node, newNode)
}

func (bst *BinarySearchTree) queryAddPreNode(newNode *Node) *Node {
	queue := NewQueue()
	preNode, currNode := bst.Head, bst.Head

	for currNode != nil {
		if currNode.Match(newNode) {
			return currNode
		}
		if currNode.More(newNode) {
			queue.Push(currNode.Left)
		}
		if currNode.Less(newNode) {
			queue.Push(currNode.Right)
		}

		preNode = currNode
		currNode = queue.Pull().(*Node)
	}

	return preNode
}

func (bst *BinarySearchTree) setLeft(node, newNode *Node) bool {
	if node.Less(newNode) {
		node.SetLeft(newNode)
		return true
	}
	return false
}

func (bst *BinarySearchTree) setRight(node, newNode *Node) bool {
	if node.More(newNode) {
		node.SetRight(newNode)
		return true
	}
	return false
}

func (bst *BinarySearchTree) setCurr(node, newNode *Node) {
	if node.Match(newNode) {
		node.CloneValue(newNode)
	}
}

func (bst *BinarySearchTree) DEL(value interface{}) bool {
	return false
}

func (bst *BinarySearchTree) queryMinByNode(node *Node) *Node {
	return nil
}

func (bst *BinarySearchTree) queryMaxByNode(node *Node) *Node {
	return nil
}

func (bst *BinarySearchTree) Query(score int) (*Node, bool) {
	return nil, false
}

func (bst *BinarySearchTree) PreOrder(node *Node) {

}

func (bst *BinarySearchTree) MidOrder(node *Node) {

}

func (bst *BinarySearchTree) Postrder(node *Node) {

}

func (bst *BinarySearchTree) LevelOrder(node *Node) {

}
