package TrieTree

/*
 * Date: 2019/03/11
 * Email: xc8801@126.com
 *
 * Insert,Delete,Search Time Complexity:O(n)
 * Space Complexity:O(n)
 */
import (
	"fmt"
	"strings"
)

const (
	ROOT_NODE = "/"
)

type TrieTree struct {
	head   *Node
	lenght uint32
	call   Call
}

func NewTrieTree(call Call) *TrieTree {
	return &TrieTree{
		lenght: 1,
		call:   call,
		head:   NewNode(),
	}
}

// @throw panic
func validate(path string) bool {
	if path == "" || string(path[0]) != ROOT_NODE {
		panic("'/' missing.")
	}

	return true
}

func parsePath(path string) ([]string, int) {

	paths := strings.Split(path, ROOT_NODE)
	paths = paths[1:len(paths)]

	return paths, len(paths)
}

func (trie *TrieTree) Add(path string, call Call) bool {
	if validate(path) && path == ROOT_NODE {
		trie.head.call = call
		return true
	}

	var node *Node
	nextNode := trie.head
	paths, size := parsePath(path)

	//fmt.Println(paths, size)
	for idx := 0; idx <= size-1; idx++ {
		node = nextNode
		key := paths[idx]

		nextNode.SetNext(key, NewNode())
		nextNode = nextNode.GetNext(key)
	}

	trie.lenght++
	node.SetCall(call)

	return true
}

func (trie *TrieTree) Get(path string) (Call, error) {
	if validate(path) && path == ROOT_NODE {
		return trie.head.call, nil
	}

	node := trie.head
	paths, size := parsePath(path)

	for idx := 0; idx < size-1; idx++ {
		key := paths[idx]

		if node.next == nil {
			return trie.call, fmt.Errorf(key)
		}

		node = node.GetNext(key)
	}

	return node.call, nil
}

func (trie *TrieTree) DepthFirstSearchShow() {

}

func (trie *TrieTree) BreadthFirstSearchShow() {

}
