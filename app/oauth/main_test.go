package main

import (
	"net/url"
	"testing"
)

func TestGetTokenUrl(t *testing.T) {
	tokenUrl, err := buildTokenUrl("clientId", "clientSecret", "code", "redirectUrl")
	if err != nil {
		t.Fatalf("buildTokenUrl error:%+v", err)
	}

	t.Logf("got tokenUrl:%s", tokenUrl)
}

func TestUrlQueryEscape(t *testing.T) {
	s := "Hello World!"
	t.Log(url.PathEscape(s))
}

func TestUrlValuesEncode(t *testing.T) {
	values := url.Values{}
	values.Add("param1", "value1")
	values.Add("param2", "value2")
	t.Log(values.Encode())
}

func TestUrlPathEscape(t *testing.T) {
	path := "Hello World!"
	t.Log(url.PathEscape(path))
}
