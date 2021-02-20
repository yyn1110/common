package common

import (
	"sync"
)

type CacheTTL interface {
	Get(key interface{}) (interface{}, bool, bool)
	Set(key interface{}, value interface{})
	Delete(key interface{})
	Clear()
	Len() int64
	HitRatio() float32
}

type SafeTTLCache struct {
	lock  sync.Mutex
	cache CacheTTL
}

func NewSafeTTLCache(cache CacheTTL) *SafeTTLCache {
	return &SafeTTLCache{cache: cache}
}

func (this *SafeTTLCache) Get(key interface{}) (interface{}, bool, bool) {
	this.lock.Lock()
	defer this.lock.Unlock()
	return this.cache.Get(key)
}

func (this *SafeTTLCache) Set(key interface{}, value interface{}) {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.cache.Set(key, value)
}

func (this *SafeTTLCache) Delete(key interface{}) {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.cache.Delete(key)
}

func (this *SafeTTLCache) Clear() {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.cache.Clear()
}

func (this *SafeTTLCache) Len() int64 {
	this.lock.Lock()
	defer this.lock.Unlock()
	return this.cache.Len()
}

func (this *SafeTTLCache) HitRatio() float32 {
	this.lock.Lock()
	defer this.lock.Unlock()
	return this.cache.HitRatio()
}
