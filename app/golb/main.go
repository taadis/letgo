// app/golb/main.go
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// User struct
type User struct {
	UserName string
}

func main() {
	fmt.Println("ready to start golb app")
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		user := User{
			UserName: "taadis",
		}
		tmp, err := template.ParseFiles("views/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		err = tmp.Execute(w, &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	//mux.HandleFunc("/system/user/", user.HandleFunc)
	//mux.Handle("/system/user/", midleware.PanicAndRecover(http.HandlerFunc(user.HandleFunc)))
	err := http.ListenAndServe(":5903", mux)
	if err != nil {
		fmt.Errorf("start golb error: ", err.Error())
	}
}
