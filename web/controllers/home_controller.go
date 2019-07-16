package controllers

import (
	"time"

	"github.com/kataras/iris/mvc"
)

type HomeController struct {
}

//
func (c *HomeController) Get() mvc.Result {
	return mvc.View{
		Name: "home/index.html",
		Data: time.Now(),
	}
}
