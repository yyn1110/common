package common

import (

	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"math/big"
)

///////////////////////////////////////////////////////////////////
// RSA EN/DECRYPTION using pem file content as key
///////////////////////////////////////////////////////////////////
func RsaEncrypt(publicKey []byte, orignal []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, ErrInvalidEncryptKey
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, orignal)
}

func RsaDecrypt(privateKey []byte, crypted []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, ErrInvalidEncryptKey
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, crypted)
}

///////////////////////////////////////////////////////////////////
// RSA EN/DECRYPTION using N,E as key
///////////////////////////////////////////////////////////////////
func RsaBlockEncrpt(blockSize int, publicKey []byte, orignal []byte) ([]byte, error) {
	CheckParam(blockSize > 0)
	dest := make([]byte, 0)
	for i := 0; i <= len(orignal)/blockSize; i++ {
		last := (i + 1) * blockSize
		if last > len(orignal) {
			last = len(orignal)
		}
		block, err := RsaEncrypt2(publicKey, orignal[i*blockSize:last])
		if err != nil {
			fmt.Printf("RSA encrypt failed:data[%d], err[%v]", i*blockSize, err)
			fmt.Printf("public[%v] data[%v]", hex.EncodeToString(publicKey), hex.EncodeToString(orignal))
			return nil, err
		}
		dest = append(dest, block...)
	}
	return dest, nil
}

func RsaBlockDecrpt(blockSize int, privateKey []byte, crypted []byte) ([]byte, error) {
	CheckParam(blockSize > 0)
	if len(crypted)%blockSize != 0 {
		fmt.Println("check message len for rsa version failed", len(crypted))
		return nil, ErrInvalidParam
	}
	dest := make([]byte, 0)
	for i := 0; i < len(crypted)/blockSize; i++ {
		block, err := RsaDecrypt(privateKey, crypted[i*blockSize:(i+1)*blockSize])
		if err != nil {
			fmt.Printf("RSA decrypt failed:data[%d], err[%v]", i*blockSize, err)
			fmt.Printf("private[%v] data[%v]", hex.EncodeToString(privateKey), hex.EncodeToString(crypted))
			return nil, err
		}
		dest = append(dest, block...)
	}
	return dest, nil
}

// using N,E for encrypt
func RsaEncrypt2(publicKey []byte, orignal []byte) ([]byte, error) {
	var key big.Int
	key.SetString(string(publicKey), 10)
	public := rsa.PublicKey{N: &key, E: 65537}
	return rsa.EncryptPKCS1v15(rand.Reader, &public, orignal)
}

///////////////////////////////////////////////////////////////////
// DES EN/DECRYPTION
///////////////////////////////////////////////////////////////////
func DesEncrypt(publicKey []byte, orignal []byte) ([]byte, error) {
	block, err := des.NewCipher(publicKey)
	if err != nil {
		return nil, err
	}
	orignal = PKCS5Padding(orignal, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, publicKey)
	crypted := make([]byte, len(orignal))
	blockMode.CryptBlocks(crypted, orignal)
	return crypted, nil
}

func DesDecrypt(publicKey []byte, crypted []byte) ([]byte, error) {
	block, err := des.NewCipher(publicKey)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, publicKey)
	orignal := make([]byte, len(crypted))
	blockMode.CryptBlocks(orignal, crypted)
	orignal = PKCS5UnPadding(orignal)
	return orignal, nil
}

///////////////////////////////////////////////////////////////////
// AES EN/DECRYPTION
///////////////////////////////////////////////////////////////////
func AesEncrypt(publicKey []byte, orignal []byte) ([]byte, error) {
	block, err := aes.NewCipher(publicKey)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	orignal = PKCS5Padding(orignal, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, publicKey[:blockSize])
	crypted := make([]byte, len(orignal))
	blockMode.CryptBlocks(crypted, orignal)
	return crypted, nil
}

func AesDecrypt(publicKey []byte, crypted []byte) ([]byte, error) {
	block, err := aes.NewCipher(publicKey)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, publicKey[:blockSize])
	orignal := make([]byte, len(crypted))
	blockMode.CryptBlocks(orignal, crypted)
	orignal = PKCS5UnPadding(orignal)
	return orignal, nil
}

//////////////////////////////////////////////////////////////////////
/// padding for the last block
//////////////////////////////////////////////////////////////////////
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(orignal []byte) []byte {
	length := len(orignal)
	unpadding := int(orignal[length-1])
	return orignal[:(length - unpadding)]
}
