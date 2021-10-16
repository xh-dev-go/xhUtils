package sha256

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

func HashFile0x(file string) (string, error) {
	var str string
	hashed, err := HashFile(file)
	if err != nil {
		return str, err
	}
	return fmt.Sprintf("%x", hashed), nil
}
func HashFile(file string) ([]byte, error) {
	var ba []byte
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return ba, err
	}
	hasher  := sha256.New()
	hasher.Write(bytes)
	return hasher.Sum(nil), nil
}
func HashFileBase64(file string) (string, error) {
	var str string
	hashed, err := HashFile(file)
	if err != nil {
		return str, err
	}
	str = base64.StdEncoding.EncodeToString(hashed)
	return str, nil
}
