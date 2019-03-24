/*
 * @Author: xc8801@126.com
 * @Date: 2019-03-24 15:02:31
 *
 */
package CompleteBinaryTree

import "fmt"

const (
	Default_Max_Lenght = 100
)

type CompleteBinaryTree struct {
	BinaryTree []interface{}
	MaxLenght  int
	Lenght     int
}

func NewCompleteBinaryTree() *CompleteBinaryTree {
	return &CompleteBinaryTree{
		MaxLenght:  Default_Max_Lenght,
		BinaryTree: make([]interface{}, Default_Max_Lenght, Default_Max_Lenght),
	}
}

func (bt *CompleteBinaryTree) SetMaxLenght(maxLenght int) *CompleteBinaryTree {
	bt.MaxLenght = maxLenght
	bt.BinaryTree = make([]interface{}, maxLenght, maxLenght)
	return bt
}

func (bt *CompleteBinaryTree) GetLenght() int {
	return bt.Lenght
}

// idx0:null, idx1:rootNode, idx2:leftNode, idx3:rightNode
//当前节点IDX:n, parentNode = int(n/2)
//当前节点IDX:n，childLeftNode = n * 2, childRightNode = n * 2 + 1
func (bt *CompleteBinaryTree) AddNode(entry interface{}) bool {
	if bt.GetLenght() > bt.MaxLenght {
		return false
	}

	idx := bt.GetLenght() + 1
	bt.BinaryTree[idx] = entry
	bt.Lenght++
	return true
}

func (bt *CompleteBinaryTree) LevelOrder() {
	for _, node := range bt.BinaryTree {
		fmt.Println(node)
	}
}
