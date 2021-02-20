package common

import (
	"fmt"
	"sync"
	"testing"
)

func TestSafeCache(t *testing.T) {
	cache := NewLRUCache(100)
	scache := NewSafeCache(cache)

	var wait sync.WaitGroup
	var routineCount int = 30
	wait.Add(routineCount)
	routine := func(start, end int) {
		for i := start; i < end; i++ {
			key := fmt.Sprintf("key%d", i)
			value := fmt.Sprintf("value%d", i)
			scache.Set(key, value)
		}
		wait.Done()
	}
	for i := 0; i < routineCount; i++ {
		go routine(0, 100)
	}
	wait.Wait()
	// check the cache infos
	// fmt.Println("wait the cache item failed")
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("key%d", i)
		value, find := scache.Get(key)
		if !find {
			t.Error("get cache item failed", key)
		} else if value != fmt.Sprintf("value%d", i) {
			t.Error("check value failed", key, value)
		}
	}
}
