package goproxy_client

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/lkzcover/easyaes"
)

func GetRequest(urlReq, queryReq string) ([]byte, error) {

	resp, err := http.Get(urlReq + decryptedReq + queryReq)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("incorrect goproxy response status. Get: %d but Expecdet: 200", resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}

func GetAesRequest(urlReq, queryReq, key string) ([]byte, error) {

	target, iv, err := encryptTargetURLReq(queryReq, key)

	urlReq = urlReq + encryptedReq + target

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

	decryptResp, err := easyaes.DecryptAesCBCStaticIV([]byte(key), iv, respBody)
	if err != nil {
		return nil, err
	}

	return decryptResp, nil
}
