package common

import (
	"testing"
)

func testingAllStack(stack Stack, t *testing.T) {
	var maxCount int64 = 1000
	if stack.Len() != 0 || !stack.Empty() {
		t.Error("check stack len failed")
	}
	for i := 0; i < int(maxCount); i++ {
		stack.Push(i)
	}
	if stack.Len() != maxCount || stack.Empty() {
		t.Error("check stack len failed")
	}

	for i := 0; i < int(maxCount); i++ {
		v := stack.Top().(int)
		if v != int(maxCount)-i-1 {
			t.Error("check pop value failed", v)
		}
		stack.Pop()
	}
	if stack.Top() != nil || stack.Len() != 0 || !stack.Empty() {
		t.Error("check stack len failed")
	}

	// pop nothing
	for i := 0; i < int(maxCount); i++ {
		stack.Pop()
	}
	if stack.Top() != nil || stack.Len() != 0 || !stack.Empty() {
		t.Error("check stack len failed")
	}
}

func TestListStack(t *testing.T) {
	stack := NewListStack()
	testingAllStack(stack, t)
}
