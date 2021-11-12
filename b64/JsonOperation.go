package b64

import "encoding/base64"

func DecodeString(str string) []byte {
	msg, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		panic(err)
	}
	return msg
}

func EncodeString(b64Data []byte) string{
	return base64.StdEncoding.EncodeToString(b64Data)
}