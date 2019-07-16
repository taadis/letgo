// main.go
package main

import (
	"letgo/web/controllers"

	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
)

// main
func main() {
	//app := iris.Default()
	app := newApp()

	// 配置和启动服务
	// 配置参考: https://www.studyiris.com/doc/irisDoc/Configuration.html
	//app.Configure(iris.WithConfiguration(iris.TOML("./config/iris.yml")))
	//ymlConfig := iris.YAML("./config/iris.yml")
	//app.Configure(iris.WithConfiguration(ymlConfig))
	app.Run(iris.Addr(":80")) // 使用默认配置
}

//
func newApp() *iris.Application {
	jwthandler := jwtmiddleware.New(jwtmiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("mySecret"), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(jwthandler.Serve)

	// 注册视图(模板引擎)
	viewEngine := iris.HTML("./web/views", ".html").Layout("shared/layout.html").Reload(true)
	app.RegisterView(viewEngine)

	// 设置静态文件目录
	app.StaticWeb("/static", "./web/static")

	// 处理异常
	app.OnAnyErrorCode(func(ctx iris.Context) {
		//ctx.ViewLayout("")
		ctx.View("shared/error.html")
	})

	// Use MVC
	//mvc.New(app.Party("/").Handle())
	mvc.New(app).Handle(new(controllers.HomeController))

	// app.Handle("GET", "/", func(ctx iris.Context) {
	// 	ctx.HTML("<h1>Welcome</h1>")
	// })
	app.Get("/hello/{name}", func(ctx iris.Context) {
		ctx.Writef("hello %s", ctx.Params().Get("name"))
	})

	return app
}
