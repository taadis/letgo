// main.go
package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	//"github.com/kataras/iris/mvc"
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
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())

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

	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome</h1>")
	})
	return app
}
