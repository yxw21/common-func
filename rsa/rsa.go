package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

func Encrypt(data string, publicKey []byte) (string, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return "", errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}
	content, err := rsa.EncryptPKCS1v15(rand.Reader, pubInterface.(*rsa.PublicKey), []byte(data))
	if err != nil {
		return string(content), err
	}
	return base64.StdEncoding.EncodeToString(content), nil
}

func Decrypt(ciphertext string, privateKey []byte) (string, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return "", errors.New("private key error!")
	}
	priKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	content, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	content, err = rsa.DecryptPKCS1v15(rand.Reader, priKey.(*rsa.PrivateKey), []byte(content))
	return string(content), err
}

func Signature(data string, privateKey []byte) (string, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return "", errors.New("private key error!")
	}
	priKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	myHash := sha256.New()
	myHash.Write([]byte(data))
	dataHashText := myHash.Sum(nil)
	cipher, err := rsa.SignPKCS1v15(rand.Reader, priKey.(*rsa.PrivateKey), crypto.SHA256, dataHashText)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(cipher), nil
}

func Verify(data, cipher string, publicKey []byte) bool {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return false
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return false
	}
	content, err := base64.StdEncoding.DecodeString(cipher)
	if err != nil {
		return false
	}
	myHash := sha256.New()
	myHash.Write([]byte(data))
	dataHashText := myHash.Sum(nil)
	err = rsa.VerifyPKCS1v15(pubInterface.(*rsa.PublicKey), crypto.SHA256, dataHashText, []byte(content))
	if err != nil {
		return false
	} else {
		return true
	}
}
