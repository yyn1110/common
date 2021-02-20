package common

import (
	"container/list"
)

type CompareFunc func(interface{}, interface{}) int

type SortedList struct {
	container *list.List
	compare   CompareFunc
}

func NewSortedList(compare CompareFunc) *SortedList {
	return &SortedList{container: list.New(), compare: compare}
}

func (this *SortedList) Len() int64 {
	return int64(this.container.Len())
}

func (this *SortedList) Front() *list.Element {
	return this.container.Front()
}

func (this *SortedList) Back() *list.Element {
	return this.container.Back()
}

func (this *SortedList) Clear() {
	if this.Len() > 0 {
		this.container = list.New()
	}
}

// Only Search from front to end, optimized for from end to front
func (this *SortedList) Insert(item interface{}) {
	element := this.container.Front()
	for element != nil {
		if this.compare(item, element.Value) <= 0 {
			this.container.InsertBefore(item, element)
			return
		}
		element = element.Next()
	}
	if element == nil {
		this.container.PushBack(item)
	}
}

// WARNING can not modify the element value
func (this *SortedList) Search(item interface{}) *list.Element {
	element := this.container.Front()
	var result int
	for element != nil {
		result = this.compare(item, element.Value)
		if result < 0 {
			return nil
		} else if result == 0 {
			return element
		}
		element = element.Next()
	}
	return nil
}

// find all item counter
func (this *SortedList) Contain(item interface{}) int64 {
	element := this.container.Front()
	var result int
	var counter int64
	for element != nil {
		result = this.compare(item, element.Value)
		if result < 0 {
			break
		} else if result == 0 {
			counter++
		}
		element = element.Next()
	}
	return counter
}

// remove the first equal with the item, if find return succ
func (this *SortedList) RemoveOne(item interface{}) bool {
	var result int
	element := this.container.Front()
	for element != nil {
		result = this.compare(item, element.Value)
		if result < 0 {
			return false
		} else if result == 0 {
			this.container.Remove(element)
			return true
		}
		element = element.Next()
	}
	return false
}

// remove all the items equal with the item, return removed count
func (this *SortedList) RemoveAll(item interface{}) int64 {
	element := this.container.Front()
	var result int
	var counter int64
	var temp *list.Element
	for element != nil {
		result = this.compare(item, element.Value)
		if result < 0 {
			break
		} else if result == 0 {
			temp = element.Next()
			this.container.Remove(element)
			counter++
			element = temp
			continue
		}
		element = element.Next()
	}
	return counter
}
