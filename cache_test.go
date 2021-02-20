package common

import (
	"math/rand"
	"testing"
	"time"
)

func TestCacheAll(t *testing.T) {
	var maxCount int64 = 1000
	cache := NewLRUCache(maxCount)
	// init fill items
	for i := 0; i < int(maxCount); i++ {
		cache.Set(i, i*2)
		if cache.Len() != int64(i+1) {
			t.Error("check cache len failed", cache.Len(), i+1)
		}
	}

	for i := 0; i < int(maxCount); i++ {
		value, find := cache.Get(i)
		if !find {
			t.Error("not find this key", i)
		} else if value != i*2 {
			t.Error("check value failed", i*2, value)
		}
	}
	if cache.Len() != maxCount {
		t.Error("check cache len failed", cache.Len(), maxCount)
	}

	// update items
	for i := 0; i < int(maxCount); i++ {
		cache.Set(i, i*3)
	}

	for i := 0; i < int(maxCount); i++ {
		value, find := cache.Get(i)
		if !find {
			t.Error("not find this key", i)
		} else if value != i*3 {
			t.Error("check value failed", i*3, value)
		}
	}
	if cache.Len() != maxCount {
		t.Error("check cache len failed", cache.Len(), maxCount)
	}
	// set more washout all the old items
	for i := int(maxCount); i < int(maxCount*2); i++ {
		cache.Set(i, i*4)
		if cache.Len() != maxCount {
			t.Error("check cache len failed", cache.Len(), maxCount)
		}
	}

	for i := 0; i < int(maxCount); i++ {
		_, find := cache.Get(i)
		if find {
			t.Error("washout should not be getted")
		}
	}
	for i := int(maxCount); i < int(maxCount*2); i++ {
		value, find := cache.Get(i)
		if !find {
			t.Error("not find this key", i)
		} else if value != i*4 {
			t.Error("check value failed", i*4, value)
		}
		cache.Delete(i)
		value, find = cache.Get(i)
		if find {
			t.Error("should not find this key", i)
		}
	}

	// clear all
	cache.Clear()
	if cache.Len() != 0 {
		t.Error("check cache len failed", cache.Len())
	}
	for i := 0; i < int(maxCount*2); i++ {
		_, find := cache.Get(i)
		if find {
			t.Error("should not find this key", i)
		}
	}
}

func TestWashout(t *testing.T) {
	var maxCount int64 = 1000
	cache := NewLRUCache(maxCount)
	// init fill items
	for i := 0; i < int(maxCount); i++ {
		cache.Set(i, i*2)
		if cache.Len() != int64(i+1) {
			t.Error("check cache len failed", cache.Len(), i+1)
		}
	}

	rand.Seed(time.Now().UTC().UnixNano())
	keys := make([]int, 0)
	// get random then put more items
	for i := 0; i < 10*int(maxCount); i++ {
		key := rand.Intn(int(maxCount))
		keys = append(keys, key)
		value, find := cache.Get(key)
		if !find {
			t.Error("not find the item", key)
		} else if value != key*2 {
			t.Error("check value failed", key, value)
		}
	}

	// fill half new items
	for i := int(maxCount); i < int(maxCount)+int(maxCount)/2; i++ {
		cache.Set(i, i*5)
	}

	for i := int(maxCount); i < int(maxCount)+int(maxCount)/2; i++ {
		value, find := cache.Get(i)
		if !find {
			t.Error("not find the item", i)
		} else if value != i*5 {
			t.Error("check value failed", i, value)
		}
	}

	keys = keys[len(keys)-int(maxCount)/2 : len(keys)]
	// the last access must exist
	for i := 0; i < int(maxCount)/2; i++ {
		value, find := cache.Get(keys[i])
		if !find {
			t.Error("not find the item", keys[i])
		} else if value != keys[i]*2 {
			t.Error("check value failed", keys[i], value)
		}
	}
}
