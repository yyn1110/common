package common

import (
	"bytes"
	"testing"
)

var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MDwwDQYJKoZIhvcNAQEBBQADKwAwKAIhALEsynDW+CmeFceZ8OHMK2wmtc8Cyvuv
cHgEjwCBjvd5AgMBAAE=
-----END PUBLIC KEY-----
`)

var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIGrAgEAAiEAsSzKcNb4KZ4Vx5nw4cwrbCa1zwLK+69weASPAIGO93kCAwEAAQIh
AK4+k7XX5OXhICBWaE1Yo1YtxNNeTnwR410j+cLhX8eBAhEAzvsMrTpcNpvD56Nj
jAuVSQIRANsiqcJ4fgLnYhupYu7mwLECEQCuSJaUBuA+20pKcjoQYnCBAhAs75O+
JDU65TXSFE8MTFdRAhAeD3CcJhSVipGdyYH+zDRX
-----END RSA PRIVATE KEY-----
`)

func TestRSA(t *testing.T) {
	message := "leverly@126.com"
	data, err := RsaEncrypt(publicKey, []byte(message))
	if err != nil {
		t.Error("encrypt data failed", err)
	}

	old, err := RsaDecrypt(privateKey, data)
	if err != nil {
		t.Error("decrypt data failed", err)
	}

	if !bytes.Equal([]byte(message), old) {
		t.Error("check decrypt result failed")
	}
}

func TestAES(t *testing.T) {
	message := "leverly@126.com"
	sessionKey := []byte("abcdefghijklmnop")
	data, err := AesEncrypt(sessionKey, []byte(message))
	if err != nil {
		t.Error("encrypt data failed", err)
	}
	old, err := AesDecrypt(sessionKey, data)
	if err != nil {
		t.Error("decrypt data failed", err)
	}

	if !bytes.Equal([]byte(message), old) {
		t.Error("check decrypt result failed")
	}
}

func TestDES(t *testing.T) {
	message := "leverly@126.com"
	sessionKey := []byte("abcdefgh")
	data, err := DesEncrypt(sessionKey, []byte(message))
	if err != nil {
		t.Error("encrypt data failed", err)
	}
	old, err := DesDecrypt(sessionKey, data)
	if err != nil {
		t.Error("decrypt data failed", err)
	}

	if !bytes.Equal([]byte(message), old) {
		t.Error("check decrypt result failed")
	}
}

func TestDecryptUsingPem(t *testing.T) {
	old := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	data, err := RsaEncrypt2([]byte("80138512665003396643737838315916663972728479914654754587175091902061894104953"), old)
	if err != nil {
		t.Error("encrypt failed", err)
	}
	old2, err := RsaDecrypt(privateKey, data)
	if err != nil {
		t.Error("decrypted failed", err)
	}

	if !bytes.Equal(old, old2) {
		t.Error("check decrypted data failed")
	}
}
