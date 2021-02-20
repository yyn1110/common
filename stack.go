package common

import (
	"container/list"
)

// LIFO Stack interface
type Stack interface {
	Push(interface{})
	Pop()
	Top() interface{}
	Len() int64
	Empty() bool
	Clear()
}

// Use list.List as container
type ListStack struct {
	container *list.List
}

func NewListStack() *ListStack {
	return &ListStack{container: list.New()}
}

func (this *ListStack) Push(item interface{}) {
	this.container.PushBack(item)
}

func (this *ListStack) Pop() {
	if this.Empty() {
		return
	}
	this.container.Remove(this.container.Back())
}

func (this *ListStack) Top() interface{} {
	if this.Empty() {
		return nil
	}
	return this.container.Back().Value
}

func (this *ListStack) Len() int64 {
	return int64(this.container.Len())
}

func (this *ListStack) Empty() bool {
	return this.Len() == 0
}

func (this *ListStack) Clear() {
	if !this.Empty() {
		this.container = list.New()
	}
}
