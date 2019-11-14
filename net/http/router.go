// router.go

package http

//
type router struct {
	routes map[string]HandlerFunc
}

//
func newRouter() *router {
	return &router{
		routes: make(map[string]HandlerFunc),
	}
}

//
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + ":" + pattern
	r.routes[key] = handler
}

//
func (r *router) handle(c *Context) {
	key := c.Request.Method + ":" + c.Request.URL.Path
	if handler, ok := r.routes[key]; ok {
		handler(c)
	} else {
		c.NotFound()
	}
}
