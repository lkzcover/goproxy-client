package goproxy_client

import "testing"

func TestGetRequest(t *testing.T) {
	resp, err := GetRequest("http://localhost:9000", "https://www.google.com")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Response: %s", string(resp))
}

func TestGetAesRequest(t *testing.T) {

	key := "test123456789012"

	resp, err := GetAesRequest("http://localhost:9000", "https://www.google.com", key)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Response: %s", string(resp))
}
