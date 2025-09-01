# tf.version

`tf.version`模块是 TensorFlow 提高的用于获取库版本以及相关信息的公共 API.

## 概述

`tf.version`模块提供了一些只读属性，用于获取TensorFlow库的版本信息、编译器版本以及图定义(GraphDef)版本等.

这些信息在调试、检查兼容性或者记录环境信息时非常有用。

## 成员属性

`tf.version`模块包含以下属性

> 这些属性都是只读的,可以直接访问以获取版本信息.

- `COMPILER_VERSION`:用于构建TensorFlow的编译器版本，例如`Ubuntu Clang 17.0.6 (++20231208085846+6009708b4367-1~exp1~20231208085949.74`
- `GIT_VERSION`:TensorFlow源代码的git版本,例如`v2.16.1-0-g5bc9d26649c`.
- `GRAPH_DEF_VERSION`:TensorFlow使用的GraphDef协议版本,例如`1766`.
- `GRAPH_DEF_VERSION_MIN_CONSUMER`:消费者支持的最小GraphDef版本,例如`0`.
- `GRAPH_DEF_VERSION_MIN_PRODUCER`:生产者支持的最小GraphDef版本,例如`0`.
- `VERSION`:TensorFlow库的版本,例如`2.16.1`

## 使用示例

要使用`tf.version`模块,需要先导入TensorFlow,然后直接访问版本模块下的这些属性.

以下为pyhon代码示例(v.py):

```python
import tensorflow as tf

# print(f"version:{tf.__version__}")

print(f"tf.verison.COMPILER_VERSION={tf.version.COMPILER_VERSION}")

print(f"tf.version.GIT_VERSION={tf.version.GIT_VERSION}")

print(f"tf.version.GRAPH_DEF_VERSION={tf.version.GRAPH_DEF_VERSION}")

print(f"tf.version.GRAPH_DEF_VERSION_MIN_CONSUMER={tf.version.GRAPH_DEF_VERSION_MIN_CONSUMER}")

print(f"tf.version.GRAPH_DEF_VERSION_MIN_PRODUCER={tf.version.GRAPH_DEF_VERSION_MIN_PRODUCER}")

print(f"tf.version.VERSION={tf.version.VERSION}")

```

不出意外的话,输出信息如下:

```
tf.verison.COMPILER_VERSION=Apple LLVM 14.0.3 (clang-1403.0.22.14.1)
tf.version.GIT_VERSION=v2.16.1-19-g810f233968c
tf.version.GRAPH_DEF_VERSION=1766
tf.version.GRAPH_DEF_VERSION_MIN_CONSUMER=0
tf.version.GRAPH_DEF_VERSION_MIN_PRODUCER=0
tf.version.VERSION=2.16.2
```

## 参考

- [Module:tf.verison](https://www.tensorflow.org/api_docs/python/tf/version)
