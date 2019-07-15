package handler

import (
	"github.com/kataras/iris"
)

//
func Add(ctx *iris.Context) {
	username := ctx.Params("username")
	ctx.JSON(iris.StatusOK)
}
