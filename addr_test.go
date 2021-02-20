package common

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestServerAddr(t *testing.T) {
	var server ServerAddr
	rand.Seed(time.Now().UnixNano())
	for i := 1000; i < 2000; i++ {
		hostname := fmt.Sprintf("%d.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256))
		server.Set(hostname, int32(i))
		if server.IP() != hostname {
			t.Errorf("check server ip failed", server.IP(), hostname)
		} else if server.Port() != int32(i) {
			t.Error("check server port failed", server.Port(), i)
		}
	}
}
