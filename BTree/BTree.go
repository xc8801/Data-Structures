package BTree

/*
 * Date: 2019/03/10
 * Email: xc8801@126.com
 *
 * Insert,Remove Time Complexity:O(logN), Search Time Complexity:O(logN)
 * Space Complexity:O(n)
 * Wiki: https://en.wikipedia.org/wiki/B-tree
 */

var (
	maxDegree = 3
)

type BTree struct {
	root *Node

	high      uint32
	childSize uint32

	maxDegree uint32
}

func NewBTree() *BTree {
	return &BTree{
		size:      0,
		high:      1,
		childSize: 0,

		maxDegree: maxDegree,
		root:      NewNode(maxDegree, nil),
	}
}

func (btree *BTree) transfer(node *Node) {
	for node.size == node.maxDegree {
		node.
	}
}

func (btree *BTree) Remove(ID uint32) bool {

}

func (btree *BTree) Insert(ID uint32, value interface{}) {

	btree.childSize++
	if btree.high == 1 && btree.root.size < btree.maxDegree {
		entry := NewEntry(ID, value, btree.root)
		btree.root.insert(entry)
		btree.transfer(btree.root)
		return
	}
}

func (btree *BTree) Search(ID uint32) (interface{}, bool) {

}
