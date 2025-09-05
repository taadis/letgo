import openwakeword
from openwakeword.model import Model
import pyaudio
import numpy as np
import time
import sys
import logging

# 设置日志以便调试
logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(levelname)s - %(message)s')

# 下载所有预训练模型（仅需运行一次）
try:
    logging.info("正在下载预训练模型...")
    openwakeword.utils.download_models(["alexa"])
except Exception as e:
    logging.error(f"下载模型失败: {e}")
    sys.exit(1)

# 初始化 openWakeWord 模型
try:
    logging.info("正在初始化 openWakeWord 模型...")
    model = Model(
        wakeword_models=["alexa"],  # 仅检测 "alexa" 热词
        vad_threshold=0.5  # 启用语音活动检测（VAD）
    )
except Exception as e:
    logging.error(f"模型初始化失败: {e}")
    sys.exit(1)

# 音频流参数
FORMAT = pyaudio.paInt16  # 16-bit 音频
CHANNELS = 1  # 单声道
RATE = 16000  # 采样率 16kHz
CHUNK = int(0.08 * RATE)  # 80ms 音频帧（1280 样本）

# 初始化 PyAudio
try:
    audio = pyaudio.PyAudio()
except Exception as e:
    logging.error(f"PyAudio 初始化失败: {e}")
    sys.exit(1)

# 打开麦克风流
try:
    stream = audio.open(
        format=FORMAT,
        channels=CHANNELS,
        rate=RATE,
        input=True,
        frames_per_buffer=CHUNK
    )
    logging.info("麦克风流已打开，开始监听...")
except Exception as e:
    logging.error(f"无法打开麦克风流: {e}")
    audio.terminate()
    sys.exit(1)

print("开始监听麦克风音频... 说 'alexa' 来触发检测")

try:
    while True:
        # 读取 80ms 的音频数据
        try:
            audio_data = stream.read(CHUNK, exception_on_overflow=False)
        except Exception as e:
            logging.warning(f"读取音频数据失败: {e}")
            continue

        # 将音频数据转换为 numpy 数组（16-bit PCM）
        try:
            audio_frame = np.frombuffer(audio_data, dtype=np.int16)
        except Exception as e:
            logging.warning(f"音频数据转换失败: {e}")
            continue

        # 获取模型预测
        try:
            prediction = model.predict(audio_frame)
            # 检查是否有热词被检测到
            for wakeword, score in prediction.items():
                if score > 0.5:  # 使用默认阈值 0.5
                    print(f"检测到热词 '{wakeword}'，置信度: {score:.2f}")
        except Exception as e:
            logging.warning(f"模型预测失败: {e}")
            continue

        # 控制循环速度
        time.sleep(0.01)

except KeyboardInterrupt:
    print("停止监听...")

finally:
    # 清理资源
    logging.info("正在清理资源...")
    try:
        stream.stop_stream()
        stream.close()
        audio.terminate()
    except Exception as e:
        logging.error(f"资源清理失败: {e}")
