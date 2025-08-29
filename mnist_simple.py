import tensorflow as tf
import numpy as np
import ssl
from datetime import datetime

# 临时禁用SSL证书验证
ssl._create_default_https_context = ssl._create_unverified_context

# 1.加载数据 - 就像拿到一堆写数字的图片和对应的标签
# 首次会从远端下载,并缓存到~/.keras/datasets/mnist.npz
# Downloading data from https://storage.googleapis.com/tensorflow/tf-keras-datasets/mnist.npz
print("📦 1.正在加载手写数字数据集...")
(train_images, train_labels), (test_images, test_labels) = tf.keras.datasets.mnist.load_data()

# 归一化像素值到0-1范围
print("🔧 2.正在预处理数据...")
train_images = train_images / 255.0
test_images = test_images / 255.0

# 3.搭建神经网络 - 配个脑子
print("🧠 3.正在构建神经网络模型...")
#  tf.keras.Sequential 这个函数的作用是什么?参数分别是什么意思?
model = tf.keras.Sequential([
    # 将28*28的图像展平为784个像素
    # 使用Input层明确指定输入形状
    tf.keras.layers.Input(shape=(28, 28)),
    tf.keras.layers.Flatten(),
    # 全连接层,128个神经元
    tf.keras.layers.Dense(128, activation='relu'),
    # 输出层,10个神经元对应0-9数字
    # 您的模型缺少了输出层的激活函数。在MNIST这样的多分类问题中，输出层需要使用softmax激活函数来将输出转换为概率分布
    tf.keras.layers.Dense(10, activation='softmax')
])

# 4.配置学习方式 - 告诉学生/电脑如何学习
print("⚙️ 4.正在编译模型...")
model.compile(
    # 学习算法:像聪明的学生,能自己调整学习速度
    optimizer='adam',
    # 损失函数:衡量猜错的程度
    loss='sparse_categorical_crossentropy',
    # 评估标准:看猜对的准确率
    metrics=['accuracy'],
)

# 5.开始训练 - 就像老师讲学生/电脑认字
print("🚀 5.开始训练模型...")
model.fit(
    # 训练数据和正确答案
    train_images,
    train_labels,
    # 学习5遍
    epochs=5,
    # todo:每次看多少张图片?
    batch_size=32,
    # 留20%的数据用来检查本次学习效果
    validation_split=0.2,
)

# 6.考试测验 - 用没见过的图片测试学习和掌握情况
print("📊 6.正在评估模型性能...")
test_loss, test_acc = model.evaluate(test_images, test_labels, verbose=2)
print(f"\n测试准确率:{test_acc:.4f}")

# 7.保存学习成果
print("💾 7.正在保存训练好的模型(.keras格式)...")
# e.g.20250829_164832
timestamp = datetime.now().strftime("%Y%m%d_%H%M%S")
model_path = f"./.output/mnist_model_{timestamp}.keras"
model.save(model_path)
print("✅ 模型已保存为 'mnist_model' 文件夹")

# 输出模型文件信息
import os
if os.path.exists(model_path):
    size_kb = os.path.getsize(model_path) / 1024
    abs_path = os.path.abspath(model_path)
    print(f"📄 模型路径: {abs_path}")
    print(f"📊 文件大小: {size_kb:.2f} KB")
else:
    print("❌ 模型文件保存失败")
