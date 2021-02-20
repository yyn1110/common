package common

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"math/big"
)

type Type int64

const (
	PKCS1 Type = iota
	PKCS8
)

const (
	SIGN int64 = iota
	VERIFY
)

type Cipher struct {
	mode       int64
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

//默认客户端，pkcs1私钥格式，pem编码
func NewCipherByPem(mode int64, key []byte) (*Cipher, error) {
	var blockPri, blockPub *pem.Block
	switch mode {
	case SIGN:
		blockPri, _ = pem.Decode([]byte(key))
		if blockPri == nil {
			return nil, errors.New("private key error")
		}
		return newCipher(blockPri.Bytes, PKCS1, mode)
	case VERIFY:
		blockPub, _ = pem.Decode([]byte(key))
		if blockPub == nil {
			return nil, errors.New("public key error")
		}
		return newCipher(blockPub.Bytes, PKCS1, mode)
	default:
		return nil, errors.New("invalid mode")
	}
}

func NewCipher(mode int64, privateKey *rsa.PrivateKey, publicKey  *rsa.PublicKey)  *Cipher {
	if privateKey == nil && publicKey == nil {
		return nil
	}
	return &Cipher{
		mode:       mode,
		privateKey: privateKey,
		publicKey:  publicKey,
	}
}

func newCipher(key []byte, privateKeyType Type, mode int64) (*Cipher, error) {
	var priKey *rsa.PrivateKey
	var pubKey *rsa.PublicKey
	var err error
	switch mode {
	case SIGN:
		priKey, err = genPriKey(key, privateKeyType)
		if err != nil {
			return nil, err
		}
	case VERIFY:
		pubKey, err = genPubKey(key)
		if err != nil {
			return nil, err
		}
	}

	return &Cipher{privateKey: priKey, publicKey: pubKey, mode: mode}, nil
}

func genPubKey(publicKey []byte) (*rsa.PublicKey, error) {
	pub, err := x509.ParsePKIXPublicKey(publicKey)
	if err != nil {
		return nil, err
	}
	return pub.(*rsa.PublicKey), nil
}

func genPriKey(privateKey []byte, privateKeyType Type) (*rsa.PrivateKey, error) {
	var priKey *rsa.PrivateKey
	var err error
	switch privateKeyType {
	case PKCS1:
		{
			priKey, err = x509.ParsePKCS1PrivateKey([]byte(privateKey))
			if err != nil {
				return nil, err
			}
		}
	case PKCS8:
		{
			prkI, err := x509.ParsePKCS8PrivateKey([]byte(privateKey))
			if err != nil {
				return nil, err
			}
			priKey = prkI.(*rsa.PrivateKey)
		}
	default:
		{
			return nil, errors.New("unsupport private key type")
		}
	}
	return priKey, nil
}


func (this *Cipher) Sign(src []byte, hash crypto.Hash) ([]byte, error) {
	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto // for simple example
	h := hash.New()
	h.Write(src)
	hashed := h.Sum(nil)
	return rsa.SignPSS(rand.Reader, this.privateKey, hash, hashed, &opts)
	//return rsa.SignPKCS1v15(rand.Reader, this.privateKey, hash, hashed)
}

func (this *Cipher) Verify(src []byte, sign []byte, hash crypto.Hash) error {
	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto // for simple example
	h := hash.New()
	h.Write(src)
	hashed := h.Sum(nil)
	return rsa.VerifyPSS(this.publicKey, hash, hashed, sign, &opts)
}

func EncodePublicKey(publicKey []byte) *rsa.PublicKey {
	var key big.Int
	key.SetString(string(publicKey), 10)
	return &rsa.PublicKey{N: &key, E: 65537}
}

func DecodePublicKey(publicKey []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	public := pubInterface.(*rsa.PublicKey)
	return public, nil
}

