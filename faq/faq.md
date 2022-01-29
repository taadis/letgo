# faq

## 结构体

空结构体的内存占用

空结构的作用

- 仅含方法的结构体
- 实现集合set
- 无实际数据类型的channel
- ...

## Encoding

- json
- xml
- [gob](encoding_gob_test.go)
- ...

## 并发相关

- [channel的流水线简单示例](channel_pipeline_test.go)
- [WaitGroup](faq_wait_group_test.go)
- channel超时处理-使用time.After
- channel超时处理-使用context.WithTimeout
