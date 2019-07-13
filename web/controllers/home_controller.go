package controllers

import(
	"github.com/kataras/iris/mvc"
)

type HomeController struct {
}

//
func (c *HomeController) Get() mvc.Result{
	return mvc.Response{
		ContentType: "text/html",
		Content: "<h1>HomeController Get()</h1>"
	}
}
