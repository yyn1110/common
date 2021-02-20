package common

import (
	"hash/fnv"
	"crypto/md5"
	"encoding/hex"
	"crypto/sha256"
	"io"
	"crypto/sha1"
	"fmt"
)

func GetMd5Hash(key []byte) uint64 {
	if key == nil {
		return 0
	}
	sum := md5.New().Sum(key)
	if len(sum) < 16 {
		return 0
	}
	b := sum[:8]
	p1 := (uint64(b[0]) << 56) |
			(uint64(b[1]) << 48) |
			(uint64(b[2]) << 40) |
			(uint64(b[3]) << 32) |
			(uint64(b[4]) << 24) |
			(uint64(b[5]) << 16) |
			(uint64(b[6]) << 8) |
			(uint64(b[7]) << 0)
	b = sum[8:]
	p2 := (uint64(b[0]) << 56) |
			(uint64(b[1]) << 48) |
			(uint64(b[2]) << 40) |
			(uint64(b[3]) << 32) |
			(uint64(b[4]) << 24) |
			(uint64(b[5]) << 16) |
			(uint64(b[6]) << 8) |
			(uint64(b[7]) << 0)
	p := p1 + p2
	return p
}

func GetFnvHash(key []byte) uint64 {
	h := fnv.New64a()
	h.Write(key)
	return h.Sum64()
}

func MD5(key []byte) string {
	h := md5.New()
	h.Write(key)
	return hex.EncodeToString(h.Sum(nil))
}

func SHA256(key []byte) string {
	hash := sha256.New()
	hash.Write(key)
	md := hash.Sum(nil)
	return hex.EncodeToString(md)
}

func SHA1(data string) string {
	t := sha1.New()
	io.WriteString(t,data)
	return fmt.Sprintf("%x",t.Sum(nil))
}
