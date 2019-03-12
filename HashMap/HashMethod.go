package HashMap

import (
	"hash/crc32"
)

func HashCode(key string, hashLen uint32) uint32 {

	return crc32.ChecksumIEEE([]byte(key)) % hashLen
}

func (hm *HashMap) SetHashCode(hashCode func(key string, hashLen uint32) uint32) *HashMap {
	hm.hashFunc = hashCode

	return hm
}
