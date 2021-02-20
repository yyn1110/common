package common

import (
	"testing"
)

// insert & delete & find
func TestIDF(t *testing.T) {
	hashmap := NewSafeMap()
	for i := 0; i < 100; i++ {
		err := hashmap.Insert(i, i+101)
		if err != nil {
			t.Error("insert new key failed", i, i+101)
		}
	}

	if hashmap.Len() != 100 {
		t.Error("check map len failed")
	}

	// Insert already exist
	for i := 0; i < 100; i++ {
		err := hashmap.Insert(i, i+201)
		if err == nil {
			t.Error("insert exist new key succ", i, i+201)
		}
	}

	// Find exist
	for i := 0; i < 100; i++ {
		v, find := hashmap.Find(i)
		if find != true {
			t.Error("not find the key", i)
		} else if v != i+101 {
			t.Error("check find value failed", i, v, i+101)
		}
	}

	// Find not exist
	for i := 100; i < 200; i++ {
		_, find := hashmap.Find(i)
		if find != false {
			t.Error("find not exist the key", i)
		}
	}

	// Delete not exist key
	for i := 100; i < 200; i++ {
		v, find := hashmap.Delete(i)
		if find != false {
			t.Error("delete not exist key, but return succ", i, v)
		}
	}

	// Delete exist key
	for i := 0; i < 100; i++ {
		v, find := hashmap.Delete(i)
		if find != true {
			t.Error("not find exist the key", i)
		} else if v != i+101 {
			t.Error("delete exist key，but check value failed", i, v)
		}
	}

	if hashmap.Len() != 0 {
		t.Error("check map not empty")
	}

	err := hashmap.Insert("test", "value")
	if err != nil {
		t.Error("insert string")
	}
	v, find := hashmap.Find("test")
	if find != true {
		t.Error("find the string value failed")
	} else if v != "value" {
		t.Error("check the string value failed")
	}
}

// replace & update
func TestUR(t *testing.T) {
	hashmap := NewSafeMap()
	for i := 0; i < 100; i++ {
		err := hashmap.Insert(i, i+101)
		if err != nil {
			t.Error("insert new key failed", i, i+101)
		}
	}
	// update exist
	for i := 0; i < 100; i++ {
		err := hashmap.Update(i, i+102)
		if err != nil {
			t.Error("update exist key failed", i, i+102)
		}
	}
	// check new value
	for i := 0; i < 100; i++ {
		v, find := hashmap.Find(i)
		if find != true {
			t.Error("not find the key", i)
		} else if v != i+102 {
			t.Error("check find value failed", i, v, i+101)
		}
	}

	// update not exist
	for i := 100; i < 200; i++ {
		err := hashmap.Update(i, i+102)
		if err == nil {
			t.Error("update not exist key failed", i, i+102)
		}
	}

	// replace exist
	for i := 0; i < 100; i++ {
		v, find := hashmap.Replace(i, i+101)
		if find != true {
			t.Error("not find the key", i)
		} else if v != i+102 {
			t.Error("check find value failed", i, v, i+101)
		}
	}
	// check new value
	for i := 0; i < 100; i++ {
		v, find := hashmap.Find(i)
		if find != true {
			t.Error("not find the key", i)
		} else if v != i+101 {
			t.Error("check find value failed", i, v, i+101)
		}
	}

	// replace not exist
	for i := 100; i < 200; i++ {
		v, find := hashmap.Replace(i, i+101)
		if find == true {
			t.Error("find the exist key", i, v)
		}
	}
	// check new value
	for i := 0; i < 200; i++ {
		v, find := hashmap.Find(i)
		if find != true {
			t.Error("not find the key", i)
		} else if v != i+101 {
			t.Error("check find value failed", i, v, i+101)
		}
	}
	if hashmap.Len() != 200 {
		t.Error("check hashmap len failed")
	}
}

func TestWalkerAll(t *testing.T) {
	array := make([]int, 0)
	walker := func(k, v interface{}) {
		array = append(array, k.(int))
	}

	hashmap := NewSafeMap()
	for i := 0; i < 100; i++ {
		err := hashmap.Insert(i, i+101)
		if err != nil {
			t.Error("insert new key failed", i, i+101)
		}
	}
	hashmap.Walk(walker)

	if len(array) != hashmap.Len() {
		t.Error("check array and hashmap len failed")
	}

	for _, key := range array {
		_, find := hashmap.Find(key)
		if !find {
			t.Error("item not find", key)
		}
	}
}

func TestWalkerWithLimit(t *testing.T) {
	array := make([]int, 0)
	walker := func(k, v interface{}) bool {
		array = append(array, k.(int))
		return true
	}

	hashmap := NewSafeMap()
	for i := 0; i < 100; i++ {
		err := hashmap.Insert(i, i+101)
		if err != nil {
			t.Error("insert new key failed", i, i+101)
		}
	}

	// limit == 50
	hashmap.WalkWithLimit(walker, 50)

	if len(array) != 50 {
		t.Error("check array and hashmap len failed")
	}

	for _, key := range array {
		_, find := hashmap.Find(key)
		if !find {
			t.Error("item not find", key)
		}
	}

	// limit == 0 (no limit)
	array = make([]int, 0)
	hashmap.WalkWithLimit(walker, 0)
	if len(array) != 100 {
		t.Error("check array and hashmap len failed")
	}
}
