package common

import (
	"math/rand"
	"testing"
	"time"
)

func TestTimeoutCache(t *testing.T) {
	var maxCount int64 = 100
	var timeout int64 = int64(time.Second)
	cache := NewTimeoutCache(maxCount, timeout)
	// fill the items
	for i := 0; i < int(maxCount); i++ {
		cache.Set(i, i*3)
	}
	for i := 0; i < int(maxCount); i++ {
		random := rand.Intn(int(maxCount))
		value, find := cache.Get(random)
		if !find {
			t.Error("not find the cache item", i)
		} else if value != random*3 {
			t.Error("check value failed", value, i*3)
		}
	}
	if cache.Len() != maxCount {
		t.Error("check len failed")
	}

	// wait for real timeout
	time.Sleep(time.Second * 2)
	// not reset the len before real washout
	if cache.Len() != maxCount {
		t.Error("check len failed")
	}

	// get all for lazy washout
	for i := 0; i < int(maxCount); i++ {
		_, find := cache.Get(i)
		if find {
			t.Error("should not get the expired item", i)
		}
	}
	if cache.Len() != 0 {
		t.Error("check len failed", cache.Len())
	}

	// update item with new timestamp
	for i := 0; i < int(maxCount); i++ {
		cache.Set(i, i*3)
	}
	time.Sleep(2 * time.Second)
	if cache.Len() != maxCount {
		t.Error("check len failed")
	}
	for i := 0; i < int(maxCount); i++ {
		cache.Set(i, i*4)
	}
	if cache.Len() != maxCount {
		t.Error("check len failed")
	}
	for i := 0; i < int(maxCount); i++ {
		value, find := cache.Get(i)
		if !find {
			t.Error("not get the cache item", i)
		} else if value != i*4 {
			t.Error("check value failed", value, i*4)
		}
	}

	// delete all
	for i := 0; i < 2*int(maxCount); i++ {
		cache.Delete(i)
	}

	if cache.Len() != 0 {
		t.Error("check len failed")
	}
}
