package HashSet

/*
 * Date: 2019/03/07
 * Email: xc8801@126.com
 *
 * Insert,Delete,Search Time Complexity:O(1)
 * Space Complexity:O(n), Map
 */

type HashSet struct {
	cache map[interface{}]bool
}

func NewHashSet() *HashSet {
	return &HashSet{
		cache: make(map[interface{}]bool),
	}
}

func (set *HashSet) Len() uint32 {
	return uint32(len(set.cache))
}

func (set *HashSet) Clear() {
	set = NewHashSet()
}

func (set *HashSet) Remove(value interface{}) {
	delete(set.cache, value)
}

func (set *HashSet) Add(value interface{}) bool {
	_, ok := set.cache[value]
	set.cache[value] = true

	return ok == true
}
