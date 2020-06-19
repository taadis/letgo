# letgo

Let's Go.

渐进式的 go 轮子.

## 结构

- app - 应用层
- conf - 配置层
- store - 存储层

## app

应用层, 对外提供具体的应用程序.

常见的应用程序:

- CLI - 命令行接口
- API - 应用程序接口 (Rest/RPC/WebService/...)
- Web - 网站
- Client - 客户端程序 (PC/Mobile App)
- ...

## conf

配置层

## store

存储层, 对外提供持久化相关的操作, 但不暴露任何具体的数据库对象.

所以如果需要切换数据库, 可以自行调整内部实现.

这样外部导入时通常是下面这样的, 也不需要单独实例化一个操作的结构体.

``` go
import (	
	userStore "gitee.com/taadis/store/user"
	orderStore "gitee.com/taadis/store/order"
)
```

## ...
