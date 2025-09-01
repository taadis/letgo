import tensorflow as tf
import numpy as np
import os
from datetime import datetime

# 极简稳定版本
x = np.array([[-1], [0], [1], [2], [3], [4]], dtype=np.float32)
y = np.array([[-3], [-1], [1], [3], [5], [7]], dtype=np.float32)

# 使用函数式API
inputs = tf.keras.layers.Input(shape=(1,))
outputs = tf.keras.layers.Dense(1)(inputs)
model = tf.keras.Model(inputs=inputs, outputs=outputs)

model.compile(optimizer='sgd', loss='mse')
model.fit(x, y, epochs=1)
print(f"model.summary=")
model.summary()

# 先保存为SavedModel格式
saved_model_path = "./.output/saved_model"
os.makedirs(saved_model_path, exist_ok=True)
model.export(saved_model_path)


converter = tf.lite.TFLiteConverter.from_saved_model(saved_model_path)
tflite_model = converter.convert()

timestamp = datetime.now().strftime("%Y%m%d_%H%M%S")
tflite_model_path = f"./.output/mnist_model_{timestamp}.tflite"
with open(tflite_model_path, 'wb') as f:
    f.write(tflite_model)

print("✅ 转换成功！")
