package common

import (
	"time"
)

type TTLValue struct {
	value		interface{}
	timestamp	int64
}

func NewTTLValue(value interface{}) *TTLValue {
	return &TTLValue{value: value, timestamp: time.Now().UTC().UnixNano()}
}

func (this *TTLValue) Expired(ttl int64) bool {
	return this.timestamp + ttl < time.Now().UTC().UnixNano()
}

type TTLCache struct {
	cache	*LRUCache
	ttl		int64
}

func NewTTLCache(count, ttl int64) *TTLCache {
	return &TTLCache{cache: NewLRUCache(count), ttl: ttl}
}

// return cache item value, found or not, expired or not
func (this *TTLCache) Get(key interface{}) (interface{}, bool, bool) {
	item, find := this.cache.Get(key)
	if find && !item.(*TTLValue).Expired(this.ttl) {
		// cache found and not expired
		return item.(*TTLValue).value, true, false
	} else if find {
		// cache found but expired
		// do not delete the cache item in ttl cache, let outside business logic decide
		return item.(*TTLValue).value, true, true
	}
	return nil, false, false
}

func (this *TTLCache) Set(key interface{}, value interface{}) {
	this.cache.Set(key, NewTTLValue(value))
}

func (this *TTLCache) UpdateTTL(key interface{}) {
	item, find := this.cache.Get(key)
	if find {
		item.(*TTLValue).timestamp = time.Now().UTC().UnixNano()
	}
}

func (this *TTLCache) Delete(key interface{}) {
	this.cache.Delete(key)
}

func (this *TTLCache) Clear() {
	this.cache.Clear()
}

func (this *TTLCache) Len() int64 {
	return this.cache.Len()
}

func (this *TTLCache) HitRatio() float32 {
	return this.cache.HitRatio()
}
