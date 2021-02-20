package common

import (
	"testing"
)

func Test(t *testing.T) {
	err := ErrNullValue
	temp := Error(err.Code())
	if err != temp {
		t.Error("not the same error")
	}
}
