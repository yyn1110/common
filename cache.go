package common

import (
	"container/list"
	"fmt"
)

type Cache interface {
	Get(key interface{}) (interface{}, bool)
	Set(key interface{}, value interface{})
	Delete(key interface{})
	Clear()
	Len() int64
	HitRatio() float32
}

type LRUCacheItem struct {
	value   interface{}
	element *list.Element
}

func NewLRUCacheItem(value interface{}, element *list.Element) *LRUCacheItem {
	return &LRUCacheItem{value: value, element: element}
}

type LRUCache struct {
	index    map[interface{}](*LRUCacheItem)
	lru      *list.List
	maxCount int64
	hit      int64
	miss     int64
}

func NewLRUCache(count int64) *LRUCache {
	return &LRUCache{index: make(map[interface{}](*LRUCacheItem)), lru: list.New(), maxCount: count}
}

func (this *LRUCache) Len() int64 {
	return int64(len(this.index))
}

func (this *LRUCache) HitRatio() float32 {
	if this.hit+this.miss > 0 {
		return float32(this.hit) / float32(this.hit+this.miss)
	}
	return 0.0
}

// only support enlarge the max count rightnow
func (this *LRUCache) SetMaxCount(count int64) {
	if this.maxCount < count {
		this.maxCount = count
	}
}

// if exist return value + true, else return nil + false
func (this *LRUCache) Get(key interface{}) (interface{}, bool) {
	item, find := this.index[key]
	if find {
		this.hit++
		this.lru.MoveToFront(item.element)
		return item.value, true
	}
	this.miss++
	return nil, false
}

func (this *LRUCache) Set(key interface{}, value interface{}) {
	item, find := this.index[key]
	if find {
		item.value = value
		this.lru.MoveToFront(item.element)
		return
	} else if this.Len() >= this.maxCount {
		// TODO remove more than one items
		delete(this.index, this.lru.Remove(this.lru.Back()))
	}
	this.index[key] = NewLRUCacheItem(value, this.lru.PushFront(key))
}

func (this *LRUCache) Delete(key interface{}) {
	item, find := this.index[key]
	if find {
		this.lru.Remove(item.element)
		delete(this.index, key)
	}
}

func (this *LRUCache) Clear() {
	if this.Len() > 0 {
		this.index = make(map[interface{}](*LRUCacheItem))
		this.lru = list.New()
	}
}

func (this *LRUCache) Debug() {
	fmt.Printf("cache length:<%d %d>, items:", this.lru.Len(), len(this.index))
	element := this.lru.Front()
	for element != nil {
		value, find := this.index[element.Value]
		if !find {
			panic("error")
		}
		fmt.Print("<", element.Value, value.value, ">")
		element = element.Next()
	}
	fmt.Println()
}
