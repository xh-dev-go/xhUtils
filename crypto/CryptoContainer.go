package cust_crypto

import (
	"github.com/xh-dev-go/xhUtils/b64"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type Ava struct {
	Pri *rsa.PrivateKey
	Pub *rsa.PublicKey
}
type Bernice struct {
	Pub *rsa.PublicKey
}
type Envelope struct {
	Data string	`json:"data"`
}
type DataContainer struct {
	Sign string	`json:"sign"`
	Data string	`json:"data"`
	Key  string	`json:"key"`
	Iv   string	`json:"iv"`
}

func RandomBytes(size int) []byte {
	key := make([]byte, size)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}
	return key
}

func AesKey() string {
	return b64.EncodeString(RandomBytes(32))
}
func AesIv() string {
	return b64.EncodeString(RandomBytes(16))
}

func EncryptToEnvelope(data string, ava Ava, bernice Bernice) Envelope {
	key := AesKey()
	iv := AesIv()
	container := DataContainer{
		Sign: "",
		Data: "",
		Key:  b64.EncodeString(EncryptRsa([]byte(key), bernice.Pub)),
		Iv:   b64.EncodeString(EncryptRsa([]byte(iv), bernice.Pub)),
	}

	aesCipher, _ := aes.NewCipher(b64.DecodeString(key))

	encrypter := cipher.NewCBCEncrypter(aesCipher, b64.DecodeString(iv))

	plainText := PKCS5Padding([]byte(data), aes.BlockSize)
	cipherText := make([]byte, len(plainText))
	encrypter.CryptBlocks(cipherText, plainText)
	cipherData := b64.EncodeString(cipherText)
	container.Data = cipherData

	rng := rand.Reader
	hashed := sha256.Sum256([]byte(data))
	sign, err := rsa.SignPKCS1v15(rng, ava.Pri, crypto.SHA256, hashed[:])
	if err != nil {
		panic(err)
	}
	container.Sign = b64.EncodeString(sign)
	x, err := json.Marshal(container)
	if err != nil {
		panic(err)
	}

	envelope := Envelope{Data: b64.EncodeString(x)}

	return envelope
}

func EncryptToEnvelopeString(data string, ava Ava, bernice Bernice) string {
	envelopeStr, err := json.Marshal(EncryptToEnvelope(data, ava, bernice))
	if err != nil {
		panic(err)
	}

	return string(envelopeStr)
}

func DecryptEnvelopeString(envelopeStr string, ava Ava, bernice Bernice) string {
	var err error
	var envelope Envelope
	err = json.Unmarshal([]byte(envelopeStr), &envelope)
	if err != nil {
		panic(err)
	}
	return DecryptEnvelope(envelope, ava, bernice).Data
}

func DecryptEnvelope(envelope Envelope, ava Ava, bernice Bernice) DataContainer {
	var err error

	var dataContainer DataContainer
	err = json.Unmarshal(b64.DecodeString(envelope.Data), &dataContainer)
	if err != nil {
		panic(err)
	}

	dataContainer.Key = string(DecryptRsa(b64.DecodeString(dataContainer.Key), ava.Pri))
	dataContainer.Iv = string(DecryptRsa(b64.DecodeString(dataContainer.Iv), ava.Pri))

	aesCipher, _ := aes.NewCipher(b64.DecodeString(dataContainer.Key))

	decrypter := cipher.NewCBCDecrypter(aesCipher, b64.DecodeString(dataContainer.Iv))

	envDataByte, _ := base64.StdEncoding.DecodeString(dataContainer.Data)
	decrypter.CryptBlocks(envDataByte, envDataByte)

	str := string(PKCS5Trimming(envDataByte))
	hs := sha256.New()
	hs.Write([]byte(str))
	hss := hs.Sum(nil)
	err = rsa.VerifyPKCS1v15(bernice.Pub, crypto.SHA256, hss, b64.DecodeString(dataContainer.Sign))

	dataContainer.Data = str

	return dataContainer
}
