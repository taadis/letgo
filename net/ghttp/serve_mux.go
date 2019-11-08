// http.go
package ghttp

import (
	"net/http"
)

//
//type HandlerFunc func(http.ResponseWriter, *http.Request)
type HandlerFunc func(*Context)

//  实现接口 ServeHTTP
type ServeMux struct {
	router map[string]HandlerFunc
}

// 构造函数
// 函数名称同 http.NewServeMux()
func NewServeMux() *ServeMux {
	return &ServeMux{
		router: make(map[string]HandlerFunc),
	}
}

//
func (mux *ServeMux) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + ":" + pattern
	mux.router[key] = handler
}

//
func (mux *ServeMux) Get(pattern string, handler HandlerFunc) {
	mux.addRoute(http.MethodGet, pattern, handler)
}

//
func (mux *ServeMux) Head(pattern string, handler HandlerFunc) {
	mux.addRoute(http.MethodHead, pattern, handler)
}

//
func (mux *ServeMux) Post(pattern string, handler HandlerFunc) {
	mux.addRoute(http.MethodPost, pattern, handler)
}

//
func (mux *ServeMux) Put(pattern string, handler HandlerFunc) {
	mux.addRoute(http.MethodPut, pattern, handler)
}

//
func (mux *ServeMux) Patch(pattern string, handler HandlerFunc) {
	mux.addRoute(http.MethodPatch, pattern, handler)
}

//
func (mux *ServeMux) Delete(pattern string, handler HandlerFunc) {
	mux.addRoute(http.MethodDelete, pattern, handler)
}

//
func (mux *ServeMux) Connect(pattern string, handler HandlerFunc) {
	mux.addRoute(http.MethodConnect, pattern, handler)
}

//
func (mux *ServeMux) Options(pattern string, handler HandlerFunc) {
	mux.addRoute(http.MethodOptions, pattern, handler)
}

//
func (mux *ServeMux) Trace(pattern string, handler HandlerFunc) {
	mux.addRoute(http.MethodTrace, pattern, handler)
}

//
func (mux *ServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + ":" + r.URL.Path
	if handler, ok := mux.router[key]; ok {
		//handler(w, r)
		context := NewContext(w, r)
		handler(context)
	} else {
		http.NotFound(w, r)
	}
}

//
//func (mux *ServeMux) handle()
