package TrieTree

import (
	"fmt"
	"testing"
)

func echo(msg string) {
	fmt.Println(msg)
}

func TestTrieTreeCase(t *testing.T) {
	trie := NewTrieTree(echo)

	trie.Add("/a", echo)
	fmt.Println("Test======>/a", trie.head, trie.head.next)

	trie.Add("/a/b", echo)
	fmt.Println("Test======>/a/b", trie.head.next["a"])

	trie.Add("/a/b/c", echo)
	fmt.Println("Test======>/a/b/c", trie.head.next["a"].next["b"])

	trie.Add("/1/2/3", echo)

	fmt.Println("Test======>/1", trie.head.next)
	fmt.Println("Test======>/1/2", trie.head.next["1"])
	fmt.Println("Test======>/1/2/3", trie.head.next["1"].next["2"])
	if call, err := trie.Get("/a/b"); err != nil {
		call(err.Error())
	}
	fmt.Println("#####END#####")
}
