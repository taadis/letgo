// app/test/main.go
package main

import (
	"fmt"
	"net/http"

	app "github.com/taadis/letgo/net/http"
)

func main() {
	fmt.Println("ready to start test app")
	mux := app.NewServeMux()
	mux.Get("/", func(c *app.Context) {
		c.Saw(http.StatusOK, []byte("hello world"))
	})
	mux.Get("/plain", func(c *app.Context) {
		c.Plain(http.StatusOK, []byte("plain"))
	})
	mux.Get("/json/", func(c *app.Context) {
		c.Json(http.StatusOK, `{"key":"value"}`)
	})
	mux.Get("/xml/", func(c *app.Context) {
		c.Xml(http.StatusOK, "<xml></xml>")
	})
	err := http.ListenAndServe(":5903", mux)
	if err != nil {
		fmt.Errorf("start golb error: ", err.Error())
	}
}
