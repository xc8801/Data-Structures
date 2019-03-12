package HashMap

/*
 * Date: 2019/03/07
 * Email: xc8801@126.com
 *
 * Insert,Delete,Search Time Complexity:O(1)
 * Space Complexity:O(n)
 */

const (
	DEFAULT_LOAD_FACTOR   = 5
	DEFAULT_HASHARRY_SIZE = 10
)

type HashMap struct {
	rebuild bool

	loadFactor uint32
	lenght     uint32

	bu   []*LinkedList
	temp []*LinkedList

	hashFunc func(key string) uint32
}

func NewHashMap() *HashMap {
	return &HashMap{
		loadFactor: DEFAULT_LOAD_FACTOR,
		temp:       make([]*LinkedList, DEFAULT_HASHARRY_SIZE),
		table:      make([]*LinkedList, DEFAULT_HASHARRY_SIZE),
	}
}

func (hm *HashMap) Get(key string) interface{} {
	return nil
}

func (hm *HashMap) Del(key string) bool {
	return true
}

func (hm *HashMap) Search(key string) *Entry {
	return nil
}

func (hm *HashMap) Set(key string, value interface{}) bool {

	return true
}

func (hm *HashMap) SetLoadFactor(loadFactor uint32) *HashMap {
	hm.loadFactor = loadFactor

	return hm
}
