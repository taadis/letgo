package http

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

// http context
type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
}

//
func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		ResponseWriter: w,
		Request:        r,
	}
}

//
func (c *Context) NotFound() {
	http.NotFound(c.ResponseWriter, c.Request)
}

//
func (c *Context) Saw(statusCode int, payload []byte) {
	c.ResponseWriter.WriteHeader(statusCode)
	c.ResponseWriter.Write(payload)
}

//
func (c *Context) Plain(statusCode int, payload []byte) {
	c.ResponseWriter.WriteHeader(statusCode)
	c.ResponseWriter.Header().Set("Context-Type", "text/plain")
	_, err := c.ResponseWriter.Write(payload)
	if err != nil {
		http.Error(c.ResponseWriter, err.Error(), http.StatusInternalServerError)
	}
}

//
func (c *Context) Json(statusCode int, payload interface{}) {
	c.ResponseWriter.WriteHeader(statusCode)
	c.ResponseWriter.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(c.ResponseWriter)
	if err := encoder.Encode(payload); err != nil {
		http.Error(c.ResponseWriter, err.Error(), http.StatusInternalServerError)
	}
}

//
func (c *Context) Xml(statusCode int, payload interface{}) {
	c.ResponseWriter.WriteHeader(statusCode)
	c.ResponseWriter.Header().Set("Context-Type", "application/xml")
	encoder := xml.NewEncoder(c.ResponseWriter)
	if err := encoder.Encode(payload); err != nil {
		http.Error(c.ResponseWriter, err.Error(), http.StatusInternalServerError)
	}
}

//
// func (c *Context) Html(statusCode, html []byte) {
// 	c.ResponseWriter.WriteHeader(statusCode)
// 	c.ResponseWriter.Header().Set("Content-Type", "text/html")
// 	c.ResponseWriter.Write(html)
// }
