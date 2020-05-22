package goproxy_client

import "testing"

func TestPostRequest(t *testing.T) {

	tesReq := struct {
		Test string
	}{Test: "hello world"}

	resp, err := PostRequest("http://localhost:9000", "https://webhook.site/1e4c45eb-b130-48b6-9b2d-1054e9885eee", ContentTypeJson, tesReq)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Response: %s", string(resp))
}

func TestPostAesRequest(t *testing.T) {

	key := "test123456789012"

	tesReq := struct {
		Test string
	}{Test: "hello encrypt world"}

	resp, err := PostAesRequest("http://localhost:9000", "https://webhook.site/1e4c45eb-b130-48b6-9b2d-1054e9885eee", key, ContentTypeJson, tesReq)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Response: %s", string(resp))
}
