# Learn Open Wake Word

- [github.com/dscripka/openWakeWord](https://github.com/dscripka/openWakeWord)

以下是一个简单的本地 Python 热词唤醒示例，展示如何使用 `openWakeWord` 库进行实时麦克风音频流的热词检测。这个示例基于文档中的说明，特别是“Usage”部分的内容，并假设你已经安装了必要的依赖。

### 前提条件
1. **安装 openWakeWord**：
   确保已安装 `openwakeword` 库：
   ```bash
   pip install openwakeword
   ```
   如果在 Linux 系统上想启用 Speex 噪声抑制，可以安装相关依赖：
   ```bash
   sudo apt-get install libspeexdsp-dev
   pip install https://github.com/dscripka/openWakeWord/releases/download/v0.1.1/speexdsp_ns-0.1.2-cp38-cp38-linux_x86_64.whl
   ```

2. **安装麦克风支持**：
   需要安装 `pyaudio` 来捕获麦克风音频：
   ```bash
   pip install pyaudio
   ```

3. **下载预训练模型**：
   示例中会自动下载预训练模型，或者你也可以手动从 GitHub 发布页面下载。

### 本地 Python 热词唤醒示例代码
以下是一个简单的 Python 脚本，展示如何使用 `openWakeWord` 从本地麦克风捕获音频并检测热词（如 "alexa" 或 "hey jarvis"）：

```python
import openwakeword
from openwakeword.model import Model
import pyaudio
import numpy as np
import time

# 下载所有预训练模型（仅需运行一次）
# openwakeword.utils.download_models()
# 仅下载指定的模型("alexa"),而不是所有
openwakeword.utils.download_models("alexa")

# 初始化 openWakeWord 模型
# 可以指定特定模型，例如 ["alexa", "hey jarvis"]，或留空以加载所有模型
model = Model(
    wakeword_models=["alexa"],  # 仅检测 "alexa" 热词
    enable_speex_noise_suppression=True,  # 启用噪声抑制（仅限 Linux）
    vad_threshold=0.5  # 启用语音活动检测（VAD），阈值设为 0.5
)

# 音频流参数
FORMAT = pyaudio.paInt16  # 16-bit 音频
CHANNELS = 1  # 单声道
RATE = 16000  # 采样率 16kHz
CHUNK = int(0.08 * RATE)  # 80ms 音频帧（1280 样本）

# 初始化 PyAudio
audio = pyaudio.PyAudio()

# 打开麦克风流
stream = audio.open(
    format=FORMAT,
    channels=CHANNELS,
    rate=RATE,
    input=True,
    frames_per_buffer=CHUNK
)

print("开始监听麦克风音频... 说 'alexa' 来触发检测")

try:
    while True:
        # 读取 80ms 的音频数据
        audio_data = stream.read(CHUNK, exception_on_overflow=False)
        
        # 将音频数据转换为 numpy 数组（16-bit PCM）
        audio_frame = np.frombuffer(audio_data, dtype=np.int16)
        
        # 获取模型预测
        prediction = model.predict(audio_frame)
        
        # 检查是否有热词被检测到
        for wakeword, score in prediction.items():
            if score > 0.5:  # 使用默认阈值 0.5
                print(f"检测到热词 '{wakeword}'，置信度: {score:.2f}")
        
        # 控制循环速度（可选）
        time.sleep(0.01)

except KeyboardInterrupt:
    print("停止监听...")

finally:
    # 清理资源
    stream.stop_stream()
    stream.close()
    audio.terminate()
```

### 代码说明
1. **模型初始化**：
   - 使用 `Model` 类初始化 `openWakeWord` 模型，指定检测的热词（例如 "alexa"）。
   - 启用了 Speex 噪声抑制（Linux 环境下）和 Silero VAD（语音活动检测）以提高在嘈杂环境中的表现。
   - `vad_threshold=0.5` 确保只有当语音活动检测得分高于 0.5 时才会触发热词检测，减少误触发。

2. **音频捕获**：
   - 使用 `pyaudio` 从麦克风捕获 16-bit、16kHz 的单声道音频。
   - 每 80ms（1280 样本）读取一次音频帧，符合 `openWakeWord` 的推荐帧长度以优化效率和延迟。

3. **热词检测**：
   - `model.predict(audio_frame)` 返回一个字典，包含每个热词的置信度得分。
   - 如果某个热词的得分超过阈值（默认 0.5），则打印检测结果。

4. **清理**：
   - 程序通过 `KeyboardInterrupt`（Ctrl+C）优雅退出，关闭音频流并释放资源。

### 运行示例
1. 保存上述代码为 `wakeword_example.py`。
2. 确保麦克风已连接并正常工作。
3. 运行脚本：
   ```bash
   python wakeword_example.py
   ```
4. 说出 "alexa"，观察终端输出是否检测到热词及置信度。

### 注意事项
- **性能优化**：
  - 如果在嘈杂环境中，启用 `enable_speex_noise_suppression=True` 和调整 `vad_threshold` 可以减少误触发。
  - 根据实际环境调整热词检测阈值（例如从 0.5 改为 0.4 或 0.6）以平衡误触发率（false-accept）和漏触发率（false-reject）。

- **模型选择**：
  - 示例中仅使用了 "alexa" 模型。你可以修改 `wakeword_models` 参数以加载其他预训练模型（如 "hey jarvis" 或 "current weather"）。

- **实时性**：
  - 该示例适用于实时检测。如果需要处理音频文件，可以使用 `model.predict_clip("path/to/wav/file")` 来测试单个 WAV 文件。

- **硬件要求**：
  - 在树莓派 3 上，单核可以同时运行 15-20 个模型，但对于低功耗微控制器（如 ESP32），建议使用更轻量的库（如 `microWakeWord`）。

### 常见问题
- **误触发率高**：
  - 尝试调整阈值（`score > 0.5` 改为更高值，如 0.7）。
  - 启用 VAD 和噪声抑制。
  - 如果仍不理想，可以训练自定义验证模型（参考文档中的“User-specific models”）。

- **支持语言**：
  - 当前仅支持英语模型。如果需要其他语言支持，可以关注未来的更新或尝试使用其他 TTS 引擎（如 Mycroft.AI 的 Mimic 3）。

- **浏览器运行**：
  - 目前不支持直接在浏览器中运行 JavaScript 版本，但可以通过 WebSocket 将音频流传输到 Python 后端（参考 `examples/web` 目录）。

### 扩展
如果你想训练自定义热词模型，可以参考文档中的 Google Colab 笔记本（简单版或详细版），通过合成语音数据快速生成新模型。训练过程需要：
1. 使用 TTS 系统生成目标热词的音频数据（几千个样本）。
2. 收集负样本数据（如 30,000 小时的非热词音频）。
3. 使用提供的训练脚本进行模型训练。

如果需要进一步帮助（如调试、优化或训练自定义模型），请提供更多细节，我可以为你提供更具体的指导！
