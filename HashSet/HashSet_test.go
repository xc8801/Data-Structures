package HashSet

import (
	"fmt"
	"testing"
)

func TestHashSetCase(t *testing.T) {
	set := NewHashSet()

	set.Add(1)
	set.Add(2)
	set.Add(3)

	set.Remove(3)
	fmt.Println(set)
}
