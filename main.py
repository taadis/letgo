import tensorflow as tf

if __name__ == "__main__":
    # 1.检查tensorflow版本
    print("tensorflow version:", tf.__version__) # 2.20.0
    # 没有单独的tf.lite__version__字段
    # print("tensorflow lite version:", tf.lite.__version__)

    # 2.检查是否有可用的GPU
    # 3.列出所有
    list_physical_devices = tf.config.list_physical_devices('CPU')
    print(f"列出所有物理设备CPU:{list_physical_devices}")
    list_physical_devices = tf.config.list_physical_devices('GPU')
    print(f"列出所有物理设备GPU:{list_physical_devices}")
    list_physical_devices = tf.config.list_physical_devices('TPU')
    print(f"列出所有物理设备TPU:{list_physical_devices}")
