package goproxy_client

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/lkzcover/easyaes"
)

func GetRequest(urlReq, queryReq string) ([]byte, error) {

	resp, err := http.Get(urlReq + "/?target=" + queryReq)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("incorrect goproxy response status. Get: %d but Expecdet: 200", resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}

func GetAesRequest(urlReq, queryReq, key string) ([]byte, error) {

	urlReq = urlReq + "/?type=e&target="

	iv := randomIV(len(key))

	target, err := easyaes.EncryptAesCBCStaticIV([]byte(key), iv, []byte(queryReq))
	if err != nil {
		return nil, err
	}

	var splitByte bytes.Buffer

	splitByte.Write(iv)
	splitByte.Write(target)

	urlReq = urlReq + url.QueryEscape(base64.StdEncoding.EncodeToString(splitByte.Bytes()))

	resp, err := http.Get(urlReq)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("incorrect goproxy response status. Get: %d but Expecdet: 200", resp.StatusCode)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	decryptResp, err := easyaes.DecryptAesCBCStaticIV([]byte(key), []byte(iv), respBody)
	if err != nil {
		return nil, err
	}

	return decryptResp, nil
}
