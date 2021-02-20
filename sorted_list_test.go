package common

import (
	"math/rand"
	"testing"
	"time"
)

// compare two int type element
func compareInt(one interface{}, two interface{}) int {
	return one.(int) - two.(int)
}

func checkListSorted(list *SortedList) bool {
	var preValue int = -1 // min
	element := list.Front()
	for element != nil {
		if compareInt(preValue, element.Value) <= 0 {
			preValue = element.Value.(int)
			element = element.Next()
			continue
		}
		return false
	}
	return true
}

func TestSortedListInsert(t *testing.T) {
	list := NewSortedList(compareInt)
	set := NewSet()
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < maxCount; i++ {
		item := rand.Intn(maxCount)
		list.Insert(item)
		set.Add(item)
		if list.Len() != int64(i+1) {
			t.Error("check list len failed")
		}
		element := list.Search(item)
		if element == nil || compareInt(element.Value, item) != 0 {
			t.Error("search exist item failed", item)
		}
		if !checkListSorted(list) {
			t.Error("check the sorted list failed")
		}
	}
	// search it
	for i := 0; i < maxCount; i++ {
		if set.Contain(i) {
			element := list.Search(i)
			if element == nil || compareInt(element.Value, i) != 0 {
				t.Error("search exist item failed", i)
			}
		} else {
			element := list.Search(i)
			if element != nil {
				t.Error("find not inserted item", i)
			}
		}
	}
	if list.Len() != int64(maxCount) {
		t.Error("check list len failed", list.Len())
	}
}

func TestSortedListDeleteOne(t *testing.T) {
	list := NewSortedList(compareInt)
	keys := make([]int, 0)
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < maxCount; i++ {
		item := rand.Intn(maxCount)
		keys = append(keys, item)
		list.Insert(item)
		if list.Len() != int64(i+1) {
			t.Error("check list len failed")
		}
		if !checkListSorted(list) {
			t.Error("check the sorted list failed")
		}
	}
	for _, i := range keys {
		count := list.Contain(i)
		if count <= 0 {
			t.Error("check the item not exist failed")
		}
	}

	for _, i := range keys {
		succ := list.RemoveOne(i)
		if !succ {
			t.Error("delte exist item failed", i)
		} else if !checkListSorted(list) {
			t.Error("after delete check sorted list failed")
		}
	}

	if list.Len() != 0 {
		t.Error("check list len failed", list.Len())
	}
}

func TestSortedListDeleteAll(t *testing.T) {
	list := NewSortedList(compareInt)
	set := NewSet()
	for i := 0; i < maxCount; i++ {
		item := rand.Intn(maxCount)
		list.Insert(item)
		set.Add(item)
		if list.Len() != int64(i+1) {
			t.Error("check list len failed")
		}
		if !checkListSorted(list) {
			t.Error("check the sorted list failed")
		}
	}
	// delete all
	var total int64
	for i := 0; i < maxCount; i++ {
		if set.Contain(i) {
			count := list.RemoveAll(i)
			if count <= 0 {
				t.Error("check remove item failed")
			} else {
				total += count
			}
			if !checkListSorted(list) || int64(maxCount)-total != list.Len() {
				t.Error("check the sorted list or len failed")
			}
		}
	}
	if list.Len() != 0 || total != int64(maxCount) {
		t.Error("check list len failed", list.Len(), total)
	}
}
