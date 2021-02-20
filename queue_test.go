package common

import (
	"testing"
)

func testingAllQueue(queue Queue, t *testing.T) {
	var maxCount int64 = 1000
	if queue.Len() != 0 || !queue.Empty() {
		t.Error("check queue len failed")
	}
	// fill all items
	for i := 0; i < int(maxCount); i++ {
		queue.Push(i)
		front := queue.Front()
		if front != 0 {
			t.Error("check front failed")
		}
		back := queue.Back()
		if back != i {
			t.Error("check back failed")
		}
	}
	if queue.Len() != maxCount || queue.Empty() {
		t.Error("check queue len failed")
	}

	// pop all
	for i := 0; i < int(maxCount); i++ {
		front := queue.Front()
		if front != i {
			t.Error("check front failed")
		}
		queue.Pop()
	}
	if queue.Len() != 0 || !queue.Empty() {
		t.Error("check queue len failed")
	}

	if queue.Front() != nil || queue.Back() != nil {
		t.Error("no item as top")
	}
	// pop nothing
	for i := 0; i < int(maxCount); i++ {
		queue.Pop()
	}
	if queue.Len() != 0 || !queue.Empty() {
		t.Error("check queue len failed")
	}
	if queue.Front() != nil || queue.Back() != nil {
		t.Error("no item as top")
	}
}

func TestListQueue(t *testing.T) {
	queue := NewListQueue()
	testingAllQueue(queue, t)
}
