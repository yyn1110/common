package common

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func checkSkipListSorted(list *SkipList) bool {
	for i := 0; i < MAX_HEIGHT; i++ {
		element := list.Head().next[i]
		var preValue int = -1 // min
		for element != nil && element.next[i] != nil {
			if compareInt(preValue, element.next[i].value) < 0 {
				preValue = element.next[i].value.(int)
				element = element.Next(i)
				continue
			}
			fmt.Printf("level:%d, pre:%d, cur:%d\n", i, preValue, element.next[i].value)
			return false
		}
	}
	return true
}

func TestSkipListSetGet(t *testing.T) {
	list := NewSkipList(compareInt)
	set := NewSet()
	rand.Seed(time.Now().UTC().UnixNano())
	var count int
	for count < maxCount {
		curCount := rand.Intn(maxCount)
		for i := 0; i < curCount; i++ {
			item := rand.Intn(curCount)
			list.Set(item, item*2)
			set.Add(item)
			element, find := list.Get(item)
			if !find || element == nil || compareInt(element.(int), item*2) != 0 {
				t.Error("search exist item failed", item)
				return
			}
			if !checkSkipListSorted(list) {
				list.printData(0)
				t.Error("check the sorted list failed")
			}
		}
		// key value
		if list.Len() != set.Len() {
			t.Error("check list len failed", set.Len(), list.Len())
		}
		// check values
		for i := 0; i < curCount; i++ {
			if set.Contain(i) {
				element, find := list.Get(i)
				if !find || element != 2*i {
					t.Error("get exist key value pair failed", find, element, i*2)
				}
				list.Set(i, i*3)
				element, find = list.Get(i)
				if !find || element != 3*i {
					t.Error("get exist updated key value pair failed", find, element, i*3)
				}
			} else {
				_, find := list.Get(i)
				if find {
					t.Error("should not find this key")
				}
			}
		}
		list.Clear()
		set.Clear()
		count++
	}
}

func TestSkipListDelete(t *testing.T) {
	list := NewSkipList(compareInt)
	set := NewSet()
	rand.Seed(time.Now().UTC().UnixNano())
	var maxCount int = 2000
	var count int
	for count < maxCount {
		// fill items
		curCount := rand.Intn(maxCount)
		for i := 0; i < curCount; i++ {
			item := rand.Intn(curCount)
			list.Set(item, item*2)
			set.Add(item)
			element, find := list.Get(item)
			if !find || element == nil || compareInt(element.(int), item*2) != 0 {
				t.Error("search exist item failed", item)
				return
			}
			if !checkSkipListSorted(list) {
				list.printData(0)
				t.Error("check the sorted list failed")
			}
		}
		for i := 0; i < curCount; i++ {
			value, find := list.Delete(i)
			if set.Contain(i) {
				if value == nil || value != i*2 || !find {
					t.Error("check value failed", i, value)
				}
			} else if find || value != nil {
				t.Error("delete not exist key succ", i, value)
			}
			if !checkSkipListSorted(list) {
				list.printData(0)
				t.Error("check the sorted list failed")
			}
		}
		if list.Len() != 0 {
			t.Error("check list len failed", list.Len())
		}
		list.Clear()
		set.Clear()
		count++
	}
}
