// app/golb/main.go
package main

import (
	"fmt"
	"net/http"

	"github.com/taadis/letgo/handler/company"
)

func main() {
	fmt.Println("ready to start golb app")
	mux := http.NewServeMux()
	mux.HandleFunc("/", company.Index)
	err := http.ListenAndServe(":5903", mux)
	if err != nil {
		fmt.Errorf("start golb error: ", err.Error())
	}
	//fmt.Println("golb running in: ", "http://localhost:5903")
}
