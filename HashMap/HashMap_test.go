package HashMap

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type Rand struct {
	*rand.Rand
}

func NewRand() *Rand {
	rand := rand.New(rand.NewSource(time.Now().Unix()))
	return &Rand{
		Rand: rand,
	}
}

// RandString 生成随机字符串
func (r *Rand) RandString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

func TestHashMapCase(t *testing.T) {
	hashMap := NewHashMap()
	randObj := NewRand()

	maxLen := 40
	testCase := make([]string, maxLen)

	for idx := 0; idx < maxLen; idx++ {

		testCase[idx] = randObj.RandString(2)
	}

	fmt.Println("testCase:", testCase)

	for idx, key := range testCase {
		hashMap.Set(key, idx)
	}

	for _, key := range testCase {
		value, ok := hashMap.Get(key)
		fmt.Println("key:", key, value, ok)
	}

	fmt.Println(hashMap.keys)
	fmt.Println(hashMap.values)
}
