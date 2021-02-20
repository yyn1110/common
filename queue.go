package common

import (
	"container/list"
	"fmt"
)

// FIFO Stack interface
type Queue interface {
	Push(interface{})
	Pop()
	Back() interface{}
	Front() interface{}
	Empty() bool
	Len() int64
	Clear()
}

// Use list.List as container
type ListQueue struct {
	container *list.List
}

func NewListQueue() *ListQueue {
	return &ListQueue{container: list.New()}
}

func (this *ListQueue) Len() int64 {
	return int64(this.container.Len())
}

func (this *ListQueue) Empty() bool {
	return this.Len() == 0
}

func (this *ListQueue) Push(item interface{}) {
	this.container.PushBack(item)
}

func (this *ListQueue) Pop() {
	if !this.Empty() {
		this.container.Remove(this.container.Front())
	}
}

func (this *ListQueue) Clear() {
	if !this.Empty() {
		this.container = list.New()
	}
}

func (this *ListQueue) Front() interface{} {
	if !this.Empty() {
		return this.container.Front().Value
	}
	return nil
}

func (this *ListQueue) Back() interface{} {
	if !this.Empty() {
		return this.container.Back().Value
	}
	return nil
}

func (this *ListQueue) Print() {
	fmt.Print(this.container)
}
