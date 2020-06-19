package http

import (
	"net/http"
	"testing"
)

// 测试 SetBasicAuth 函数是否有用
// 可参考官方测试:
// https://github.com/golang/go/blob/1fd3f8bd67a36e330c8be07941d1ab09870ff932/src/net/http/request_test.go#L587
func TestSetBasicAuth(t *testing.T) {
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
		t.FailNow()
	}
	req.SetBasicAuth("name", "123456")
	username, password, ok := req.BasicAuth()
	if !ok || username != "name" || password != "123456" {
		t.Errorf("test fail")
	}
}

// TestGet
func TestGet(t *testing.T) {
	request, err := http.NewRequest()
}
