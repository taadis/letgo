


前端请求

```
curl -H "traceparent: 00-4bf92f3577b34da6a3ce929d0e0e4736-00f067aa0ba902b7-01" http://localhost:8080/
```

这个请求会产生类似这样的日志输出：

```
{"time":"2024-01-01T12:00:00+08:00","level":"INFO","msg":"home log...","trace_id":"4bf92f3577b34da6a3ce929d0e0e4736","span_id":"新生成的span_id"}
```

traceparent 头的格式说明：

- 00 : 版本号
- 4bf92f3577b34da6a3ce929d0e0e4736 : trace ID
- 00f067aa0ba902b7 : parent span ID
- 01 : trace flags (采样标志)

当前的中间件代码已经正确支持了这个功能，它会：

1. 从请求头中提取 traceparent
2. 保持相同的 trace ID
3. 生成新的 span ID 作为子 span
4. 将上下文传递给后续的处理函数


主要改动：

1. 添加了 extractContext 方法来统一处理 trace context 的提取
2. 优先使用标准的 W3C trace context（traceparent）
3. 如果没有 traceparent，则尝试使用自定义的 x-trace-id
4. 提取逻辑复用于 Handler 和 HandlerFunc
现在可以通过以下两种方式传递 trace ID：

1. 标准方式：
```bash
curl -H "traceparent: 00-4bf92f3577b34da6a3ce929d0e0e4736-00f067aa0ba902b7-01" http://localhost:8080/
```

2. 自定义方式：
```bash
curl -H "x-trace-id: 4bf92f3577b34da6a3ce929d0e0e4736" http://localhost:8080/
```
两种方式都会正确地传递 trace ID 到后端系统。
