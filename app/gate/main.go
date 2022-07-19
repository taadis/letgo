package main

import (
	"context"

	"go-micro.dev/v4/api"
	"go-micro.dev/v4/logger"
)

// TODO:
// [2022-07-19] go-micro/v4 重新调整了 api package 来重新实现v1/v2版本的 api gateway, 过段时间再看看能不能实际使用?
func main() {
	//  Usage:
	//
	// 	proto.RegisterHandler(service.Server(), new(Handler), api.WithEndpoint(
	//		&api.Endpoint{
	//			Name: "Greeter.Hello",
	//			Path: []string{"/greeter"},
	//		},
	//	))
	//api.WithEndpoint(&api.Endpoint{Name: "Greeter.Hello", Path: []string{"/"}})
	//api.WithEndpoint(&api.Endpoint{Name: "User.Login", Path: []string{"/user/login"}})

	a := api.NewApi()

	var err error
	err = a.Init()
	if err != nil {
		logger.Errorf("api init error:%+v", err)
		return
	}

	err = a.Run(context.Background())
	if err != nil {
		logger.Errorf("api run error:%+v", err)
	}
}
