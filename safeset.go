package common

import (
	"sync"
)

type SafeSet struct {
	lock      sync.RWMutex
	container *Set
}

func NewSafeSet() *SafeSet {
	return &SafeSet{container: NewSet()}
}

func (this *SafeSet) Add(elem interface{}) bool {
	this.lock.Lock()
	defer this.lock.Unlock()
	return this.container.Add(elem)
}

func (this *SafeSet) Remove(elem interface{}) bool {
	this.lock.Lock()
	defer this.lock.Unlock()
	return this.container.Remove(elem)
}

func (this *SafeSet) Contain(elem interface{}) bool {
	this.lock.RLock()
	defer this.lock.RUnlock()
	return this.container.Contain(elem)
}

func (this *SafeSet) Len() int64 {
	this.lock.RLock()
	defer this.lock.RUnlock()
	return this.container.Len()
}

func (this *SafeSet) String() string {
	this.lock.RLock()
	defer this.lock.RUnlock()
	return this.container.String()
}

func (this *SafeSet) Walk(callback func(elem interface{})) {
	this.lock.RLock()
	defer this.lock.RUnlock()
	this.container.Walk(callback)
}

func (this *SafeSet) Clear() {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.container.Clear()
}