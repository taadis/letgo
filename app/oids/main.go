package oids

import "net/http"

// authorize 授权请求处理
func authorize(w http.ResponseWriter, r *http.Request) {
	err := srv.HandleAuthorizeRequest(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

// token 令牌请求处理
func token(w http.ResponseWriter, r *http.Request) {
	err := srv.HandleTokenRequest(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

// todo:用哪个库做的例子?继续补充学习下
func main() {
	manager := manage.NewManager()

	http.HandleFunc("/authorize", authorize)
	http.HandleFunc("/token", token)
}
