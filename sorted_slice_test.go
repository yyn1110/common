package common

import (
	"math/rand"
	"testing"
	"time"
)

func checkSliceSorted(slice *SortedSlice) bool {
	var preValue int = -1 // min
	for i := 0; i < int(slice.Len()); i++ {
		if slice.At(int64(i)).(int) < preValue {
			return false
		}
		preValue = slice.At(int64(i)).(int)
	}
	return true
}

func compareListArray(slice *SortedSlice, list *SortedList) bool {
	var i int64
	element := list.Front()
	if slice.Len() != list.Len() {
		return false
	}
	for element != nil {
		if slice.At(int64(i)).(int) != element.Value.(int) {
			return false
		}
		element = element.Next()
		i++
	}
	return true
}

const maxCount int = 300

func TestSortedSliceInsert(t *testing.T) {
	array := NewSortedSlice(int64(maxCount/5), compareInt)
	list := NewSortedList(compareInt)
	set := NewSet()
	rand.Seed(time.Now().UTC().UnixNano())
	var count int
	for count < maxCount {
		for i := 0; i < maxCount; i++ {
			item := rand.Intn(maxCount)
			array.Insert(item)
			list.Insert(item)
			set.Add(item)
			if array.Len() != int64(i+1) {
				t.Error("check list len failed", array.Len(), i+1)
			}
			if !checkSliceSorted(array) || !compareListArray(array, list) {
				t.Error("check the sorted list failed")
			}
			itemCount := array.Search(item)
			if itemCount < 0 {
				t.Error("search the exist item failed", item)
			}
		}
		array.Clear()
		list.Clear()
		set.Clear()
		if array.Len() != 0 {
			t.Error("after clear check len failed", array.Len())
		}
		count++
	}
}

func TestSortedSliceSearch(t *testing.T) {
	array := NewSortedSlice(0, compareInt)
	list := NewSortedList(compareInt)
	set := NewSet()
	rand.Seed(time.Now().UTC().UnixNano())
	var count int
	for count < maxCount {
		curCount := rand.Intn(maxCount)
		for i := 0; i < curCount; i++ {
			item := rand.Intn(curCount)
			array.Insert(item)
			list.Insert(item)
			set.Add(item)
			if array.Len() != int64(i+1) {
				t.Error("check list len failed", array.Len(), i+1)
			}
			if !checkSliceSorted(array) || !compareListArray(array, list) {
				t.Error("check the sorted list failed")
			}
			itemCount := array.Search(item)
			if itemCount < 0 {
				t.Error("search the exist item failed", item)
			}
		}
		// search the item
		var total int64
		for i := 0; i < curCount; i++ {
			find := set.Contain(i)
			pos := array.Search(i)
			count := array.Contain(i)
			if find {
				if pos < 0 || count < 0 || array.At(pos).(int) != i {
					t.Error("exist search failed", i, pos)
				}
				if pos > 0 && array.At(pos-1).(int) == i {
					t.Error("return not the frist item", i, pos)
				}
			} else if pos > 0 || count != 0 {
				t.Error("not exist search ok")
			}
			total += count
		}
		if total != int64(curCount) {
			t.Error("check total count failed", total, curCount)
		}
		array.Clear()
		set.Clear()
		list.Clear()
		count++
	}
}

func TestSortedSliceRemove(t *testing.T) {
	array := NewSortedSlice(int64(maxCount/5), compareInt)
	list := NewSortedList(compareInt)
	set := NewSet()
	rand.Seed(time.Now().UTC().UnixNano())
	var count int
	for count < maxCount {
		curCount := rand.Intn(maxCount)
		for i := 0; i < curCount; i++ {
			item := rand.Intn(curCount)
			array.Insert(item)
			list.Insert(item)
			set.Add(item)
			if array.Len() != int64(i+1) {
				t.Error("check list len failed", array.Len(), i+1)
			}
		}
		if !checkSliceSorted(array) || !compareListArray(array, list) {
			t.Error("check the sorted list failed")
		}

		// remove items
		var total int
		for array.Len() > 0 {
			pos := rand.Intn(int(array.Len()))
			value := array.Remove(int64(pos))
			if value == nil {
				t.Error("remove item failed")
			}
			if !checkSliceSorted(array) {
				t.Error("check the sorted list failed")
			}
			total++
			if array.Len() != int64(curCount-total) {
				t.Error("check count failed")
			}
		}

		if array.Len() != 0 {
			t.Error("check array len failed", array.Len())
		}
		value := array.At(0)
		if value != nil {
			t.Error("should not return value")
		}
		// nothing to remove
		value = array.Remove(0)
		if value != nil {
			t.Error("should return nothing")
		}

		array.Clear()
		set.Clear()
		list.Clear()
		count++
	}
}
