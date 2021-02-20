package common

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestCompareBasicType(t *testing.T) {
	var a int = 1
	var b int32 = 1
	// diff type
	if Compare(a, b, true) {
		t.Error("compare two type succ")
	}

	if Compare(1, 2, true) {
		t.Error("compare 1 2 failed")
	}
	if !Compare(1, 1, true) {
		t.Error("compare 1, 1 failed")
	}
	rand.Seed(time.Now().UnixNano())
	// array
	x := rand.Perm(30)
	y := rand.Perm(30)

	// should not equal with random algorithm, but not absolutely
	if Compare(x, y, true) {
		t.Error("not equal", x, y)
	} else if !Compare(x, y, false) {
		t.Error("should equal", x, y)
	}
}

func TestCompareNilType(t *testing.T) {
	if !Compare(nil, nil, true) {
		t.Error("compare nil should succ")
	}
	if Compare(nil, 1, false) {
		t.Error("compare nil should false")
	}
	if Compare(make([]int, 0), nil, false) {
		t.Error("compare nil should false")
	}
	a := make([]interface{}, 0)
	a = append(a, nil)
	b := make([]int, 0)
	b = append(b, 1)
	if Compare(a, b, false) {
		t.Error("not the same type")
	}
	/*
		c := make([]interface{}, 0)
		c = append(c, 1)
		if Compare(a, c, false) {
			t.Error("not the same type")
		}
	*/
}

func TestCompareComplexType(t *testing.T) {
	// map
	a := make(map[string]bool)
	b := make(map[string]bool)
	for i := 0; i < 10; i++ {
		a[fmt.Sprintf("%d", i)] = false
		b[fmt.Sprintf("%d", i)] = false
	}
	if !Compare(a, b, true) {
		t.Error("should equal", a, b)
	}
	b["x"] = false
	if Compare(a, b, true) {
		t.Error("not equal", a, b)
	}

	// array map
	c := make([]map[string]int, 0)
	d := make([]map[string]int, 0)
	for i := 0; i < 10; i++ {
		p := make(map[string]int)
		q := make(map[string]int)
		p[fmt.Sprintf("%d", i)] = i + 1
		q[fmt.Sprintf("%d", i)] = i + 1
		c = append(c, p)
		d = append(d, q)
	}
	if !Compare(c, d, true) {
		t.Error("compare failed", c, d)
	}
	if !Compare(d, c, true) {
		t.Error("compare failed", c, d)
	}
	d[0]["0"] = 100
	if Compare(c, d, true) {
		t.Error("compare should failed", c, d)
	}
	if Compare(d, c, true) {
		t.Error("compare should failed", c, d)
	}

	// reset and remove one item set another new item
	d[0]["0"] = 0 + 1
	if !Compare(c, d, true) {
		t.Error("compare failed", c, d)
	}
	if !Compare(d, c, true) {
		t.Error("compare failed", c, d)
	}
	delete(d[0], "0")
	d[0]["100"] = 100 + 1
	if Compare(c, d, true) {
		t.Error("compare should failed", c, d)
	}
	if Compare(d, c, true) {
		t.Error("compare should failed", c, d)
	}
	// reset to equal
	delete(d[0], "100")
	d[0]["0"] = 0 + 1
	if !Compare(c, d, true) {
		t.Error("compare failed", c, d)
	}
	if !Compare(d, c, true) {
		t.Error("compare failed", c, d)
	}

	// swap
	d[0], d[1] = d[1], d[0]
	if Compare(c, d, true) {
		t.Error("compare should failed", c, d)
	}
	if Compare(d, c, true) {
		t.Error("compare should failed", c, d)
	}
	// without strict check
	if !Compare(c, d, false) {
		t.Error("compare failed", c, d)
	}
	if !Compare(d, c, false) {
		t.Error("compare failed", c, d)
	}
}

const arrayCount int = 20

type A struct {
	name string
	id   uint64
	list [arrayCount]int
}

type B struct {
	A
	other int32
}

// struct array compare
func TestCompareStruct(t *testing.T) {
	var x B
	x.name = "test"
	x.id = 10
	for i := 0; i < len(x.list); i++ {
		x.list[i] = i
	}
	x.other = 11

	var y B
	y.name = "test"
	y.id = 10
	for i := 0; i < len(y.list); i++ {
		y.list[i] = arrayCount - i - 1
	}
	y.other = 11

	var z A
	z.id = 10
	z.name = "test"
	for i := 0; i < len(z.list); i++ {
		z.list[i] = arrayCount - i - 1
	}

	// fmt.Println(x, y, z)
	if Compare(x, z, true) {
		t.Error("should not equal")
	}

	if Compare(x, y, true) {
		t.Error("should not equal")
	}
	if Compare(y, x, true) {
		t.Error("should not equal")
	}

	if !Compare(x, y, false) {
		t.Error("should compare succ")
	}

	if !Compare(y, x, false) {
		t.Error("should compare succ")
	}

	y.other = 12
	if Compare(x, y, false) {
		t.Error("should not equal")
	}
}
