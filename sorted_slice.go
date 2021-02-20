package common

import (
	"fmt"
)

type SortedSlice struct {
	container []interface{}
	compare   CompareFunc
}

func NewSortedSlice(reserve int64, compare CompareFunc) *SortedSlice {
	return &SortedSlice{container: make([]interface{}, 0, reserve), compare: compare}
}

func (this *SortedSlice) Len() int64 {
	return int64(len(this.container))
}

func (this *SortedSlice) Insert(item interface{}) {
	pos := this.binarySearch(this.compare, item, this.container, 0, this.Len()-1)
	// fmt.Printf("pos[%d], len[%d], item[%d]\n", pos, this.Len(), item)
	if pos == this.Len() {
		this.container = append(this.container, item)
	} else {
		if this.Len() == int64(cap(this.container)) {
			temp := make([]interface{}, 0, cap(this.container)<<2)
			temp = append(temp, this.container[:pos]...)
			temp = append(temp, item)
			temp = append(temp, this.container[pos:]...)
			this.container = temp
		} else {
			len := this.Len()
			this.container = append(this.container, item)
			for i := len; i > pos; i-- {
				this.container[i] = this.container[i-1]
			}
			this.container[pos] = item
		}
	}
	// this.Debug()
}

func (this *SortedSlice) At(pos int64) interface{} {
	if pos < 0 || pos >= this.Len() {
		return nil
	}
	return this.container[pos]
}

func (this *SortedSlice) Clear() {
	if this.Len() <= 0 {
		return
	}
	this.container = make([]interface{}, 0, cap(this.container))
}

// if find return first pos, if not exist return -1
func (this *SortedSlice) Search(item interface{}) int64 {
	pos := this.binarySearch(this.compare, item, this.container, 0, this.Len()-1)
	if this.At(pos) == item {
		for pos >= 0 && this.At(pos) == item {
			pos--
		}
		return pos + 1
	}
	return -1
}

// not the frist one only the right pos for insert, return pos between [0, len]
func (this *SortedSlice) binarySearch(compare CompareFunc, item interface{}, array []interface{}, start, end int64) int64 {
	for start <= end {
		mid := (start + end) >> 1
		result := compare(item, array[mid])
		if result == 0 {
			return mid
		} else if result > 0 {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return start
}

func (this *SortedSlice) Remove(pos int64) interface{} {
	item := this.At(pos)
	if pos >= 0 && pos < this.Len()-1 {
		copy(this.container[pos:this.Len()-1], this.container[pos+1:this.Len()])
		this.container = this.container[:this.Len()-1]
		// or this.container = append(this.container[:pos], this.container[pos+1:])
	} else if pos == this.Len()-1 {
		this.container = this.container[:pos]
	}
	return item
}

// return item count
func (this *SortedSlice) Contain(item interface{}) int64 {
	pos := this.binarySearch(this.compare, item, this.container, 0, this.Len()-1)
	if pos >= this.Len() || this.compare(item, this.At(pos)) != 0 {
		return 0
	}
	var result int
	var count int64 = 1
	// forward
	index := pos - 1
	for index >= 0 {
		result = this.compare(item, this.At(index))
		if result == 0 {
			count++
			index--
		} else {
			break
		}
	}
	// backword
	index = pos + 1
	for index < this.Len() {
		result = this.compare(item, this.At(index))
		if result == 0 {
			count++
			index++
		} else {
			break
		}
	}
	return count
}

func (this *SortedSlice) Debug() {
	fmt.Printf("slice length:<%d>, content:", this.Len())
	for _, i := range this.container {
		fmt.Print("<", i, ">")
	}
	fmt.Println()
}
