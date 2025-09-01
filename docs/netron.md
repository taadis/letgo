# netron

netron 是一个功能强大的开源可视化工具，专门用于可视化神经网络、深度学习和机器学习模型。

它支持包括TFlite(`.tflite`)在内的多种模型格式（如ONNX/Keras/CoreML等）。

## 安装和使用

安装

```
pip install netron
```

运行以下命令使用, 或者直接把`.tflite`等模型文件拖到 [https://netron.app/](Netron网页版) 查看.

```
netron your_model.tflite
```

不出意外的话，输出类似如下,打开本地地址访问即可.

```
Serving your_model.tflite' at http://localhost:8080
```
