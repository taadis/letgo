// ghttp.go
package ghttp

import (
	"net/http"
)

//
type HandlerFunc func(http.ResponseWriter, *http.Request)

//  实现接口 ServeHTTP
type GHttp struct {
	router map[string]HandlerFunc
}

// 构造函数
// 函数名称同 http.NewServeMux()
func NewServeMux() *GHttp {
	return &GHttp{router: make(map[string]HandlerFunc)}
}

//
func (g *GHttp) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + ":" + pattern
	g.router[key] = handler
}

//
func (g *GHttp) Get(pattern string, handler HandlerFunc) {
	g.addRoute(http.MethodGet, pattern, handler)
}

//
func (g *GHttp) Head(pattern string, handler HandlerFunc) {
	g.addRoute(http.MethodHead, pattern, handler)
}

//
func (g *GHttp) Post(pattern string, handler HandlerFunc) {
	g.addRoute(http.MethodPost, pattern, handler)
}

//
func (g *GHttp) Put(pattern string, handler HandlerFunc) {
	g.addRoute(http.MethodPut, pattern, handler)
}

//
func (g *GHttp) Patch(pattern string, handler HandlerFunc) {
	g.addRoute(http.MethodPatch, pattern, handler)
}

//
func (g *GHttp) Delete(pattern string, handler HandlerFunc) {
	g.addRoute(http.MethodDelete, pattern, handler)
}

//
func (g *GHttp) Connect(pattern string, handler HandlerFunc) {
	g.addRoute(http.MethodConnect, pattern, handler)
}

//
func (g *GHttp) Options(pattern string, handler HandlerFunc) {
	g.addRoute(http.MethodOptions, pattern, handler)
}

//
func (g *GHttp) Trace(pattern string, handler HandlerFunc) {
	g.addRoute(http.MethodTrace, pattern, handler)
}

//
func (g *GHttp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + ":" + r.URL.Path
	if handler, ok := g.router[key]; ok {
		handler(w, r)
	} else {
		http.NotFound(w, r)
	}
}
