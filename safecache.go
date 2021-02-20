package common

import (
	"sync"
)

type SafeCache struct {
	lock  sync.Mutex
	cache Cache
}

func NewSafeCache(cache Cache) *SafeCache {
	return &SafeCache{cache: cache}
}

func (this *SafeCache) Get(key interface{}) (interface{}, bool) {
	this.lock.Lock()
	defer this.lock.Unlock()
	return this.cache.Get(key)
}

func (this *SafeCache) Set(key interface{}, value interface{}) {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.cache.Set(key, value)
}

func (this *SafeCache) Delete(key interface{}) {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.cache.Delete(key)
}

func (this *SafeCache) Clear() {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.cache.Clear()
}

func (this *SafeCache) Len() int64 {
	this.lock.Lock()
	defer this.lock.Unlock()
	return this.cache.Len()
}

func (this *SafeCache) HitRatio() float32 {
	this.lock.Lock()
	defer this.lock.Unlock()
	return this.cache.HitRatio()
}
