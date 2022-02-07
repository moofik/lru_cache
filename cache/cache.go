package cache

import (
	"encoding/binary"
	heap "lrucache/heap"
	"time"
)

func WithInt64Cache(fn func(key int64) interface{}) (func(key int64) interface{}, func()) {
	cache := make(map[int64]interface{})
	hp := heap.CreateHeap(100)

	return func(key int64) interface{} {
			currentTime := time.Now().Nanosecond()
			if el, ok := cache[key]; ok {
				bs := make([]byte, 8)
				binary.BigEndian.PutUint64(bs, uint64(key))
				hp.UpdateKeyPriority(bs, int64(currentTime))
				return el
			}

			result := fn(key)

			if len(cache) == 100 {
				minKey := hp.ExtractMin()
				delete(cache, int64(binary.BigEndian.Uint64(minKey)))
			}

			cache[key] = result
			bs := make([]byte, 8)
			binary.BigEndian.PutUint64(bs, uint64(key))
			hp.Insert(bs, int64(currentTime))

			return result
		}, func() {
			hp.Dump()
		}
}
