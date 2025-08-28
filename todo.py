# # 简单的验证脚本
# import tensorflow as tf
# import numpy as np

# print("TensorFlow Version:", tf.__version__)

# # 创建一个简单的模型进行测试
# model = tf.keras.Sequential([
#     tf.keras.layers.Dense(units=1, input_shape=[1])
# ])
# model.compile(optimizer='sgd', loss='mean_squared_error')

# # 转换为TFLite格式
# converter = tf.lite.TFLiteConverter.from_keras_model(model)
# tflite_model = converter.convert()

# print("TFLite模型转换成功！模型大小:", len(tflite_model), "bytes")


# # 示例代码-图像分类模型
# # python scripts/basic_example.py

# import tensorflow as tf
# import numpy as np

# # 1. 创建一个简单的线性回归模型
# def create_and_train_model():
#     # 训练数据
#     x_train = np.array([1, 2, 3, 4, 5], dtype=float)
#     y_train = np.array([2, 4, 6, 8, 10], dtype=float)
    
#     # 创建模型
#     model = tf.keras.Sequential([
#         tf.keras.layers.Dense(units=1, input_shape=[1])
#     ])
    
#     model.compile(optimizer='sgd', loss='mean_squared_error')
    
#     # 训练模型
#     model.fit(x_train, y_train, epochs=100, verbose=0)
    
#     return model

# # 2. 转换为TFLite格式
# def convert_to_tflite(model):
#     # 原来的转换方法
#     # converter = tf.lite.TFLiteConverter.from_keras_model(model)
#     # tflite_model = converter.convert()
#     model.save('temp_model.h5')
    
#     # 保存模型
#     with open('models/linear_regression.tflite', 'wb') as f:
#         f.write(tflite_model)
    
#     return tflite_model

# # 3. 加载和运行TFLite模型
# def run_tflite_model(tflite_model_path):
#     # 加载TFLite模型
#     interpreter = tf.lite.Interpreter(model_path=tflite_model_path)
#     interpreter.allocate_tensors()
    
#     # 获取输入输出细节
#     input_details = interpreter.get_input_details()
#     output_details = interpreter.get_output_details()
    
#     # 准备测试数据
#     test_data = np.array([[6.0]], dtype=np.float32)
    
#     # 设置输入
#     interpreter.set_tensor(input_details[0]['index'], test_data)
    
#     # 运行推理
#     interpreter.invoke()
    
#     # 获取输出
#     output = interpreter.get_tensor(output_details[0]['index'])
#     print(f"输入: 6.0, 预测输出: {output[0][0]}")

# if __name__ == "__main__":
#     print("创建和训练模型...")
#     model = create_and_train_model()
    
#     print("转换为TFLite格式...")
#     tflite_model = convert_to_tflite(model)
    
#     print("运行TFLite模型...")
#     run_tflite_model('models/linear_regression.tflite')