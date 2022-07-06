# letgo

Let's Go.

渐进式的 go 轮子.

## 技能栈

- [gin]()
- [gin cors](github.com/gin-contrib/cors)
- ...

## 结构

- app - 应用层
- conf - 配置层
- store - 存储层
- util - 工具层

## app

应用层, 对外提供具体的应用程序.

常见的应用程序:

- CLI - 命令行接口
- API - 应用程序接口 (Rest/RPC/WebService/...)
- Web - 网站
- Client - 客户端程序 (PC/Mobile App)
- ...

## conf

配置层, 提供统一的配置操作.

常见的配置文件:

- ini
- xml
- json
- yml
- ...

## store

存储层, 对外提供持久化相关的操作, 但不暴露任何具体的数据库对象.

我们所说的持久化通常是把数据存储到某个地方, 比如:

- memory
- sqlite
- oracle
- mssql
- mysql
- pgsql
- tidb
- redis
- ...

所以如果需要切换数据库, 可以自行调整内部实现.

这样外部导入时通常是下面这样的, 也不需要单独实例化一个操作的结构体.

``` go
import (	
	userStore "github.com/taadis/store/user"
	orderStore "github.com/taadis/store/order"
)
```

## util

工具层, 封装一些小工具都放到这里, 方便使用.

## ...
