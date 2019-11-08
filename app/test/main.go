// app/test/main.go
package main

import (
	"fmt"
	"net/http"

	"github.com/taadis/letgo/net/ghttp"
)

func main() {
	fmt.Println("ready to start test app")
	mux := ghttp.NewServeMux()
	mux.Get("/", func(c *ghttp.Context) {
		c.Saw(http.StatusOK, []byte("hello world"))
	})
	err := http.ListenAndServe(":5903", mux)
	if err != nil {
		fmt.Errorf("start golb error: ", err.Error())
	}
}
