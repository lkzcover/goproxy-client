package goproxy_client

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/lkzcover/easyaes"
)

const (
	ContentTypeJson = "application/json"
)

func PostRequest(urlReq, queryReq, contentType string, bodyReq interface{}) ([]byte, error) {

	body, err := json.Marshal(bodyReq)
	if err != nil {
		return nil, fmt.Errorf("marshal bodyReq error: %s", err)
	}

	resp, err := http.Post(urlReq+decryptedReq+queryReq, contentType, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("incorrect goproxy response status. Get: %d but Expecdet: 200", resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}

func PostAesRequest(urlReq, queryReq, key, contentType string, bodyReq interface{}) ([]byte, error) {

	target, iv, err := encryptTargetURLReq(queryReq, key)

	urlReq = urlReq + encryptedReq + target

	body, err := json.Marshal(bodyReq)
	if err != nil {
		return nil, fmt.Errorf("marshal bodyReq error: %s", err)
	}

	encryptedBody, err := easyaes.EncryptAesCBCStaticIV([]byte(key), iv, body)
	if err != nil {
		return nil, fmt.Errorf("encrypted body POST request error: %s", err)
	}

	resp, err := http.Post(urlReq, contentType, bytes.NewReader([]byte(base64.StdEncoding.EncodeToString(encryptedBody)))) // TODO оптимизировать
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

	decryptResp, err := easyaes.DecryptAesCBCStaticIV([]byte(key), iv, respBody)
	if err != nil {
		return nil, err
	}

	return decryptResp, nil
}
