package common

import (
	"net"
	"strconv"
	"strings"
)

type ServerAddr struct {
	ip   int64
	port int32
}

func (this *ServerAddr) Set(ip string, port int32) {
	this.ip = IPv4Atoi(ip)
	this.port = port
}

func (this *ServerAddr) IP() string {
	return IPv4Itoa(this.ip)
}

func (this *ServerAddr) Port() int32 {
	return this.port
}

func (this ServerAddr) String() string {
	return IPv4Itoa(this.ip) + ":" + strconv.Itoa(int(this.port))
}

// convert int64 to string
func IPv4Itoa(ip int64) string {
	var bytes [4]byte
	bytes[0] = byte(ip & 0xFF)
	bytes[1] = byte((ip >> 8) & 0xFF)
	bytes[2] = byte((ip >> 16) & 0xFF)
	bytes[3] = byte((ip >> 24) & 0xFF)
	return net.IPv4(bytes[3], bytes[2], bytes[1], bytes[0]).String()
}

// convert string to int64
func IPv4Atoi(ip string) int64 {
	bits := strings.Split(ip, ".")
	if len(bits) != 4 {
		return 0
	}
	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])
	var sum int64
	sum += int64(b0) << 24
	sum += int64(b1) << 16
	sum += int64(b2) << 8
	sum += int64(b3)
	return sum
}
