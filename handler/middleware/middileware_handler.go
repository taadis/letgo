package midleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

//
type CustomHandler struct {
}

//
type CustomModel struct {
	Type string `json:"type,omitempty"`
}

// 处理函数
func Custom(w http.ResponseWriter, r *http.Request) {
	//http.NotFound(w, r)
	http.Error(w, "custom handle func not implemented", http.StatusNotImplemented)
}

// JSON 响应
func Json(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	model := &CustomModel{
		Type: "json",
	}
	jsonBytes, err := json.Marshal(model)
	if err != nil {
		panic(err)
	}
	w.Write(jsonBytes)
}

// 301 永久转移
func MovedPermanently(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://taadis.com")
	w.WriteHeader(http.StatusMovedPermanently)
}

// 302 临时转移
func Found(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://taadis.com")
	w.WriteHeader(http.StatusFound)
}

// 501 服务器内出错
func InternalServerError(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "服务器开小差了, 请稍后重试!", http.StatusInternalServerError)
}

// 中间件处理器
func MiddlewareHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("middleware handler before")
		next.ServeHTTP(w, r)
		log.Println("middleware handler after")
	})
}

// 请求日志处理器
func LoggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.RawPath)
		next.ServeHTTP(w, r)
		log.Printf("Comleted %s %s in %v", r.Method, r.URL.Path, time.Since(start))
	})
}

// 崩溃恢复
// 规避 panic 情况下应用直接跪了的问题.
// 当然也可以用来更好友好的返回响应.
func PanicAndRecover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rc := recover(); rc != nil {
				s := fmt.Sprintf("panic recover %v", rc)
				fmt.Println(s)
				http.Error(w, s, http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
