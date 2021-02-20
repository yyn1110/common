package common

import (
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := NewPool(100)
	// push all elems
	for i := 0; i < 100; i++ {
		err := pool.Put(i, 1000)
		if err != nil {
			t.Error("put obj failed")
		}
	}

	// all timeout
	for i := 0; i < 100; i++ {
		err := pool.Put(i, 1000)
		if err != ErrTimeout {
			t.Error("put obj not timeout")
		}
	}

	// pop all the elems
	for i := 0; i < 100; i++ {
		temp := pool.Get(1000 * 1000)
		if temp == nil {
			t.Error("check get obj failed")
		}
	}

	for i := 0; i < 100; i++ {
		temp := pool.Get(1000 * 1000)
		if temp != nil {
			t.Error("check get obj failed")
		}
	}

	// push all elems again
	for i := 0; i < 100; i++ {
		err := pool.Put(i, 1000)
		if err != nil {
			t.Error("put obj failed")
		}
	}

	go func() {
		for i := 0; i < 100; i++ {
			temp := pool.Get(1000 * 1000)
			if temp == nil {
				t.Error("check get obj failed")
			} else {
				// fmt.Println("pop succ", i)
			}
			time.Sleep(time.Millisecond * 2)
		}
	}()

	// put the last 2 timeout
	for i := 0; i < 102; i++ {
		err := pool.Put(i, time.Second.Nanoseconds())
		if i < 100 {
			if err != nil {
				t.Error("check put failed", err)
			}
		}
	}
}
