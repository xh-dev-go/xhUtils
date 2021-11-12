package cust_crypto

import (
	"bytes"
	"github.com/xh-dev-go/xhUtils/b64"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

func GetPrivKey(key string) *rsa.PrivateKey {
	block,_ := pem.Decode([]byte(key))
	if block==nil{
		panic(errors.New("private key not valid"))
	}
	priv,_ := x509.ParsePKCS1PrivateKey(block.Bytes)
	return priv
}
func GetPublicKey(key string) *rsa.PublicKey{
	block,_ :=pem.Decode([]byte(key))
	if block==nil {
		panic(errors.New("public key not valid"))
	}
	pub, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err!=nil {
		panic(err)
	}
	return pub
}


func DecryptRsa(cipherText []byte , priv *rsa.PrivateKey) []byte{
	decryptedMessage, _ := rsa.DecryptPKCS1v15(rand.Reader, priv, cipherText)
	return decryptedMessage
}

func EncryptRsa(plainText []byte , pub *rsa.PublicKey) []byte{
	bytes,err :=rsa.EncryptPKCS1v15(rand.Reader, pub,  plainText)
	if err != nil{
		panic(err)
	}
	return bytes
}
func EncryptAsBase64(msg string, key string) string{
	return b64.EncodeString(EncryptRsaFromStr(msg,key))
}
func EncryptRsaFromStr(msg string, key string) []byte{
	pk := GetPublicKey(key)
	return EncryptRsa([]byte(msg), pk)
}

// Copied from https://gist.github.com/ltyyz/b306746041b48a8366d0f63507a4e7f3
func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
// Copied from https://gist.github.com/ltyyz/b306746041b48a8366d0f63507a4e7f3
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}