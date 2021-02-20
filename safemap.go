package common

import (
	"sync"
)

type SafeMap struct {
	lock      sync.RWMutex
	container map[interface{}](interface{})
}

func NewSafeMap() *SafeMap {
	conn := make(map[interface{}](interface{}))
	if conn != nil {
		return &SafeMap{container: conn}
	}
	return nil
}

// all kv pair count
func (this *SafeMap) Len() int {
	this.lock.RLock()
	l := len(this.container)
	this.lock.RUnlock()
	return l
}

// if key exist update, else insert
func (this *SafeMap) Replace(key interface{}, value interface{}) (interface{}, bool) {
	this.lock.Lock()
	temp, find := this.container[key]
	this.container[key] = value
	this.lock.Unlock()
	return temp, find
}

// if key exist return err, else return nil
func (this *SafeMap) Insert(key interface{}, value interface{}) error {
	this.lock.Lock()
	defer this.lock.Unlock()
	_, find := this.container[key]
	if find {
		return ErrEntryExist
	}
	this.container[key] = value
	return nil
}

// if key exist return nil, else return err
func (this *SafeMap) Update(key interface{}, value interface{}) error {
	this.lock.Lock()
	defer this.lock.Unlock()
	_, find := this.container[key]
	if !find {
		return ErrEntryNotExist
	}
	this.container[key] = value
	return nil
}

// if key exist return value + true, else return nil + false
func (this *SafeMap) Delete(key interface{}) (interface{}, bool) {
	this.lock.Lock()
	value, find := this.container[key]
	delete(this.container, key)
	this.lock.Unlock()
	return value, find
}

// if key exist return value + true, else return nil + false
func (this *SafeMap) Find(key interface{}) (interface{}, bool) {
	this.lock.RLock()
	value, find := this.container[key]
	this.lock.RUnlock()
	return value, find
}

// walk through all the key value node then callback
func (this *SafeMap) Walk(callback func(k, v interface{})) {
	this.lock.RLock()
	defer this.lock.RUnlock()
	for key, value := range this.container {
		callback(key, value)
	}
}

// 扫描 map 并执行 callback；当返回值等于 true 时，将其算作是有效的 limit 限制增一。
// Note:
//  当 limit == 0 时表示不进行限制。
func (this *SafeMap) WalkWithLimit(callback func(k, v interface{}) bool, limit int) {
	this.lock.RLock()
	defer this.lock.RUnlock()
	i := 0
	for key, value := range this.container {
		if limit == 0 || i < limit {
			if callback(key, value) == true {
				i += 1
			}
		} else {
			break
		}
	}
}

// clear all the data
func (this *SafeMap) Clear() {
	this.lock.Lock()
	if len(this.container) > 0 {
		this.container = make(map[interface{}](interface{}))
	}
	this.lock.Unlock()
}
