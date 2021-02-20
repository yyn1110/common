package common

import (
	"testing"
)

// Basic ADD\REMOVE\CONTAIN\LEN
func TestBasicSet(t *testing.T) {
	s := NewSet()
	// not exist
	for i := 0; i < 100; i++ {
		succ := s.Add(i)
		if !succ {
			t.Error("find the elem:", i)
		}
	}
	if s.Len() != 100 {
		t.Error("check set length failed", s.Len())
	}

	// contain
	for i := 0; i < 100; i++ {
		succ := s.Contain(i)
		if !succ {
			t.Error("not find the elem:", i)
		}
	}

	// already exist
	for i := 0; i < 100; i++ {
		succ := s.Add(i)
		if succ {
			t.Error("not find the elem:", i)
		}
	}
	if s.Len() != 100 {
		t.Error("check set length failed", s.Len())
	}

	// remove exist
	for i := 0; i < 100; i++ {
		succ := s.Remove(i)
		if !succ {
			t.Error("not find the elem:", i)
		}
	}

	// not contain
	for i := 0; i < 100; i++ {
		succ := s.Contain(i)
		if succ {
			t.Error("not find the elem:", i)
		}
	}

	if s.Len() != 0 {
		t.Error("check set len failed", s.Len())
	}
	// remove not exist
	for i := 0; i < 100; i++ {
		succ := s.Remove(i)
		if succ {
			t.Error("not find the elem:", i)
		}
	}

	if s.Len() != 0 {
		t.Error("check set len failed", s.Len())
	}
}

func TestSetUnion(t *testing.T) {
	// Union
	s1 := NewSet()
	for i := 0; i < 100; i++ {
		succ := s1.Add(i)
		if !succ {
			t.Error("find the elem:", i)
		}
	}
	s2 := NewSet()
	for i := 100; i < 200; i++ {
		succ := s2.Add(i)
		if !succ {
			t.Error("find the elem:", i)
		}
	}

	if s1.IsSubset(s2) || s2.IsSubset(s1) {
		t.Error("s2 is not s1 subset")
	}

	// 0 - 200
	s3 := s1.Union(s2)
	for i := 0; i < 200; i++ {
		succ := s3.Contain(i)
		if !succ {
			t.Error("not find the elem:", i)
		}
	}

	if !s1.IsSubset(s3) || !s2.IsSubset(s3) {
		t.Error("s1 or s2 is not s3 subset")
	}

	s4 := s3.Union(s1)
	for i := 0; i < 200; i++ {
		succ := s3.Contain(i)
		if !succ {
			t.Error("not find the elem:", i)
		}
	}

	s5 := NewSet()
	for i := 150; i < 300; i++ {
		succ := s5.Add(i)
		if !succ {
			t.Error("find the elem:", i)
		}
	}

	if s1.IsSubset(s5) || s2.IsSubset(s5) {
		t.Error("s1 or s2 is s5 subset")
	}

	s6 := s5.Union(s4)
	for i := 0; i < 300; i++ {
		succ := s6.Contain(i)
		if !succ {
			t.Error("not find the elem:", i)
		}
	}
	if !(s1.IsSubset(s6) && s2.IsSubset(s6) && s3.IsSubset(s6) && s4.IsSubset(s6) && s6.IsSubset(s6)) {
		t.Error("sx is not s6 subset")
	}
}

func TestSetIntersect(t *testing.T) {
	// Intersect
	s1 := NewSet()
	for i := 0; i < 100; i++ {
		succ := s1.Add(i)
		if !succ {
			t.Error("find the elem:", i)
		}
	}
	s2 := NewSet()
	for i := 100; i < 200; i++ {
		succ := s2.Add(i)
		if !succ {
			t.Error("find the elem:", i)
		}
	}
	s3 := s1.Intersect(s2)
	if s3.Len() != 0 {
		t.Error("check interset set failed", s3.Len())
	}

	s4 := s1.Intersect(s1)
	if s4.Len() != 100 {
		t.Error("check interset failed", s4.Len())
	}

	// 50-100
	s5 := NewSet()
	for i := 50; i < 150; i++ {
		succ := s5.Add(i)
		if !succ {
			t.Error("find the elem:", i)
		}
	}

	s6 := s1.Intersect(s5)
	if s6.Len() != 50 {
		t.Error("check interset failed", s6.Len())
	}
	for i := 50; i < 100; i++ {
		succ := s6.Contain(i)
		if !succ {
			t.Error("not find the elem:", i)
		}
	}

	s7 := s5.Intersect(s1)
	if s7.Len() != s6.Len() {
		t.Error("check result failed", s7.Len(), s6.Len())
	}

	if !s7.IsSubset(s6) || !s6.IsSubset(s7) {
		t.Error("check set not equal")
	}
}

func TestSetDifference(t *testing.T) {
	// difference
	s1 := NewSet()
	for i := 0; i < 100; i++ {
		succ := s1.Add(i)
		if !succ {
			t.Error("find the elem:", i)
		}
	}
	s2 := NewSet()
	for i := 100; i < 200; i++ {
		succ := s2.Add(i)
		if !succ {
			t.Error("find the elem:", i)
		}
	}

	// (0-100).Diff(100-200)=(0-200)
	s3 := s1.Difference(s2)
	for i := 0; i < 200; i++ {
		succ := s3.Contain(i)
		if !succ {
			t.Error("find the elem:", i)
		}
	}
	if s3.Len() != 200 {
		t.Error("check diff len failed", s3.Len())
	}

	// (1-100).Diff(1-100) = empty
	s4 := s1.Difference(s1)
	if s4.Len() != 0 {
		t.Error("check no diff failed")
	}

	// 50-150
	s6 := NewSet()
	for i := 50; i < 150; i++ {
		succ := s6.Add(i)
		if !succ {
			t.Error("find the elem", i)
		}
	}
	// (0-100).Diff(50-150) = (0-50).Union(100-150)
	s7 := s6.Difference(s1)
	for i := 0; i < 150; i++ {
		if i < 50 || i >= 100 {
			succ := s7.Contain(i)
			if !succ {
				t.Error("not find exist elem", i)
			}
		} else {
			succ := s7.Contain(i)
			if succ {
				t.Error("find not exist elem", i)
			}
		}
	}
	if s7.Len() != 100 {
		t.Error("check diff len failed", s7.Len())
	}

	// (0-200).Diff((0-50).Union(100-150)) = (50-100).Unoin(150-200).Interset(1-100)
	s8 := s3.Difference(s7).Intersect(s1)
	if s8.Len() != 50 {
		t.Error("check len s8 failed")
	}
	for i := 50; i < 100; i++ {
		succ := s8.Contain(i)
		if !succ {
			t.Error("not find the elem:", i)
		}
	}
}

func TestClear(t *testing.T) {
	s := NewSet()
	s.Clear()
	if s.Len() != 0 {
		t.Error("check len failed")
	}
	for i := 0; i < 100; i++ {
		succ := s.Add(i)
		if !succ {
			t.Error("find the elem:", i)
		} else if s.Len() != int64(i+1) {
			t.Error("set len failed", s.Len())
		}
	}
	s.Clear()
	if s.Len() != 0 {
		t.Error("check len failed")
	}
	// add again
	for i := 0; i < 100; i++ {
		succ := s.Add(i)
		if !succ {
			t.Error("find the elem:", i)
		} else if s.Len() != int64(i+1) {
			t.Error("set len failed", s.Len())
		}
	}
}

func TestSetWalk(t *testing.T) {
	s := NewSet()
	// first add
	for i := 0; i < 10; i++ {
		succ := s.Add(i)
		if !succ {
			t.Error("find the elem:", i)
		}
	}
	if s.Len() != 10 {
		t.Error("check set length failed", s.Len())
	}
	count := 0
	sum := 0
	walker := func(elem interface{}) {
		count ++
		sum += elem.(int)
	}
	s.Walk(walker)
	if count != 10 && sum != 45 {
		t.Error("walk failed:", count)
	}
}