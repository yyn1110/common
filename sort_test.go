package common

import (
	"testing"
	"math/rand"
)

func TestSort(t *testing.T) {
	perm := rand.Perm(1001)
	var values = make([]int64, 0, 1000)
	for _, value := range perm {
		values = append(values, int64(value))
	}

	SortInt64Slice(values)
	for i := 0; i < 1001; i++ {
		if values[i] != int64(i) {
			t.Error(perm, values)
		}
	}
}
