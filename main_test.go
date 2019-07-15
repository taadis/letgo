package main

import (
	"testing"

	"github.com/kataras/iris/httptest"
)

// 参考:
// 1. https://github.com/kataras/iris/tree/master/_examples/testing/httptest
// 2. https://studyiris.com/example/request/customViaUnmarshaler.html
func TestMain(t *testing.T) {
	app := newApp()
	e := httptest.New(t, app)
	e.GET("/").Expect().Status(httptest.StatusOK)
}
