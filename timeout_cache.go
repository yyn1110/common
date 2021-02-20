package common

import (
	"time"
)

type TimeoutValue struct {
	value     interface{}
	timestamp int64
}

func NewTimeoutValue(value interface{}) *TimeoutValue {
	return &TimeoutValue{value: value, timestamp: time.Now().UTC().UnixNano()}
}

func (this *TimeoutValue) Expired(timeout int64) bool {
	return this.timestamp+timeout < time.Now().UTC().UnixNano()
}

// Lazy washout the timeout entries
type TimeoutCache struct {
	cache   *LRUCache
	timeout int64
}

func NewTimeoutCache(count, timeout int64) *TimeoutCache {
	return &TimeoutCache{cache: NewLRUCache(count), timeout: timeout}
}

func (this *TimeoutCache) Get(key interface{}) (interface{}, bool) {
	item, find := this.cache.Get(key)
	if find && !item.(*TimeoutValue).Expired(this.timeout) {
		return item.(*TimeoutValue).value, true
	} else if find {
		this.cache.Delete(key)
	}
	return nil, false
}

func (this *TimeoutCache) Set(key interface{}, value interface{}) {
	this.cache.Set(key, NewTimeoutValue(value))
}

func (this *TimeoutCache) Delete(key interface{}) {
	this.cache.Delete(key)
}

func (this *TimeoutCache) Clear() {
	this.cache.Clear()
}

// WARNING not exclude all the timeout entries
func (this *TimeoutCache) Len() int64 {
	return this.cache.Len()
}

func (this *TimeoutCache) HitRatio() float32 {
	return this.cache.HitRatio()
}