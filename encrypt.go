package goproxy_client

import (
	"bytes"
	"encoding/base64"
	"math/rand"

	"github.com/lkzcover/easyaes"
)

const letterBytes = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomIV(n int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return b
}

func encryptTargetURLReq(queryReq, key string) (string, []byte, error) {

	iv := randomIV(len(key))

	target, err := easyaes.EncryptAesCBCStaticIV([]byte(key), iv, []byte(queryReq))
	if err != nil {
		return "", nil, err
	}

	var splitByte bytes.Buffer

	splitByte.Write(iv)
	splitByte.Write(target)

	return base64.URLEncoding.EncodeToString(splitByte.Bytes()), iv, nil

}
