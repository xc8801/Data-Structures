package HashMap

/*
 * Date: 2019/03/07
 * Email: xc8801@126.com
 *
 * Insert,Delete,Search Time Complexity:O(1)
 * Space Complexity:O(n), linkedlist(keys + values) + keys(array) + value(array)
 */
import (
	"sync"
)

const (
	DEFAULT_LOAD_FACTOR   = 0.75
	DEFAULT_HASHARRY_SIZE = 4
)

type HashMap struct {
	sync.Mutex

	loadFactor float32
	lenght     uint32

	bucketCap uint32
	bucket    []*LinkedList //Real Store
	temp      []*LinkedList //Resize

	//Python Keys(), Values()
	keys   []string
	values []interface{}

	//HashMethod
	hashFunc func(key string, hashLen uint32) uint32
}

func NewHashMap() *HashMap {
	return &HashMap{
		hashFunc:   HashCode,
		bucketCap:  DEFAULT_HASHARRY_SIZE,
		loadFactor: DEFAULT_LOAD_FACTOR,

		//temp:       make([]*LinkedList, DEFAULT_HASHARRY_SIZE),
		bucket: make([]*LinkedList, DEFAULT_HASHARRY_SIZE),
	}
}

func (hm *HashMap) Get(key string) (interface{}, bool) {
	if node, ok := hm.Search(key); ok {
		return node.value, true
	}

	return nil, false
}

func (hm *HashMap) Del(key string) bool {
	hm.Lock()
	defer hm.Unlock()

	var ok bool
	var keysIdx uint32

	bucketIdx := hm.hashFunc(key, hm.bucketCap)
	if hm.bucket[bucketIdx] == nil || hm.bucket[bucketIdx].Empty() {
		return false
	}

	if keysIdx, ok = hm.bucket[bucketIdx].Remove(key); !ok {
		return false
	}

	hm.lenght--
	//TODO, 删除数据可以存放临时链表队列，下次添加直接替换
	hm.keys = append(hm.keys[:keysIdx], hm.keys[keysIdx+1:]...)
	hm.values = append(hm.values[:keysIdx], hm.values[keysIdx+1:]...)

	return true
}

func (hm *HashMap) Search(key string) (*LinkedListNode, bool) {
	hm.Lock()
	defer hm.Unlock()

	bucketIdx := hm.hashFunc(key, hm.bucketCap)
	node, _, ok := hm.bucket[bucketIdx].Search(key)

	return node, ok
}

func (hm *HashMap) Set(key string, value interface{}) {
	hm.Lock()
	defer hm.Unlock()

	hm.lenght++
	hm.keys = append(hm.keys, key)
	hm.values = append(hm.values, value)

	//Resize && Transfer
	if hm.loadFactor < hm.GetRealLoadFactor() {
		//fmt.Println("#####Resize && Transfer", hm.bucketCap, hm.GetRealLoadFactor())
		newBucketCap := hm.resize()
		hm.transfer(newBucketCap)
	}

	bucketIdx := hm.hashFunc(key, hm.bucketCap)
	hm.addLinkedList(key, value, bucketIdx, hm.lenght)
}

func (hm *HashMap) addLinkedList(key string, value interface{}, bucketIdx, idx uint32) {
	//fmt.Println("######AddLinkedList:", bucketIdx, hm.bucketCap)
	if hm.bucket[bucketIdx] == nil {
		hm.bucket[bucketIdx] = NewLinkedList()
	}

	hm.bucket[bucketIdx].Insert(key, value, idx)
}

func (hm *HashMap) resize() uint32 {
	newBucketCap := hm.bucketCap * 2
	hm.temp = make([]*LinkedList, newBucketCap)

	return newBucketCap
}

func (hm *HashMap) transfer(bucketCap uint32) {
	for idx, oldKey := range hm.keys {
		if oldKey == "" {
			continue
		}
		oldValue := hm.values[idx]
		bucketIdx := hm.hashFunc(oldKey, bucketCap)
		hm.addTempLinkedList(oldKey, oldValue, bucketIdx, uint32(idx))
	}

	hm.bucket = hm.temp
	hm.bucketCap = bucketCap
	hm.temp = nil
}

func (hm *HashMap) addTempLinkedList(key string, value interface{}, bucketIdx, idx uint32) {
	if hm.temp[bucketIdx] == nil {
		hm.temp[bucketIdx] = NewLinkedList()
	}

	hm.temp[bucketIdx].Insert(key, value, idx)
}

func (hm *HashMap) GetRealLoadFactor() float32 {

	return float32(hm.lenght) / float32(hm.bucketCap)
}

func (hm *HashMap) SetLoadFactor(loadFactor float32) *HashMap {
	if loadFactor > 0 || loadFactor < 1 {
		return hm
	}

	hm.loadFactor = loadFactor

	return hm
}
