package common

import (
	"fmt"
	"strings"
)

type Set struct {
	data map[interface{}]struct{}
}

func NewSet() *Set {
	return &Set{data: make(map[interface{}]struct{})}
}

// if not exist return true
func (this *Set) Add(elem interface{}) bool {
	_, find := this.data[elem]
	if !find {
		this.data[elem] = struct{}{}
	}
	return !find
}

// if exist return true
func (this *Set) Contain(elem interface{}) bool {
	_, find := this.data[elem]
	return find
}

// if exist return true
func (this *Set) Remove(elem interface{}) bool {
	_, find := this.data[elem]
	if find {
		delete(this.data, elem)
	}
	return find
}

func (this *Set) Clear() {
	if this.Len() > 0 {
		this.data = make(map[interface{}]struct{})
	}
}

func (this *Set) Len() int64 {
	return int64(len(this.data))
}

func (this *Set) String() string {
	items := make([]string, 0, this.Len())
	for elem := range this.data {
		items = append(items, fmt.Sprintf("%v", elem))
	}
	return fmt.Sprintf("Set{%s}", strings.Join(items, ", "))
}

func (this *Set) IsSubset(other *Set) bool {
	for elem := range this.data {
		if !other.Contain(elem) {
			return false
		}
	}
	return true
}

func (this *Set) Union(other *Set) *Set {
	set := NewSet()
	for elem := range other.data {
		set.Add(elem)
	}
	for elem := range this.data {
		set.Add(elem)
	}
	return set
}

func (this *Set) Intersect(other *Set) *Set {
	set := NewSet()
	var smaller, bigger *Set
	if this.Len() > other.Len() {
		smaller = other
		bigger = this
	} else {
		smaller = this
		bigger = other
	}
	for elem := range smaller.data {
		if bigger.Contain(elem) {
			set.Add(elem)
		}
	}
	return set
}

func (this *Set) Difference(other *Set) *Set {
	set := NewSet()
	for elem := range other.data {
		if !this.Contain(elem) {
			set.Add(elem)
		}
	}
	for elem := range this.data {
		if !other.Contain(elem) {
			set.Add(elem)
		}
	}
	return set
}

// walk through all value then callback
func (this *Set) Walk(callback func(elem interface{})) {
	for elem := range this.data {
		callback(elem)
	}
}
