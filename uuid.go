package common

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

// A UUID is a 128 bit (16 byte) Universal Unique IDentifier as defined in RFC
// 4122.
type UUID []byte

// xvalues returns the value of a byte as a hexadecimal digit or 255.
var xvalues = []byte{
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 255, 255, 255, 255, 255, 255,
	255, 10, 11, 12, 13, 14, 15, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 10, 11, 12, 13, 14, 15, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255,
}

// xtob converts the the first two hex bytes of x into a byte.
func xtob(x string) (byte, bool) {
	b1 := xvalues[x[0]]
	b2 := xvalues[x[1]]
	return (b1 << 4) | b2, b1 != 255 && b2 != 255
}

func GenUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := rand.Read(uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}

	// RFC 4122
	uuid[8] = 0x80 // variant bits see page 5
	uuid[4] = 0x40 // version 4 pseudo random, see page 7

	return hex.EncodeToString(uuid), nil
}

// Parse decodes s into a UUID or returns nil.  Both the UUID form of
// xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx and
// urn:uuid:xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx are decoded.
func Parse(s string) (UUID, error) {
	var tmp string
	if len(s) == 8 {
		tmp = "000000000000000000000000" + s
	} else if len(s) == 16 {
		tmp = "0000000000000000" + s
	} else if len(s) == 32 {
		tmp = s
	} else {
		return nil, ErrInvalidParam
	}

	uuid := make([]byte, 16)
	for i, x := range []int{
		0, 2, 4, 6,
		8, 10, 12, 14,
		16, 18, 20, 22,
		24, 26, 28, 30} {
		if v, ok := xtob(tmp[x:]); !ok {
			return nil, ErrInvalidParam
		} else {
			uuid[i] = v
		}
	}
	return uuid, nil
}

// Equal returns true if uuid1 and uuid2 are equal.
func Equal(uuid1, uuid2 UUID) bool {
	return bytes.Equal(uuid1, uuid2)
}

// String returns the string form of uuid, xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
// , or "" if uuid is invalid.
func (uuid UUID) String() string {
	if uuid == nil || len(uuid) != 16 {
		return ""
	}
	b := []byte(uuid)
	return fmt.Sprintf("%08X%04X%04X%04X%012X",
		b[:4], b[4:6], b[6:8], b[8:10], b[10:])
}
