# h5,keras,savedmodel等格式有什么区别?

🎯 主要模型格式对比

## 1.`.h5`格式(HDF5)

特点:

- 传统格式:TensorFlow 1.x时期的主要格式
- 单个文件:所有内容存储在一个`.h5`文件中
- 兼容性好:支持大多数深度学习框架

使用场景:

- 需要与旧代码或工具兼容时
- 简单的模型保存和加载
- 跨平台分享模型

示例:

```python
# 保存
model.save('your/path/model.h5')

# 加载
model = tf.keras.models.load_model('model.h5')
```

## 2.`.keras`格式(Keras v3)

特点:

- 新标准:Keras 3.0+的推荐格式
- 基于Zip:实际上是zip压缩包,内含多个文件
- 框架无关:可以在TensorFlow/JAX/PyTorch等多种后端使用

适用场景:

- 新的项目(推荐使用)
- 需要跨框架兼容性
- 希望使用新特性等

示例:

```python
# 保存
model.save('your/path/model.keras')

# 加载
model = tf.keras.models.load_model('your/path/model.keras')
```

## 3.SavedModel格式(文件夹)

特点:

- TensorFlow原生: TensorFlow的官方格式
- 文件夹结构:包含多个文件(saved_model.pb+variables目录等)
- 生产就绪:支持TensorFlow Serving/TFLite等

适用场景:

- 生产环境部署
- 适用TensorFlow Serving
- 需要转化为TFLite移动端模型
- 需要完整的签名定义和元数据

示例:

```
# 保存(需要使用export方法)
model.export('saved_model/')

# 记载
model = tf.saved_model.load('saved_model/')
```

## 🎯 总结建议

对于不同项目和阶段以及兼容性要求，按需保存或导出为需要的格式.

比如对于mnist项目,可以:

- 开发阶段:使用`.keras`格式
- 生产阶段:使用`SavedModel`格式
- 若要兼容旧版,可以使用`.h5`格式
