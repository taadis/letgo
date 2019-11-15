package controllers

import (
	"fmt"
	"net/http"

	"github.com/taadis/letgo/app/golb/vm"
)

// 定义home结构体，我理解的类似于php的类文件，之后定义了一个registerRoutes方法，接收路由，下面是 indexHandler 和loginHandler方法
type home struct{}

func (h home) registerRoutes() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", loginHandler)
}

//
func indexHandler(w http.ResponseWriter, r *http.Request) {
	vop := vm.IndexViewModelOp{}
	v := vop.GetVM()
	templates["index.html"].Execute(w, &v)
}

//
func loginHandler(w http.ResponseWriter, r *http.Request) {
	tpName := "login.html"
	vop := vm.LoginViewModelOp{}
	v := vop.GetVM()
	if r.Method == http.MethodGet {
		templates[tpName].Execute(w, &v)
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		fmt.Fprintf(w, "Username:%s,Password%s:", username, password)
	}
}
