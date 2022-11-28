# echo-server

echo-server是用Go语言实现一个简单echo服务器.

- 1.创建一个监听套接字
- 2.通过Accept获得每次请求的新链接
- 3.调用handle函数异步处理读/写

主要了解下Listen\Accept\Read\Write四个函数的使用和实现.

## 测试

通过telnet命令行工具来测试

```
telnet 127.0.0.1 5903
Trying 127.0.0.1...
Connected to localhost.
Escape character is '^]'.
1
1
2
2
ping
ping
^]
telnet> quit
Connection closed.
	
```

通过nc命令行工具来测试

```
nc 127.0.0.1 5903
1
1
2
2
3
3
^C

```
