package common

import (
	"fmt"
	"math/rand"
)

type SkipListItem struct {
	key   interface{}     // list node item key
	value interface{}     // list node item value
	next  []*SkipListItem // every level next node
}

func NewSkipListItem(height int) *SkipListItem {
	// fill nil with height length
	return &SkipListItem{next: make([]*SkipListItem, height)}
}

func (this *SkipListItem) Next(height int) *SkipListItem {
	return this.next[height]
}

func (this *SkipListItem) SetNext(height int, item *SkipListItem) {
	this.next[height] = item
}

const MAX_HEIGHT int = 12

// The Sorted key-value List, not support duplicate keys
type SkipList struct {
	height  int             // current height
	header  *SkipListItem   // list header node
	compare CompareFunc     // key compare function
	path    []*SkipListItem // pre-nodes search path
	length  int64           // current len
}

func NewSkipList(compare CompareFunc) *SkipList {
	// the path and the header next both fill all nil with length
	return &SkipList{height: 1, header: NewSkipListItem(MAX_HEIGHT), compare: compare, path: make([]*SkipListItem, MAX_HEIGHT)}
}

func (this *SkipList) Head() *SkipListItem {
	return this.header
}

func (this *SkipList) Len() int64 {
	return this.length
}

func (this *SkipList) Clear() {
	if this.Len() > 0 {
		this.height = 1
		this.header = NewSkipListItem(MAX_HEIGHT)
		this.path = make([]*SkipListItem, MAX_HEIGHT)
		this.length = 0
	}
}

// if not exist, insert the key-value, else update the value
func (this *SkipList) Set(key, value interface{}) {
	next := this.findGEItem(key, this.path)
	// if exist, only update the value
	if next != nil && this.compare(key, next.key) == 0 {
		next.value = value
		return
	}
	height := this.randomHeight()
	// init search path higher level to header
	if height > this.height {
		for i := this.height; i < height; i++ {
			this.path[i] = this.header
		}
		this.height = height
	}
	node := NewSkipListItem(height)
	node.key = key
	node.value = value
	// insert the node to the skiplist, pre-nodes next to the new node
	for i := 0; i < height; i++ {
		node.SetNext(i, this.path[i].next[i])
		this.path[i].SetNext(i, node)
	}
	this.length++
}

// if not exist return false, else return true
func (this *SkipList) Contain(key interface{}) bool {
	node := this.findGEItem(key, nil)
	if node != nil && this.compare(key, node.key) == 0 {
		return true
	}
	return false
}

// if not exist return nil + false, else return value + true
func (this *SkipList) Get(key interface{}) (interface{}, bool) {
	node := this.findGEItem(key, nil)
	if node != nil && this.compare(key, node.key) == 0 {
		return node.value, true
	}
	return nil, false
}

// if not exist return nil + false, else return value + true
func (this *SkipList) Delete(key interface{}) (interface{}, bool) {
	node := this.findGEItem(key, this.path)
	if node != nil && this.compare(key, node.key) == 0 {
		this.length--
		height := len(node.next)
		// decrease the current height
		if height == this.height {
			for i := height; i > 1; i-- {
				if node.next[i-1] == nil && this.header.next[i-1] == node {
					this.height = i - 1
				}
			}
		}
		// remove the item from the list
		for i := 0; i < height; i++ {
			this.path[i].next[i] = node.next[i]
		}
		return node.value, true
	}
	return nil, false
}

//////////////////////////////////////////////////////////////////
/// private interface
//////////////////////////////////////////////////////////////////
// print all level data for debug
func (this *SkipList) printData(level int) {
	if level < 0 {
		for i := MAX_HEIGHT - 1; i >= 0; i-- {
			fmt.Print("level:", i, " members:")
			node := this.header.next[i]
			for node != nil {
				fmt.Print("<", node.key, node.value, ">")
				node = node.next[i]
			}
			fmt.Println()
		}
	} else {
		fmt.Print("level:", level, " members:")
		node := this.header.next[level]
		for node != nil {
			fmt.Print("<", node.key, node.value, ">")
			node = node.next[level]
		}
		fmt.Println()
	}
}

// print find path for debug
func (this *SkipList) printPath(heigth int) {
	fmt.Printf("print %d path:", heigth)
	for i := 0; i <= heigth; i++ {
		fmt.Print("<", i, this.path[i].value, ">")
	}
	fmt.Println()
}

// the key is GE than the node key
func (this *SkipList) isGE(key interface{}, node *SkipListItem) bool {
	return (node != nil) && (this.compare(node.key, key) < 0)
}

// get random height for a new node
func (this *SkipList) randomHeight() int {
	const kBranching = 4
	var height int = 1
	for height < MAX_HEIGHT && rand.Int()%kBranching == 0 {
		height++
	}
	return height
}

// find the first GE element
func (this *SkipList) findGEItem(key interface{}, path []*SkipListItem) *SkipListItem {
	level := this.height - 1
	cur := this.header
	var next *SkipListItem
	for {
		next = cur.next[level]
		if this.isGE(key, next) {
			cur = next
		} else {
			if path != nil {
				path[level] = cur
			}
			if level == 0 {
				return next
			} else {
				level--
			}
		}
	}
	return nil
}
