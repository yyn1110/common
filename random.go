package common

import (
	crand "crypto/rand"
	"math/rand"
	"time"
)

// generate random key not readable string
func GenerateRandomKey(length int) []byte {
	rand.Seed(time.Now().UTC().UnixNano())
	key := make([]byte, length)
	for i := 0; i < length; i++ {
		key[i] = byte(rand.Intn(256))
	}
	return key
}

// readable string char table
const RTABLE = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// generate random readable string consisting of letters and numbers
func GenerateRandomString(length int) string {
	var bytes = make([]byte, length)
	tableLen := len(RTABLE)
	crand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = RTABLE[b%byte(tableLen)]
	}
	return string(bytes)
}

// token key for direct communication
func GenerateDeviceToken(length int) string {
	return GenerateRandomString(length)
}
