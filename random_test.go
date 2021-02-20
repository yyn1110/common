package common

import (
	"strings"
	"testing"
)

func TestGenerateRandomString(t *testing.T) {
	for i := 0; i < 32; i++ {
		token := GenerateRandomString(i)
		if len(token) != i {
			t.Error(token)
		}
		for _, c := range token {
			if !strings.ContainsRune(RTABLE, c) {
				t.Error("not find the char", c)
			}
		}
	}
}
